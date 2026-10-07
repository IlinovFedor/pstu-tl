package usecases

import (
	"context"
	"errors"
	"testing"
	"uuid"

	"github.com/xuri/excelize/v2"

	"github.com/IlinovFedor/pstu-tl/internal/models"
)

type fakeTransactor struct{ calls int }

func (f *fakeTransactor) WithTx(ctx context.Context, fn func(ctx context.Context) error) error {
	f.calls++
	return fn(ctx)
}

type fakeDisciplineRepo struct {
	DisciplineRepository
	calls int
	ids   map[string]uuid.UUID
}

func (f *fakeDisciplineRepo) UpsertDisciplines(_ context.Context, ds []*models.Discipline) ([]models.Discipline, error) {
	f.calls++
	res := make([]models.Discipline, 0, len(ds))
	for _, d := range ds {
		id, ok := f.ids[d.Name()]
		if !ok {
			id = uuid.NewV7()
			f.ids[d.Name()] = id
		}
		res = append(res, *models.NewDiscipline(id, d.Name(), false))
	}
	return res, nil
}

type fakeFacultyRepo struct {
	FacultyRepository
	calls int
	ids   map[string]uuid.UUID
}

func (f *fakeFacultyRepo) UpsertFaculties(_ context.Context, fs []*models.Faculty) ([]models.Faculty, error) {
	f.calls++
	res := make([]models.Faculty, 0, len(fs))
	for _, fc := range fs {
		id, ok := f.ids[fc.Name()]
		if !ok {
			id = uuid.NewV7()
			f.ids[fc.Name()] = id
		}
		res = append(res, *models.NewFaculty(id, fc.Name(), false))
	}
	return res, nil
}

type fakeSpecializationRepo struct {
	SpecializationRepository
	calls int
	ids   map[string]uuid.UUID
}

func (f *fakeSpecializationRepo) UpsertSpecializations(_ context.Context, ss []*models.Specialization) ([]models.Specialization, error) {
	f.calls++
	res := make([]models.Specialization, 0, len(ss))
	for _, s := range ss {
		id, ok := f.ids[s.Name()]
		if !ok {
			id = uuid.NewV7()
			f.ids[s.Name()] = id
		}
		res = append(res, *models.NewSpecialization(id, s.Name(), false))
	}
	return res, nil
}

type groupKey struct {
	facultyID, specializationID uuid.UUID
	year, seq                   int
	groupType                   models.GroupType
}

func keyOfGroup(g *models.Group) groupKey {
	return groupKey{g.FacultyID(), g.SpecializationID(), g.YearOfEnrollment(), g.SequenceNumber(), g.GroupType()}
}

type fakeGroupRepo struct {
	GroupRepository
	calls int
	ids   map[groupKey]uuid.UUID
	got   []models.Group
}

func (f *fakeGroupRepo) UpsertGroups(_ context.Context, gs []*models.Group) ([]models.Group, error) {
	f.calls++
	res := make([]models.Group, 0, len(gs))
	for _, g := range gs {
		f.got = append(f.got, *g)
		k := keyOfGroup(g)
		id, ok := f.ids[k]
		if !ok {
			id = uuid.NewV7()
			f.ids[k] = id
		}
		res = append(res, *models.NewGroup(id, k.facultyID, k.specializationID, k.year, k.seq, k.groupType, false))
	}
	return res, nil
}

type fakeFileImportRepo struct {
	FileImportRepository
	byHashCalls int
	existing    map[string]*models.FileImport
	inserted    []models.FileImport
}

func (f *fakeFileImportRepo) GetFileImportByHash(_ context.Context, _ int, hash string) (*models.FileImport, error) {
	f.byHashCalls++
	if fi, ok := f.existing[hash]; ok {
		return fi, nil
	}
	return nil, models.ErrUnknownFile
}

func (f *fakeFileImportRepo) InsertFileImport(_ context.Context, fi models.FileImport) error {
	f.inserted = append(f.inserted, fi)
	return nil
}

type fakeRecordsRepo struct {
	RecordsRepository
	calls    int
	inserted []models.Record
}

func (f *fakeRecordsRepo) InsertRecords(_ context.Context, rs []*models.Record) error {
	f.calls++
	for _, r := range rs {
		f.inserted = append(f.inserted, *r)
	}
	return nil
}

type fakeStorage struct {
	files map[uuid.UUID][]byte
}

func (f *fakeStorage) CreateFile(id uuid.UUID, data []byte) error {
	f.files[id] = data
	return nil
}

func (f *fakeStorage) GetFile(id uuid.UUID) ([]byte, error) { return f.files[id], nil }

func (f *fakeStorage) DeleteFile(id uuid.UUID) error {
	delete(f.files, id)
	return nil
}

type importFixture struct {
	svc         *FileImportService
	tx          *fakeTransactor
	files       *fakeFileImportRepo
	disciplines *fakeDisciplineRepo
	faculties   *fakeFacultyRepo
	specials    *fakeSpecializationRepo
	groups      *fakeGroupRepo
	records     *fakeRecordsRepo
	storage     *fakeStorage
}

func newImportFixture() *importFixture {
	f := &importFixture{
		tx:          &fakeTransactor{},
		files:       &fakeFileImportRepo{existing: map[string]*models.FileImport{}},
		disciplines: &fakeDisciplineRepo{ids: map[string]uuid.UUID{}},
		faculties:   &fakeFacultyRepo{ids: map[string]uuid.UUID{}},
		specials:    &fakeSpecializationRepo{ids: map[string]uuid.UUID{}},
		groups:      &fakeGroupRepo{ids: map[groupKey]uuid.UUID{}},
		records:     &fakeRecordsRepo{},
		storage:     &fakeStorage{files: map[uuid.UUID][]byte{}},
	}
	f.svc = &FileImportService{
		transactor:          f.tx,
		fileRepo:            f.files,
		disciplineRepo:      f.disciplines,
		facultyRepo:         f.faculties,
		specializationsRepo: f.specials,
		groupRepo:           f.groups,
		recordsRepo:         f.records,
		storage:             f.storage,
	}
	return f
}

func (f *importFixture) assertNoWrites(t *testing.T) {
	t.Helper()
	if f.tx.calls != 0 || f.files.byHashCalls != 0 || len(f.files.inserted) != 0 ||
		f.disciplines.calls != 0 || f.faculties.calls != 0 || f.specials.calls != 0 ||
		f.groups.calls != 0 || f.records.calls != 0 || len(f.storage.files) != 0 {
		t.Fatalf("expected no repo/storage calls, got tx=%d byHash=%d files=%d disc=%d fac=%d spec=%d groups=%d records=%d storage=%d",
			f.tx.calls, f.files.byHashCalls, len(f.files.inserted), f.disciplines.calls, f.faculties.calls,
			f.specials.calls, f.groups.calls, f.records.calls, len(f.storage.files))
	}
}

// buildXLSX собирает xlsx в памяти: ключ - имя листа, значение - строки
func buildXLSX(t *testing.T, sheets map[string][][]string) []byte {
	t.Helper()
	file := excelize.NewFile()
	defer func() { _ = file.Close() }()
	for name, rows := range sheets {
		if _, err := file.NewSheet(name); err != nil {
			t.Fatalf("new sheet: %v", err)
		}
		for i, row := range rows {
			cellName, _ := excelize.CoordinatesToCellName(1, i+1)
			vals := make([]any, len(row))
			for j, v := range row {
				vals[j] = v
			}
			if err := file.SetSheetRow(name, cellName, &vals); err != nil {
				t.Fatalf("set row: %v", err)
			}
		}
	}
	if err := file.DeleteSheet("Sheet1"); err != nil {
		t.Fatalf("delete default sheet: %v", err)
	}
	buf, err := file.WriteToBuffer()
	if err != nil {
		t.Fatalf("write xlsx: %v", err)
	}
	return buf.Bytes()
}

func TestImportFile_ParseErrorsFromBothSheets_NoWrites(t *testing.T) {
	f := newImportFixture()
	data := buildXLSX(t, map[string][][]string{
		models.TermFirst.String(): {
			{"Математика", "ЭТФ АСУ-22-1б", "25", "", "Лек", "36"},
			{"Физика", "ЭТФ АСУ-22-1б", "x", "", "Лек", "36"},
		},
		models.TermSecond.String(): {
			{"Химия", "ЭТФ АСУ-22-1б", "25", "", "Лекция", "36"},
		},
	})

	err := f.svc.ImportFile(context.Background(), 2025, data, "plan.xlsx")

	var parseErr *models.ParseError
	if !errors.As(err, &parseErr) {
		t.Fatalf("want *models.ParseError, got %v", err)
	}
	got := map[string]models.Term{}
	for _, c := range parseErr.Cells() {
		got[c.Cell()] = c.Term()
	}
	want := map[string]models.Term{"C2": models.TermFirst, "E1": models.TermSecond}
	if len(got) != len(want) || got["C2"] != want["C2"] || got["E1"] != want["E1"] {
		t.Fatalf("cells: want %v, got %v", want, got)
	}
	f.assertNoWrites(t)
}

// validSheets: 2 дисциплины, 2 факультета, 2 специальности, 3 группы (АСУ-22-1 встречается дважды)
func validSheets() map[string][][]string {
	return map[string][][]string{
		models.TermFirst.String(): {
			{"Математика", "ЭТФ АСУ-22-1,2б", "25", "", "Лек", "36"},
			{"Математика", "ЭТФ АСУ-22-1б", "12", "1", "Пр", "18"},
		},
		models.TermSecond.String(): {
			{"Физика", "ФПММ ПМИ-23-1б", "20", "", "Лаб", "10"},
		},
	}
}

type wantGroup struct {
	faculty, specialization string
	year, seq               int
	subgroup                models.SubgroupType
}

type wantRecord struct {
	discipline     string
	term           models.Term
	recordType     models.RecordType
	studentsAmount int
	goalHours      int
	groups         []wantGroup
}

func validSheetsWant() []wantRecord {
	return []wantRecord{
		{"Математика", models.TermFirst, models.RecordTypeLecture, 25, 36, []wantGroup{
			{"ЭТФ", "АСУ", 22, 1, models.SubgroupTypeBoth},
			{"ЭТФ", "АСУ", 22, 2, models.SubgroupTypeBoth},
		}},
		{"Математика", models.TermFirst, models.RecordTypePractice, 12, 18, []wantGroup{
			{"ЭТФ", "АСУ", 22, 1, models.SubgroupTypeFirst},
		}},
		{"Физика", models.TermSecond, models.RecordTypeLab, 20, 10, []wantGroup{
			{"ФПММ", "ПМИ", 23, 1, models.SubgroupTypeBoth},
		}},
	}
}

func disciplineName(s string) string { return models.NewDiscipline(uuid.Nil(), s, false).Name() }
func facultyName(s string) string    { return models.NewFaculty(uuid.Nil(), s, false).Name() }
func specializationName(s string) string {
	return models.NewSpecialization(uuid.Nil(), s, false).Name()
}

// assertRecords проверяет, что записи ссылаются на ID, выданные репозиториями
func (f *importFixture) assertRecords(t *testing.T, fileID uuid.UUID, want []wantRecord) {
	t.Helper()
	if len(f.records.inserted) != len(want) {
		t.Fatalf("records: want %d, got %d", len(want), len(f.records.inserted))
	}
	recordIDs := map[uuid.UUID]bool{}
	for _, w := range want {
		found := false
		for _, r := range f.records.inserted {
			if r.Term() != w.term || r.RecordType() != w.recordType || r.StudentsAmount() != w.studentsAmount ||
				r.GoalHours() != w.goalHours || r.DisciplineId() != f.disciplines.ids[disciplineName(w.discipline)] ||
				len(r.Groups()) != len(w.groups) {
				continue
			}
			match := true
			for _, wg := range w.groups {
				groupID := f.groups.ids[groupKey{
					facultyID:        f.faculties.ids[facultyName(wg.faculty)],
					specializationID: f.specials.ids[specializationName(wg.specialization)],
					year:             wg.year,
					seq:              wg.seq,
					groupType:        models.GroupTypeBachelor,
				}]
				hasGroup := false
				for _, g := range r.Groups() {
					if g.GroupID() == groupID && g.Subgroup() == wg.subgroup {
						hasGroup = true
					}
				}
				match = match && hasGroup
			}
			if !match {
				continue
			}
			found = true
			if r.ImportFileID() != fileID {
				t.Errorf("record %+v: importFileID %v, want %v", w, r.ImportFileID(), fileID)
			}
			if r.Id() == uuid.Nil() || recordIDs[r.Id()] {
				t.Errorf("record %+v: id %v is nil or duplicated", w, r.Id())
			}
			recordIDs[r.Id()] = true
		}
		if !found {
			t.Errorf("record not found: %+v", w)
		}
	}
}

func TestImportFile_Valid_UpsertsGraphInBatches(t *testing.T) {
	f := newImportFixture()
	data := buildXLSX(t, validSheets())

	if err := f.svc.ImportFile(context.Background(), 2025, data, "plan.xlsx"); err != nil {
		t.Fatalf("ImportFile: %v", err)
	}

	if f.tx.calls != 1 {
		t.Errorf("tx calls: want 1, got %d", f.tx.calls)
	}
	if f.disciplines.calls != 1 || f.faculties.calls != 1 || f.specials.calls != 1 ||
		f.groups.calls != 1 || f.records.calls != 1 {
		t.Errorf("each batch must be called once: disc=%d fac=%d spec=%d groups=%d records=%d",
			f.disciplines.calls, f.faculties.calls, f.specials.calls, f.groups.calls, f.records.calls)
	}
	if len(f.disciplines.ids) != 2 || len(f.faculties.ids) != 2 || len(f.specials.ids) != 2 || len(f.groups.got) != 3 {
		t.Errorf("deduplication: disc=%d fac=%d spec=%d groups=%d",
			len(f.disciplines.ids), len(f.faculties.ids), len(f.specials.ids), len(f.groups.got))
	}
	for _, g := range f.groups.got {
		if g.FacultyID() == uuid.Nil() || g.SpecializationID() == uuid.Nil() {
			t.Errorf("group passed to upsert without faculty/specialization id: %+v", g)
		}
	}

	if len(f.files.inserted) != 1 {
		t.Fatalf("file imports: want 1, got %d", len(f.files.inserted))
	}
	fi := f.files.inserted[0]
	if fi.Id() == uuid.Nil() || fi.YearName() != 2025 || fi.Name() != "plan.xlsx" || fi.Hash() == "" {
		t.Errorf("unexpected file import: id=%v year=%d name=%q hash=%q", fi.Id(), fi.YearName(), fi.Name(), fi.Hash())
	}
	if string(f.storage.files[fi.Id()]) != string(data) {
		t.Errorf("storage must contain file data under file id")
	}

	f.assertRecords(t, fi.Id(), validSheetsWant())
}

func TestImportFile_ExistingRows_KeepIDs(t *testing.T) {
	f := newImportFixture()
	disciplineID, facultyID, specializationID, groupID := uuid.NewV7(), uuid.NewV7(), uuid.NewV7(), uuid.NewV7()
	f.disciplines.ids[disciplineName("Математика")] = disciplineID
	f.faculties.ids[facultyName("ЭТФ")] = facultyID
	f.specials.ids[specializationName("АСУ")] = specializationID
	existingGroup := groupKey{facultyID, specializationID, 22, 1, models.GroupTypeBachelor}
	f.groups.ids[existingGroup] = groupID

	if err := f.svc.ImportFile(context.Background(), 2025, buildXLSX(t, validSheets()), "plan.xlsx"); err != nil {
		t.Fatalf("ImportFile: %v", err)
	}

	if f.disciplines.ids[disciplineName("Математика")] != disciplineID ||
		f.faculties.ids[facultyName("ЭТФ")] != facultyID ||
		f.specials.ids[specializationName("АСУ")] != specializationID ||
		f.groups.ids[existingGroup] != groupID {
		t.Fatalf("existing ids must not change")
	}
	if len(f.disciplines.ids) != 2 || len(f.faculties.ids) != 2 || len(f.specials.ids) != 2 || len(f.groups.ids) != 3 {
		t.Errorf("existing rows must not be duplicated: disc=%d fac=%d spec=%d groups=%d",
			len(f.disciplines.ids), len(f.faculties.ids), len(f.specials.ids), len(f.groups.ids))
	}
	f.assertRecords(t, f.files.inserted[0].Id(), validSheetsWant())
}

func TestImportFile_DuplicateHash_NoWrites(t *testing.T) {
	first := newImportFixture()
	if err := first.svc.ImportFile(context.Background(), 2025, buildXLSX(t, validSheets()), "plan.xlsx"); err != nil {
		t.Fatalf("first ImportFile: %v", err)
	}
	existing := first.files.inserted[0]

	f := newImportFixture()
	f.files.existing[existing.Hash()] = &existing
	reordered := validSheets()
	autumn := reordered[models.TermFirst.String()]
	autumn[0], autumn[1] = autumn[1], autumn[0]

	err := f.svc.ImportFile(context.Background(), 2025, buildXLSX(t, reordered), "other-name.xlsx")

	if !errors.Is(err, models.ErrDuplicateFileData) {
		t.Fatalf("want ErrDuplicateFileData, got %v", err)
	}
	if f.files.byHashCalls != 1 {
		t.Errorf("hash lookup calls: want 1, got %d", f.files.byHashCalls)
	}
	if len(f.files.inserted) != 0 || f.disciplines.calls != 0 || f.faculties.calls != 0 ||
		f.specials.calls != 0 || f.groups.calls != 0 || f.records.calls != 0 || len(f.storage.files) != 0 {
		t.Errorf("duplicate must not write anything")
	}
}

func TestImportFile_InvalidFileData_NoWrites(t *testing.T) {
	tests := []struct {
		name   string
		sheets map[string][][]string
	}{
		{name: "no known sheets", sheets: map[string][][]string{
			"Лист1": {{"Математика", "ЭТФ АСУ-22-1б", "25", "", "Лек", "36"}},
		}},
		{name: "no records", sheets: map[string][][]string{
			models.TermFirst.String(): {{"короткая строка"}},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := newImportFixture()

			err := f.svc.ImportFile(context.Background(), 2025, buildXLSX(t, tt.sheets), "plan.xlsx")

			if !errors.Is(err, models.ErrInvalidFileData) {
				t.Fatalf("want ErrInvalidFileData, got %v", err)
			}
			f.assertNoWrites(t)
		})
	}
}
