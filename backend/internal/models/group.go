package models

import (
	"uuid"
)

// GroupType тип группы (заоч, оч, асп)
type GroupType int

const (
	GroupTypePartTime GroupType = iota
	GroupTypeUndergraduate
	GroupTypePostgraduate
)

// Group базовая модель группы.
type Group struct {
	id               uuid.UUID
	facultyID        uuid.UUID
	specializationID uuid.UUID
	yearOfEnrollment int
	sequenceNumber   int
	groupType        GroupType
	studentsAmount   int
	isDeleted        bool
}

func NewGroup(id uuid.UUID, facultyID uuid.UUID, specializationID uuid.UUID, yearOfEnrollment int, sequenceNumber int, groupType GroupType, studentsAmount int, isDeleted bool) *Group {
	return &Group{id: id, facultyID: facultyID, specializationID: specializationID, yearOfEnrollment: yearOfEnrollment, sequenceNumber: sequenceNumber, groupType: groupType, studentsAmount: studentsAmount, isDeleted: isDeleted}
}

func (g *Group) SetId(id uuid.UUID) {
	g.id = id
}

func (g *Group) SetSpecializationID(specializationID uuid.UUID) {
	g.specializationID = specializationID
}

func (g *Group) FacultyID() uuid.UUID {
	return g.facultyID
}

func (g *Group) SetFacultyID(facultyID uuid.UUID) {
	g.facultyID = facultyID
}

func (g *Group) Id() uuid.UUID {
	return g.id
}

func (g *Group) SpecializationID() uuid.UUID {
	return g.specializationID
}

func (g *Group) SetSpecialization(specializationID uuid.UUID) {
	g.specializationID = specializationID
}

func (g *Group) YearOfEnrollment() int {
	return g.yearOfEnrollment
}

func (g *Group) SetYearOfEnrollment(yearOfEnrollment int) {
	g.yearOfEnrollment = yearOfEnrollment
}

func (g *Group) SequenceNumber() int {
	return g.sequenceNumber
}

func (g *Group) SetSequenceNumber(sequenceNumber int) {
	g.sequenceNumber = sequenceNumber
}

func (g *Group) GroupType() GroupType {
	return g.groupType
}

func (g *Group) SetGroupType(groupType GroupType) {
	g.groupType = groupType
}

func (g *Group) StudentsAmount() int {
	return g.studentsAmount
}

func (g *Group) SetStudentsAmount(studentsAmount int) {
	g.studentsAmount = studentsAmount
}

func (g *Group) IsDeleted() bool {
	return g.isDeleted
}

func (g *Group) SetIsDeleted(isDeleted bool) {
	g.isDeleted = isDeleted
}

// GroupAction модель изменения группы. Получается только из репозитория, append-only
type GroupAction struct {
	id                  uuid.UUID
	actorName           string
	revertsActionID     *uuid.UUID
	groupId             uuid.UUID
	newFacultyId        uuid.UUID
	newSpecialization   string
	newYearOfEnrollment int
	newSequenceNumber   int
	newGroupType        GroupType
	newStudentsAmount   int
	newIsDeleted        bool
}

func NewGroupAction(id uuid.UUID, actorName string, revertsActionID *uuid.UUID, groupId uuid.UUID, newFacultyId uuid.UUID, newSpecialization string, newYearOfEnrollment int, newSequenceNumber int, newGroupType GroupType, newStudentsAmount int, newIsDeleted bool) *GroupAction {
	return &GroupAction{id: id, actorName: actorName, revertsActionID: revertsActionID, groupId: groupId, newFacultyId: newFacultyId, newSpecialization: newSpecialization, newYearOfEnrollment: newYearOfEnrollment, newSequenceNumber: newSequenceNumber, newGroupType: newGroupType, newStudentsAmount: newStudentsAmount, newIsDeleted: newIsDeleted}
}

func (g GroupAction) NewIsDeleted() bool {
	return g.newIsDeleted
}

func (g GroupAction) NewStudentsAmount() int {
	return g.newStudentsAmount
}

func (g GroupAction) NewGroupType() GroupType {
	return g.newGroupType
}

func (g GroupAction) NewSequenceNumber() int {
	return g.newSequenceNumber
}

func (g GroupAction) NewYearOfEnrollment() int {
	return g.newYearOfEnrollment
}

func (g GroupAction) NewSpecialization() string {
	return g.newSpecialization
}

func (g GroupAction) NewFacultyId() uuid.UUID {
	return g.newFacultyId
}

func (g GroupAction) GroupId() uuid.UUID {
	return g.groupId
}

func (g GroupAction) RevertsActionID() *uuid.UUID {
	return g.revertsActionID
}

func (g GroupAction) ActorName() string {
	return g.actorName
}

func (g GroupAction) Id() uuid.UUID {
	return g.id
}
