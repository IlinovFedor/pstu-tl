package models

import (
	"uuid"
)

// Faculty базовая модель факультета.
type Faculty struct {
	id        uuid.UUID
	name      string
	isDeleted bool
}

func (d *Faculty) IsDeleted() bool {
	return d.isDeleted
}

func (d *Faculty) SetIsDeleted(isDeleted bool) {
	d.isDeleted = isDeleted
}

func (d *Faculty) SetName(name string) {
	d.name = name
}

func (d Faculty) Id() uuid.UUID {
	return d.id
}

func (d Faculty) Name() string {
	return d.name
}

func NewFaculty(name string) *Faculty {
	return &Faculty{
		id:   uuid.NewV7(),
		name: name}
}

// FacultyAction модель изменения дисциплины. Получается только из репозитория, append-only
type FacultyAction struct {
	id              uuid.UUID
	actorName       string
	revertsActionID *uuid.UUID
	disciplineID    uuid.UUID
	newName         string
	newIsDeleted    bool
}

func NewFacultyAction(id uuid.UUID, actorName string, revertsActionID *uuid.UUID, disciplineID uuid.UUID, newName string, newIsDeleted bool) *FacultyAction {
	return &FacultyAction{id: id, actorName: actorName, revertsActionID: revertsActionID, disciplineID: disciplineID, newName: newName, newIsDeleted: newIsDeleted}
}

func (d FacultyAction) Id() uuid.UUID {
	return d.id
}

func (d FacultyAction) ActorName() string {
	return d.actorName
}

func (d FacultyAction) RevertsActionID() *uuid.UUID {
	return d.revertsActionID
}

func (d FacultyAction) FacultyID() uuid.UUID {
	return d.disciplineID
}

func (d FacultyAction) NewName() string {
	return d.newName
}

func (d FacultyAction) NewIsDeleted() bool {
	return d.newIsDeleted
}
