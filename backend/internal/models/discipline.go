package models

import (
	"strings"
	"time"
	"uuid"
)

// Discipline базовая модель дисциплины.
type Discipline struct {
	id        uuid.UUID
	name      string
	deletedAt *time.Time
}

func NewDiscipline(id uuid.UUID, name string, deletedAt *time.Time) *Discipline {
	name = strings.ToLower(name)
	return &Discipline{id: id, name: strings.TrimSpace(name), deletedAt: deletedAt}
}

func (d *Discipline) SetId(id uuid.UUID) {
	d.id = id
}

func (d *Discipline) DeletedAt() *time.Time {
	return d.deletedAt
}

func (d *Discipline) SetDeletedAt(deletedAt *time.Time) {
	d.deletedAt = deletedAt
}

func (d *Discipline) SetName(name string) {
	name = strings.ToLower(name)
	d.name = strings.TrimSpace(name)
}

func (d Discipline) Id() uuid.UUID {
	return d.id
}

func (d Discipline) Name() string {
	return d.name
}

// DisciplineAction модель изменения дисциплины. Получается только из репозитория, append-only
type DisciplineAction struct {
	id              uuid.UUID
	actorName       string
	revertsActionID *uuid.UUID
	disciplineID    uuid.UUID
	newName         string
	newDeletedAt    *time.Time
}

func NewDisciplineAction(id uuid.UUID, actorName string, revertsActionID *uuid.UUID, disciplineID uuid.UUID, newName string, newDeletedAt *time.Time) *DisciplineAction {
	return &DisciplineAction{id: id, actorName: actorName, revertsActionID: revertsActionID, disciplineID: disciplineID, newName: newName, newDeletedAt: newDeletedAt}
}

func (d DisciplineAction) Id() uuid.UUID {
	return d.id
}

func (d DisciplineAction) ActorName() string {
	return d.actorName
}

func (d DisciplineAction) RevertsActionID() *uuid.UUID {
	return d.revertsActionID
}

func (d DisciplineAction) DisciplineID() uuid.UUID {
	return d.disciplineID
}

func (d DisciplineAction) NewName() string {
	return d.newName
}

func (d DisciplineAction) NewDeletedAt() *time.Time {
	return d.newDeletedAt
}
