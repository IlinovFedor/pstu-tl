package models

import (
	"strings"
	"uuid"
)

// Specialization базовая модель дисциплины.
type Specialization struct {
	id        uuid.UUID
	name      string
	isDeleted bool
}

func NewSpecialization(id uuid.UUID, name string, isDeleted bool) *Specialization {
	return &Specialization{id: id, name: strings.TrimSpace(name), isDeleted: isDeleted}
}

func (d *Specialization) SetId(id uuid.UUID) {
	d.id = id
}

func (d *Specialization) IsDeleted() bool {
	return d.isDeleted
}

func (d *Specialization) SetIsDeleted(isDeleted bool) {
	d.isDeleted = isDeleted
}

func (d *Specialization) SetName(name string) {
	d.name = strings.TrimSpace(name)
}

func (d Specialization) Id() uuid.UUID {
	return d.id
}

func (d Specialization) Name() string {
	return d.name
}

// SpecializationAction модель изменения дисциплины. Получается только из репозитория, append-only
type SpecializationAction struct {
	id               uuid.UUID
	actorName        string
	revertsActionID  *uuid.UUID
	specializationID uuid.UUID
	newName          string
	newIsDeleted     bool
}

func NewSpecializationAction(id uuid.UUID, actorName string, revertsActionID *uuid.UUID, SpecializationID uuid.UUID, newName string, newIsDeleted bool) *SpecializationAction {
	return &SpecializationAction{id: id, actorName: actorName, revertsActionID: revertsActionID, specializationID: SpecializationID, newName: newName, newIsDeleted: newIsDeleted}
}

func (d SpecializationAction) Id() uuid.UUID {
	return d.id
}

func (d SpecializationAction) ActorName() string {
	return d.actorName
}

func (d SpecializationAction) RevertsActionID() *uuid.UUID {
	return d.revertsActionID
}

func (d SpecializationAction) SpecializationID() uuid.UUID {
	return d.specializationID
}

func (d SpecializationAction) NewName() string {
	return d.newName
}

func (d SpecializationAction) NewIsDeleted() bool {
	return d.newIsDeleted
}
