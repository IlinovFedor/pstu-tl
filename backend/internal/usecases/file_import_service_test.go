package usecases

import (
	"bytes"
	"fmt"
	"os"
	"testing"
	"uuid"

	"github.com/IlinovFedor/pstu-tl/internal/models"
	"github.com/xuri/excelize/v2"
)

// TODO: finish test
func Test_parser_parseSheet(t *testing.T) {
	data, err := os.ReadFile("testdata/test1.xlsx")
	if err != nil {
		t.Fatalf("cannot read file: %v", err)
	}
	file, err := excelize.OpenReader(bytes.NewReader(data))
	if err != nil {
		t.Fatalf("cannot open xlsx reader: %v", err)
	}

	p := newParser()
	if p.parseSheet(file, models.TermSecond, uuid.Nil()) != nil {
		t.Errorf("cannot parse sheet %v", err)
	}

	fmt.Println("Disciplines")
	for imported, model := range p.disciplines {
		fmt.Printf("%s\t%s\n", imported.Name(), model.Name())
	}
	fmt.Println("Faculties")
	for imported, model := range p.faculties {
		fmt.Printf("%s\t%s\n", imported.Name(), model.Name())
	}
	fmt.Println("Specializations")
	for imported, model := range p.specials {
		fmt.Printf("%s\t%s\n", imported.Name(), model.Name())
	}
	fmt.Println("Groups")
	for imported, model := range p.groups {
		var strType string
		strType, err = model.GroupType().String()
		if err != nil {
			t.Errorf("cannot get group type: %v", err)
		}
		fmt.Printf("%s %s-%d-%d%s\n",
			imported.faculty.Name(),
			imported.specialization.Name(),
			imported.YearOfEnrollment(),
			imported.SequenceNumber(),
			strType)
	}

	for i, record := range p.records {
		var recType string
		recType, err = record.RecordType().String()
		if err != nil {
			t.Errorf("cannot get record type: %v", err)
		}
		fmt.Printf("%d) %s %s\n\tHours: %d\n\tStudents amount: %d\n", i, record.discipline.Name(), recType, record.GoalHours(), record.StudentsAmount())
		fmt.Printf("\tGroups: ")
		for _, subgroup := range record.subgroups {
			var strType string
			strType, err = subgroup.group.GroupType().String()
			if err != nil {
				t.Errorf("cannot get group type: %v", err)
			}
			fmt.Printf("%s %s-%d-%d%s пг%d ",
				subgroup.group.faculty.Name(),
				subgroup.group.specialization.Name(),
				subgroup.group.YearOfEnrollment(),
				subgroup.group.SequenceNumber(),
				strType,
				subgroup.subgroup)
		}
		fmt.Printf("\n")
	}
}
