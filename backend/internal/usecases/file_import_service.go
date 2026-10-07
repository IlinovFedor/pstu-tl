package usecases

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"uuid"

	"github.com/IlinovFedor/pstu-tl/internal/models"
	"github.com/xuri/excelize/v2"
)

type FileImportService struct {
	transactor          Transactor
	fileRepo            FileImportRepository
	disciplineRepo      DisciplineRepository
	facultyRepo         FacultyRepository
	specializationsRepo SpecializationRepository
	groupRepo           GroupRepository
	recordsRepo         RecordsRepository
	storage             FileStorage
}

const (
	colDiscipline = iota
	colGroup
	colStudentsAmount
	colSubgroupNumber
	colRecordType
	colGoalHours
)

var regexGroupName = regexp.MustCompile(`^([А-яЁё]+) +([А-яЁё\d]+) *-(\d\d)-(\d+(?:,\d+)*)([А-яЁё]+)$`)

const (
	submatchFaculty = iota + 1
	submatchSpecialization
	submatchYearOfEnrollment
	submatchSequenceNumbers
	submatchGroupType
)

var mapSheetNameToTerm = map[string]models.Term{
	"Осенний семестр":  models.TermFirst,
	"Весенний семестр": models.TermSecond,
}

var mapTermToSheetName = map[models.Term]string{
	models.TermFirst:  "Осенний семестр",
	models.TermSecond: "Весенний семестр",
}

type importDiscipline struct {
	models.Discipline
}

func newImportDiscipline(discipline models.Discipline) *importDiscipline {
	return &importDiscipline{Discipline: discipline}
}

type importFaculty struct {
	models.Faculty
}

func newImportFaculty(faculty models.Faculty) *importFaculty {
	return &importFaculty{Faculty: faculty}
}

type importSpecialization struct {
	models.Specialization
}

func newImportSpecialization(specialization models.Specialization) *importSpecialization {
	return &importSpecialization{Specialization: specialization}
}

type importGroup struct {
	models.Group
	faculty        *importFaculty
	specialization *importSpecialization
}

type importSubgroups struct {
	group    *importGroup
	subgroup models.SubgroupType
}

func newImportSubgroups(group *importGroup, subgroup models.SubgroupType) *importSubgroups {
	return &importSubgroups{group: group, subgroup: subgroup}
}

func newImportGroup(group models.Group, faculty *importFaculty, specialization *importSpecialization) *importGroup {
	return &importGroup{Group: group, faculty: faculty, specialization: specialization}
}

type importRecord struct {
	models.Record
	discipline *importDiscipline
	subgroups  []importSubgroups
}

func newImportRecord(record models.Record, discipline *importDiscipline, subgroups []importSubgroups) *importRecord {
	return &importRecord{Record: record, discipline: discipline, subgroups: subgroups}
}

type parser struct {
	disciplines map[importDiscipline]*importDiscipline
	faculties   map[importFaculty]*importFaculty
	specials    map[importSpecialization]*importSpecialization
	groups      map[importGroup]*importGroup
	records     []importRecord
}

func newParser() *parser {
	return &parser{
		disciplines: make(map[importDiscipline]*importDiscipline),
		faculties:   make(map[importFaculty]*importFaculty),
		specials:    make(map[importSpecialization]*importSpecialization),
		groups:      make(map[importGroup]*importGroup),
		records:     make([]importRecord, 0),
	}
}

func intern[T comparable](m map[T]*T, v *T) *T {
	if ref, ok := m[*v]; ok {
		return ref
	}
	m[*v] = v
	return v
}

// TODO: привести в порядок
func (p *parser) parseSheet(file *excelize.File, term models.Term, fileID uuid.UUID) (err error) {
	rows, err := file.Rows(
		mapTermToSheetName[term],
	)
	if err != nil {
		return
	}
	defer func(rows *excelize.Rows) {
		errClose := rows.Close()
		if errClose != nil {
			err = errors.Join(err, errClose)
		}
	}(rows)

	parseErr := models.NewParseError(
		make([]models.CellError, 0))

	var (
		columns   []string
		rowNum    = 0
		rowFailed bool
	)

	cell := func(col int) string {
		if col < len(columns) {
			return strings.TrimSpace(columns[col])
		}
		return ""
	}
	fail := func(col int, reason, cause error) {
		rowFailed = true
		name, errName := excelize.CoordinatesToCellName(col+1, rowNum)
		if errName != nil {
			name = fmt.Sprintf("R%dC%d", rowNum, col+1)
		}
		if cause != nil {
			reason = fmt.Errorf("%w: %v", reason, cause)
		}
		parseErr.AddErr(models.NewCellError(reason, term, name, cell(col)))
	}
	atoi := func(col int, s string, reason error) (int, bool) {
		n, errAtoi := strconv.Atoi(strings.TrimSpace(s))
		if errAtoi != nil {
			fail(col, reason, errAtoi)
			return 0, false
		}
		return n, true
	}

	for rows.Next() {
		rowNum++
		// TODO: брать данные для столбцов из merged columns
		var errGetColumns error
		columns, errGetColumns = rows.Columns()
		if errGetColumns != nil {
			return errGetColumns
		}
		if len(columns) < 6 {
			continue
		}

		studentsAmount, errParseStudentsAmount := strconv.Atoi(strings.TrimSpace(columns[colStudentsAmount]))
		if errParseStudentsAmount != nil {
			fail(colStudentsAmount, models.ErrUnparsableStudentsAmount, errParseStudentsAmount)
		}

		recordType, errRecordType := models.NewRecordType(columns[colRecordType])
		if errRecordType != nil {
			fail(colRecordType, models.ErrUnknownRecordType, errRecordType)
		}

		goalHours, errParseGoalHours := strconv.Atoi(columns[colGoalHours])
		if errParseGoalHours != nil {
			fail(colGoalHours, models.ErrUnparsableGoalHours, errParseGoalHours)
		}

		subgroupType := models.SubgroupTypeBoth
		if cell(colSubgroupNumber) != "" {
			if n, ok := atoi(colSubgroupNumber, cell(colSubgroupNumber), models.ErrUnparsableSubgroupNumber); ok {
				subgroupType = models.SubgroupType(n)
			}
		}

		submatch := regexGroupName.FindStringSubmatch(columns[colGroup])
		if submatch == nil {
			fail(colGroup, models.ErrUnparsableGroupCell, nil)
			continue
		}

		yearOfEnrollment, _ := atoi(colGroup, submatch[submatchYearOfEnrollment], models.ErrUnparsableYearOfEnrollment)

		groupType, errGT := models.NewGroupTypeFromString(submatch[submatchGroupType])
		if errGT != nil {
			fail(colGroup, models.ErrUnknownGroupType, errGT)
		}

		var seqNums []int
		for s := range strings.SplitSeq(submatch[submatchSequenceNumbers], ",") {
			if n, ok := atoi(colGroup, s, models.ErrUnparsableGroupSequenceNumber); ok {
				seqNums = append(seqNums, n)
			}
		}

		disciplineName := strings.TrimSpace(columns[colDiscipline])
		if disciplineName == "" {
			fail(colDiscipline, models.ErrEmptyDisciplineCol, nil)
		}

		if rowFailed {
			continue
		}

		discipline := intern(p.disciplines, newImportDiscipline(
			*models.NewDiscipline(uuid.Nil(), disciplineName, false)))
		faculty := intern(p.faculties, newImportFaculty(
			*models.NewFaculty(uuid.Nil(), submatch[submatchFaculty], false)))
		specialization := intern(p.specials, newImportSpecialization(
			*models.NewSpecialization(uuid.Nil(), submatch[submatchSpecialization], false)))

		subgroups := make([]importSubgroups, 0, len(seqNums))
		for _, num := range seqNums {
			group := intern(p.groups, newImportGroup(
				*models.NewGroup(uuid.Nil(), uuid.Nil(), uuid.Nil(), yearOfEnrollment, num, groupType, false),
				faculty, specialization))
			subgroups = append(subgroups, *newImportSubgroups(group, subgroupType))
		}

		p.records = append(p.records,
			*newImportRecord(
				*models.NewRecord(
					uuid.Nil(),
					fileID,
					nil,
					studentsAmount,
					uuid.Nil(),
					term,
					recordType,
					goalHours,
					nil,
					false,
					false,
				),
				discipline,
				subgroups),
		)
	}

	if parseErr.HasErrors() {
		return parseErr
	}
	return nil
}

func (f *FileImportService) ImportFile(ctx context.Context, yearName int, data []byte, fileName string) (err error) {
	file, err := excelize.OpenReader(bytes.NewReader(data))
	if err != nil {
		return errors.New("failed to open excel file: " + err.Error())
	}
	defer func(reader *excelize.File) {
		errClose := reader.Close()
		if errClose != nil {
			err = errors.Join(err, errClose)
		}
	}(file)

	sheets := file.GetSheetList()

	//pars := newParser()
	for _, sheet := range sheets {
		_, ok := mapSheetNameToTerm[sheet]
		if !ok {
			continue
		}

		//pars.parseSheet(file, sheet)
	}

	// logic

	return err
}
