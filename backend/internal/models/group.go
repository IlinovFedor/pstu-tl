package models

import (
	"strings"
	"time"
	"uuid"
)

// GroupType тип группы (заоч, оч, асп)
type GroupType int

const (
	GroupTypeBachelorPartTime GroupType = iota
	GroupTypeBachelorPartTimeAccelerated
	GroupTypeBachelor
	GroupTypeSpecialistPartTime
	GroupTypeSpecialistPartTimeAccelerated
	GroupTypeSpecialist
	GroupTypePostgraduate
)

func NewGroupTypeFromString(s string) (GroupType, error) {
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)
	switch s {
	case "б":
		return GroupTypeBachelor, nil
	case "бз":
		return GroupTypeBachelorPartTime, nil
	case "бзу":
		return GroupTypeBachelorPartTimeAccelerated, nil
	case "с":
		return GroupTypeSpecialist, nil
	case "сз":
		return GroupTypeSpecialistPartTime, nil
	case "сзу":
		return GroupTypeSpecialistPartTimeAccelerated, nil
	case "м":
		return GroupTypePostgraduate, nil
	default:
		return 0, ErrUnexpectedGroupType
	}
}

func (g GroupType) String() (string, error) {
	switch g {
	case GroupTypeBachelor:
		return "б", nil
	case GroupTypeBachelorPartTime:
		return "бз", nil
	case GroupTypeBachelorPartTimeAccelerated:
		return "бзу", nil
	case GroupTypeSpecialist:
		return "с", nil
	case GroupTypeSpecialistPartTime:
		return "сз", nil
	case GroupTypeSpecialistPartTimeAccelerated:
		return "сзу", nil
	case GroupTypePostgraduate:
		return "м", nil
	default:
		return "", ErrUnexpectedGroupType
	}
}

// Group базовая модель группы.
type Group struct {
	id               uuid.UUID
	facultyID        uuid.UUID
	specializationID uuid.UUID
	yearOfEnrollment int
	sequenceNumber   int
	groupType        GroupType
	deletedAt        *time.Time
}

func NewGroup(id uuid.UUID, facultyID uuid.UUID, specializationID uuid.UUID, yearOfEnrollment int, sequenceNumber int, groupType GroupType, deletedAt *time.Time) *Group {
	return &Group{id: id, facultyID: facultyID, specializationID: specializationID, yearOfEnrollment: yearOfEnrollment, sequenceNumber: sequenceNumber, groupType: groupType, deletedAt: deletedAt}
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

func (g *Group) DeletedAt() *time.Time {
	return g.deletedAt
}

func (g *Group) SetDeletedAt(deletedAt *time.Time) {
	g.deletedAt = deletedAt
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
	newDeletedAt        *time.Time
}

func NewGroupAction(id uuid.UUID, actorName string, revertsActionID *uuid.UUID, groupId uuid.UUID, newFacultyId uuid.UUID, newSpecialization string, newYearOfEnrollment int, newSequenceNumber int, newGroupType GroupType, newStudentsAmount int, newDeletedAt *time.Time) *GroupAction {
	return &GroupAction{id: id, actorName: actorName, revertsActionID: revertsActionID, groupId: groupId, newFacultyId: newFacultyId, newSpecialization: newSpecialization, newYearOfEnrollment: newYearOfEnrollment, newSequenceNumber: newSequenceNumber, newGroupType: newGroupType, newStudentsAmount: newStudentsAmount, newDeletedAt: newDeletedAt}
}

func (g GroupAction) NewDeletedAt() *time.Time {
	return g.newDeletedAt
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
