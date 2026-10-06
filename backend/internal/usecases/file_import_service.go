package usecases

import (
	"bytes"
	"context"
	"errors"
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
	defaultRowsCnt = 1024
)

var regexGroupName = regexp.MustCompile(`^([А-яЁё]+) +([А-яЁё\d]+) *-(\d\d)-(\d+(?:,\d+)*)([А-яЁё]+)$`)

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
	disciplines map[importDiscipline]*models.Discipline
	faculties   map[importFaculty]*models.Faculty
	specials    map[importSpecialization]*models.Specialization
	groups      map[importGroup]*models.Group
	records     []*importRecord
}

func newParser() *parser {
	return &parser{
		disciplines: make(map[importDiscipline]*models.Discipline),
		faculties:   make(map[importFaculty]*models.Faculty),
		specials:    make(map[importSpecialization]*models.Specialization),
		groups:      make(map[importGroup]*models.Group),
		records:     make([]*importRecord, defaultRowsCnt),
	}
}

// TODO: привести в порядок
func (i *parser) parseSheet(file *excelize.File, term models.Term, fileID uuid.UUID) (err error) {
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

	for rows.Next() {
		var columns []string
		// TODO: брать данные для столбцов из merged columns
		// TODO: оборачивать ошибки доп инфой
		// TODO: вероятно количество студентов нужно сделать общим для records
		columns, err = rows.Columns()
		if err != nil {
			return err
		}
		if len(columns) > 6 {
			continue
		}
		submatch := regexGroupName.FindStringSubmatch(columns[1])
		if submatch == nil {
			continue
		}

		discipline := newImportDiscipline(*models.NewDiscipline(uuid.Nil(), columns[0], false))
		if _, ok := i.disciplines[*discipline]; !ok {
			i.disciplines[*discipline] = &discipline.Discipline
		}

		faculty := newImportFaculty(*models.NewFaculty(uuid.Nil(), submatch[1], false))
		if _, ok := i.faculties[*faculty]; !ok {
			i.faculties[*faculty] = &faculty.Faculty
		}

		specialization := newImportSpecialization(*models.NewSpecialization(uuid.Nil(), submatch[2], false))
		if _, ok := i.specials[*specialization]; !ok {
			i.specials[*specialization] = &specialization.Specialization
		}

		var studentsAmount int
		studentsAmount, err = strconv.Atoi(strings.TrimSpace(columns[2]))
		if err != nil {
			return err
		}

		var yearOfEnrollment int
		yearOfEnrollment, err = strconv.Atoi(submatch[3])
		if err != nil {
			return err
		}
		groupsSeqNums := strings.Split(submatch[4], ",")
		subgroups := make([]importSubgroups, 0)

		var subgroupInt int
		subgroupType := models.SubgroupTypeBoth
		subgroupInt, err = strconv.Atoi(strings.TrimSpace(columns[3]))
		if err == nil {
			subgroupType = models.SubgroupType(subgroupInt)
		}

		for _, numStr := range groupsSeqNums {
			var num int
			num, err = strconv.Atoi(strings.TrimSpace(numStr))
			if err != nil {
				return err
			}

			group := newImportGroup(*models.NewGroup(uuid.Nil(),
				uuid.Nil(),
				uuid.Nil(),
				yearOfEnrollment,
				num,
				models.GroupTypePartTime,
				studentsAmount,
				false,
			),
				faculty,
				specialization,
			)

			if _, ok := i.groups[*group]; !ok {
				i.groups[*group] = &group.Group
			}
			subgroups = append(subgroups, *newImportSubgroups(group, subgroupType))
		}

		var recordType models.RecordType
		recordType, err = models.MapStringToRecordType(columns[4])
		if err != nil {
			return err
		}

		var hoursGoal int
		hoursGoal, err = strconv.Atoi(columns[5])
		if err != nil {
			return err
		}

		i.records = append(i.records,
			newImportRecord(
				*models.NewRecord(uuid.Nil(),
					fileID,
					nil,
					uuid.Nil(),
					term,
					recordType,
					hoursGoal,
					nil,
					false,
					false,
				),
				discipline,
				subgroups),
		)
	}

	return
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
