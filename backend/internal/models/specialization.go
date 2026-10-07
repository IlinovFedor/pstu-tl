package models

import (
	"strings"
	"time"
	"uuid"
)

// Specialization базовая модель дисциплины.
type Specialization struct {
	id        uuid.UUID
	name      string
	deletedAt *time.Time
}

func NewSpecialization(id uuid.UUID, name string, deletedAt *time.Time) *Specialization {
	name = strings.ToLower(name)
	return &Specialization{id: id, name: strings.TrimSpace(name), deletedAt: deletedAt}
}

func (d *Specialization) SetId(id uuid.UUID) {
	d.id = id
}

func (d *Specialization) DeletedAt() *time.Time {
	return d.deletedAt
}

func (d *Specialization) SetDeletedAt(deletedAt *time.Time) {
	d.deletedAt = deletedAt
}

func (d *Specialization) SetName(name string) {
	name = strings.ToLower(name)
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
	newDeletedAt     *time.Time
}

func NewSpecializationAction(id uuid.UUID, actorName string, revertsActionID *uuid.UUID, SpecializationID uuid.UUID, newName string, newDeletedAt *time.Time) *SpecializationAction {
	return &SpecializationAction{id: id, actorName: actorName, revertsActionID: revertsActionID, specializationID: SpecializationID, newName: newName, newDeletedAt: newDeletedAt}
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

func (d SpecializationAction) NewDeletedAt() *time.Time {
	return d.newDeletedAt
}
