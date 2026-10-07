package models

import (
	"strings"
	"time"
	"uuid"
)

// Faculty базовая модель факультета.
type Faculty struct {
	id        uuid.UUID
	name      string
	deletedAt *time.Time
}

func NewFaculty(id uuid.UUID, name string, deletedAt *time.Time) *Faculty {
	name = strings.ToLower(name)
	return &Faculty{id: id, name: strings.TrimSpace(name), deletedAt: deletedAt}
}

func (d *Faculty) SetId(id uuid.UUID) {
	d.id = id
}

func (d *Faculty) DeletedAt() *time.Time {
	return d.deletedAt
}

func (d *Faculty) SetDeletedAt(deletedAt *time.Time) {
	d.deletedAt = deletedAt
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
	newDeletedAt    *time.Time
}

func NewFacultyAction(id uuid.UUID, actorName string, revertsActionID *uuid.UUID, disciplineID uuid.UUID, newName string, newDeletedAt *time.Time) *FacultyAction {
	return &FacultyAction{id: id, actorName: actorName, revertsActionID: revertsActionID, disciplineID: disciplineID, newName: newName, newDeletedAt: newDeletedAt}
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

func (d FacultyAction) NewDeletedAt() *time.Time {
	return d.newDeletedAt
}
