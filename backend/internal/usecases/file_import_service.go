package usecases

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"regexp"
	"slices"
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
	rows, err := file.Rows(term.String())
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

	fileID := uuid.NewV7()
	pars := newParser()
	parseErr := models.NewParseError(make([]models.CellError, 0))
	for _, sheet := range file.GetSheetList() {
		term, errTerm := models.NewTermFromString(sheet)
		if errTerm != nil {
			continue
		}

		errSheet := pars.parseSheet(file, term, fileID)
		var sheetParseErr *models.ParseError
		switch {
		case errSheet == nil:
		case errors.As(errSheet, &sheetParseErr):
			for _, cellErr := range sheetParseErr.Cells() {
				parseErr.AddErr(cellErr)
			}
		default:
			return fmt.Errorf("failed to parse sheet %q: %w", sheet, errSheet)
		}
	}
	if parseErr.HasErrors() {
		return parseErr
	}
	if len(pars.records) == 0 {
		return fmt.Errorf("%w: no records found", models.ErrInvalidFileData)
	}

	fileImport := models.NewFileImport(fileID, yearName, fileName, pars.hash())

	fileStored := false
	err = f.transactor.WithTx(ctx, func(ctx context.Context) error {
		_, errFind := f.fileRepo.GetFileImportByHash(ctx, yearName, fileImport.Hash())
		switch {
		case errFind == nil:
			return models.ErrDuplicateFileData
		case !errors.Is(errFind, models.ErrUnknownFile):
			return fmt.Errorf("failed to find file import by hash: %w", errFind)
		}

		if errUpsert := f.upsertGraph(ctx, pars); errUpsert != nil {
			return errUpsert
		}
		if errInsert := f.fileRepo.InsertFileImport(ctx, fileImport); errInsert != nil {
			return fmt.Errorf("failed to insert file import: %w", errInsert)
		}
		if errInsert := f.recordsRepo.InsertRecords(ctx, pars.buildRecords()); errInsert != nil {
			return fmt.Errorf("failed to insert records: %w", errInsert)
		}
		if errStore := f.storage.CreateFile(fileID, data); errStore != nil {
			return fmt.Errorf("failed to store file: %w", errStore)
		}
		fileStored = true
		return nil
	})
	if err != nil && fileStored {
		if errDelete := f.storage.DeleteFile(fileID); errDelete != nil {
			err = errors.Join(err, errDelete)
		}
	}
	return err
}

// upsertGraph сохраняет справочники парсера пачками и записывает в значения map реальные объекты из репозиториев.
// Группы сохраняются последними, так как ссылаются на факультеты и специальности
func (f *FileImportService) upsertGraph(ctx context.Context, p *parser) error {
	if err := upsertBatch(ctx, p.faculties,
		func(v *importFaculty) *models.Faculty { return &v.Faculty },
		f.facultyRepo.UpsertFaculties); err != nil {
		return fmt.Errorf("failed to upsert faculties: %w", err)
	}
	if err := upsertBatch(ctx, p.specials,
		func(v *importSpecialization) *models.Specialization { return &v.Specialization },
		f.specializationsRepo.UpsertSpecializations); err != nil {
		return fmt.Errorf("failed to upsert specializations: %w", err)
	}
	if err := upsertBatch(ctx, p.disciplines,
		func(v *importDiscipline) *models.Discipline { return &v.Discipline },
		f.disciplineRepo.UpsertDisciplines); err != nil {
		return fmt.Errorf("failed to upsert disciplines: %w", err)
	}

	for _, group := range p.groups {
		group.SetFacultyID(group.faculty.Id())
		group.SetSpecializationID(group.specialization.Id())
	}
	if err := upsertBatch(ctx, p.groups,
		func(v *importGroup) *models.Group { return &v.Group },
		f.groupRepo.UpsertGroups); err != nil {
		return fmt.Errorf("failed to upsert groups: %w", err)
	}
	return nil
}

// upsertBatch отправляет все значения map одной пачкой. Репозиторий обязан вернуть результаты
// в том же порядке, что и входные данные: i-й результат записывается в i-й переданный объект
func upsertBatch[K comparable, M any](
	ctx context.Context,
	items map[K]*K,
	model func(*K) *M,
	upsert func(context.Context, []*M) ([]M, error),
) error {
	batch := make([]*M, 0, len(items))
	for _, item := range items {
		batch = append(batch, model(item))
	}

	upserted, err := upsert(ctx, batch)
	if err != nil {
		return err
	}
	if len(upserted) != len(batch) {
		return fmt.Errorf("repository returned %d items, expected %d", len(upserted), len(batch))
	}
	for i := range batch {
		*batch[i] = upserted[i]
	}
	return nil
}

// buildRecords проставляет записям ID и ссылки на сохранённые дисциплины и группы
func (p *parser) buildRecords() []*models.Record {
	records := make([]*models.Record, 0, len(p.records))
	for i := range p.records {
		record := &p.records[i]
		groups := make([]models.GroupsSubgroups, 0, len(record.subgroups))
		for _, subgroup := range record.subgroups {
			groups = append(groups, models.NewGroupsSubgroups(subgroup.group.Id(), subgroup.subgroup))
		}
		record.SetId(uuid.NewV7())
		record.SetDisciplineId(record.discipline.Id())
		record.SetGroups(groups)
		records = append(records, &record.Record)
	}
	return records
}

// hash считает хеш по содержимому записей, а не по байтам xlsx, так как метаданные файла
// меняются при каждом сохранении. Строки сортируются, поэтому порядок записей в файле не важен
func (p *parser) hash() string {
	lines := make([]string, 0, len(p.records))
	for _, record := range p.records {
		groups := make([]string, 0, len(record.subgroups))
		for _, subgroup := range record.subgroups {
			group := subgroup.group
			groups = append(groups, fmt.Sprintf("%s/%s/%d/%d/%d/%d",
				group.faculty.Name(), group.specialization.Name(), group.YearOfEnrollment(),
				group.SequenceNumber(), group.GroupType(), subgroup.subgroup))
		}
		slices.Sort(groups)
		lines = append(lines, fmt.Sprintf("%d|%s|%d|%d|%d|%s",
			record.Term(), record.discipline.Name(), record.StudentsAmount(),
			record.RecordType(), record.GoalHours(), strings.Join(groups, ";")))
	}
	slices.Sort(lines)

	sum := sha256.Sum256([]byte(strings.Join(lines, "\n")))
	return hex.EncodeToString(sum[:])
}
