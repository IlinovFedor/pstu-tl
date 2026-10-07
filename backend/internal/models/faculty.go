package models

import (
	"strings"
	"uuid"
)

// Faculty базовая модель факультета.
type Faculty struct {
	id        uuid.UUID
	name      string
	isDeleted bool
}

func NewFaculty(id uuid.UUID, name string, isDeleted bool) *Faculty {
	name = strings.ToLower(name)
	return &Faculty{id: id, name: strings.TrimSpace(name), isDeleted: isDeleted}
}

func (d *Faculty) SetId(id uuid.UUID) {
	d.id = id
}

func (d *Faculty) IsDeleted() bool {
	return d.isDeleted
}

func (d *Faculty) SetIsDeleted(isDeleted bool) {
	d.isDeleted = isDeleted
}

func (d *Faculty) SetName(name string) {
	name = strings.ToLower(name)
	d.name = strings.TrimSpace(name)
}

func (d Faculty) Id() uuid.UUID {
	return d.id
}

func (d Faculty) Name() string {
	return d.name
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
