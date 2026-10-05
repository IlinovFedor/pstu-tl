package models

import (
	"strings"
	"uuid"
)

// Discipline базовая модель дисциплины.
type Discipline struct {
	id        uuid.UUID
	name      string
	isDeleted bool
}

func NewDiscipline(id uuid.UUID, name string, isDeleted bool) *Discipline {
	return &Discipline{id: id, name: strings.TrimSpace(name), isDeleted: isDeleted}
}

func (d *Discipline) SetId(id uuid.UUID) {
	d.id = id
}

func (d *Discipline) IsDeleted() bool {
	return d.isDeleted
}

func (d *Discipline) SetIsDeleted(isDeleted bool) {
	d.isDeleted = isDeleted
}

func (d *Discipline) SetName(name string) {
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
	newIsDeleted    bool
}

func NewDisciplineAction(id uuid.UUID, actorName string, revertsActionID *uuid.UUID, disciplineID uuid.UUID, newName string, newIsDeleted bool) *DisciplineAction {
	return &DisciplineAction{id: id, actorName: actorName, revertsActionID: revertsActionID, disciplineID: disciplineID, newName: newName, newIsDeleted: newIsDeleted}
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

func (d DisciplineAction) NewIsDeleted() bool {
	return d.newIsDeleted
}
