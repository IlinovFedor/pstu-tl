package usecases

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"uuid"

	"github.com/google/go-cmp/cmp"
	"github.com/xuri/excelize/v2"

	"github.com/IlinovFedor/pstu-tl/internal/models"
)

var update = flag.Bool("update", false, "update golden files")

func openXLSX(t *testing.T, path string) *excelize.File {
	t.Helper()
	f, err := excelize.OpenFile(path)
	if err != nil {
		t.Fatalf("cannot open %s: %v", path, err)
	}
	t.Cleanup(func() { _ = f.Close() })
	return f
}

func Test_parser_parseSheet_golden(t *testing.T) {
	file := openXLSX(t, "testdata/test1.xlsx")

	p := newParser()
	if err := p.parseSheet(file, models.TermSecond, uuid.Nil()); err != nil {
		t.Fatalf("parseSheet: %v", err)
	}

	got := dumpParser(t, p)
	golden := filepath.Join("testdata", "test1.golden")

	if *update {
		if err := os.WriteFile(golden, []byte(got), 0o644); err != nil {
			t.Fatalf("update golden: %v", err)
		}
	}

	want, err := os.ReadFile(golden)
	if err != nil {
		t.Fatalf("read golden (run with -update): %v", err)
	}
	if diff := cmp.Diff(string(want), got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func dumpParser(t *testing.T, p *parser) string {
	t.Helper()
	var b strings.Builder

	var lines []string
	for imp, m := range p.disciplines {
		lines = append(lines, fmt.Sprintf("%s\t%s", imp.Name(), m.Name()))
	}
	writeSection(&b, "Disciplines", lines)

	lines = nil
	for imp, m := range p.faculties {
		lines = append(lines, fmt.Sprintf("%s\t%s", imp.Name(), m.Name()))
	}
	writeSection(&b, "Faculties", lines)

	lines = nil
	for imp, m := range p.specials {
		lines = append(lines, fmt.Sprintf("%s\t%s", imp.Name(), m.Name()))
	}
	writeSection(&b, "Specializations", lines)

	lines = nil
	for imp, m := range p.groups {
		strType, err := m.GroupType().String()
		if err != nil {
			t.Fatalf("cannot get group type: %v", err)
		}
		lines = append(lines, fmt.Sprintf("%s %s-%d-%d%s",
			imp.faculty.Name(), imp.specialization.Name(),
			imp.YearOfEnrollment(), imp.SequenceNumber(), strType))
	}
	writeSection(&b, "Groups", lines)

	b.WriteString("Records\n")
	for i, r := range p.records {
		recType, err := r.RecordType().String()
		if err != nil {
			t.Fatalf("cannot get record type: %v", err)
		}
		fmt.Fprintf(&b, "%d) %s %s\n\tHours: %d\n\tStudents amount: %d\n\tGroups:",
			i, r.discipline.Name(), recType, r.GoalHours(), r.StudentsAmount())
		for _, sg := range r.subgroups {
			gt, err := sg.group.GroupType().String()
			if err != nil {
				t.Fatalf("cannot get group type: %v", err)
			}
			fmt.Fprintf(&b, " %s %s-%d-%d%s пг%d;",
				sg.group.faculty.Name(), sg.group.specialization.Name(),
				sg.group.YearOfEnrollment(), sg.group.SequenceNumber(), gt, sg.subgroup)
		}
		b.WriteString("\n")
	}
	return b.String()
}

func writeSection(b *strings.Builder, title string, lines []string) {
	sort.Strings(lines)
	b.WriteString(title + "\n")
	for _, l := range lines {
		b.WriteString(l + "\n")
	}
}

type cellErr struct {
	Cell  string
	Term  models.Term
	Value string
	Err   string
}

func Test_parser_parseSheet_errors(t *testing.T) {
	tests := []struct {
		name string
		file string
		want []cellErr
	}{
		{
			name: "test2",
			file: "testdata/test2.xlsx",
			want: []cellErr{
				{Cell: "B1", Term: 2, Value: "группа", Err: "cannot parse group cell"},
				{
					Cell:  "C1",
					Term:  2,
					Value: "число студентов",
					Err:   `cannot parse record students amount: strconv.Atoi: parsing "число студентов": invalid syntax`,
				},
				{
					Cell:  "D1",
					Term:  2,
					Value: "подгруппа",
					Err:   `cannot parse record subgroup number: strconv.Atoi: parsing "подгруппа": invalid syntax`,
				},
				{
					Cell:  "E1",
					Term:  2,
					Value: "лекция",
					Err:   "unknown record type: unknown record type",
				},
				{
					Cell:  "F1",
					Term:  2,
					Value: "часы",
					Err:   `cannot parse record hours: strconv.Atoi: parsing "часы": invalid syntax`,
				},
			},
		},
		{
			name: "test3",
			file: "testdata/test3.xlsx",
			want: []cellErr{
				{Cell: "A2", Term: 2, Err: "empty discipline"},
				{Cell: "A8", Term: 2, Err: "empty discipline"},
				{Cell: "B1", Term: 2, Value: "группа1", Err: "cannot parse group cell"},
				{Cell: "B3", Term: 2, Value: "группа2", Err: "cannot parse group cell"},
				{Cell: "B9", Term: 2, Err: "cannot parse group cell"},
				{
					Cell:  "C1",
					Term:  2,
					Value: "число студентов1",
					Err:   `cannot parse record students amount: strconv.Atoi: parsing "число студентов1": invalid syntax`,
				},
				{
					Cell: "C10",
					Term: 2,
					Err:  `cannot parse record students amount: strconv.Atoi: parsing "": invalid syntax`,
				},
				{
					Cell:  "C4",
					Term:  2,
					Value: "число студентов2",
					Err:   `cannot parse record students amount: strconv.Atoi: parsing "число студентов2": invalid syntax`,
				},
				{
					Cell:  "D1",
					Term:  2,
					Value: "подгруппа1",
					Err:   `cannot parse record subgroup number: strconv.Atoi: parsing "подгруппа1": invalid syntax`,
				},
				{
					Cell:  "D5",
					Term:  2,
					Value: "подгруппа2",
					Err:   `cannot parse record subgroup number: strconv.Atoi: parsing "подгруппа2": invalid syntax`,
				},
				{
					Cell:  "E1",
					Term:  2,
					Value: "лекция1",
					Err:   "unknown record type: unknown record type",
				},
				{Cell: "E12", Term: 2, Err: "unknown record type: unknown record type"},
				{
					Cell:  "E6",
					Term:  2,
					Value: "лекция2",
					Err:   "unknown record type: unknown record type",
				},
				{
					Cell:  "F1",
					Term:  2,
					Value: "часы1",
					Err:   `cannot parse record hours: strconv.Atoi: parsing "часы1": invalid syntax`,
				},
				{
					Cell: "F13",
					Term: 2,
					Err:  `cannot parse record hours: strconv.Atoi: parsing "  ": invalid syntax`,
				},
				{
					Cell:  "F7",
					Term:  2,
					Value: "часы2",
					Err:   `cannot parse record hours: strconv.Atoi: parsing "часы2": invalid syntax`,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := openXLSX(t, tt.file)

			err := newParser().parseSheet(file, models.TermSecond, uuid.Nil())

			var pe *models.ParseError
			if !errors.As(err, &pe) {
				t.Fatalf("want *models.ParseError, got %T: %v", err, err)
			}

			var got []cellErr
			for _, c := range pe.Cells() {
				got = append(got, cellErr{
					Cell:  c.Cell(),
					Term:  c.Term(),
					Value: c.Value(),
					Err:   c.Err().Error(),
				})
			}
			sort.Slice(got, func(i, j int) bool { return got[i].Cell < got[j].Cell })
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("cell errors mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
