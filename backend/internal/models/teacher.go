package models

import (
	"strings"
	"uuid"
)

type TeacherJob int

const (
	TeacherJobAssistant TeacherJob = iota
	TeacherJobAssociateProfessor
	TeacherJobDepartmentHead
	TeacherJovProfessor
	TeacherJobSeniorLecturer
)

type TeacherWage float32

const (
	WageQuarter TeacherWage = 0.25
	WageHalf    TeacherWage = 0.5
	WageFull    TeacherWage = 1.0
)

type Teacher struct {
	id        uuid.UUID
	name      string
	job       TeacherJob
	wage      TeacherWage
	isDeleted bool
}

func NewTeacher(id uuid.UUID, name string, job TeacherJob, wage TeacherWage, isDeleted bool) *Teacher {
	name = strings.ToLower(name)
	return &Teacher{id: id, name: strings.TrimSpace(name), job: job, wage: wage, isDeleted: isDeleted}
}

func (t *Teacher) IsDeleted() bool {
	return t.isDeleted
}

func (t *Teacher) SetIsDeleted(isDeleted bool) {
	t.isDeleted = isDeleted
}

func (t *Teacher) SetName(name string) {
	name = strings.ToLower(name)
	t.name = strings.TrimSpace(name)
}

func (t *Teacher) SetJob(job TeacherJob) {
	t.job = job
}

func (t *Teacher) SetWage(wage TeacherWage) {
	t.wage = wage
}

func (t Teacher) Id() uuid.UUID {
	return t.id
}

func (t Teacher) Name() string {
	return t.name
}

func (t Teacher) Job() TeacherJob {
	return t.job
}

func (t Teacher) Wage() TeacherWage {
	return t.wage
}

// TeacherAction модель изменения преподавателя. Получается только из репозитория, append-only
type TeacherAction struct {
	id              uuid.UUID
	actorName       string
	revertsActionID *uuid.UUID
	teacherID       uuid.UUID
	newName         string
	newJob          TeacherJob
	newWage         TeacherWage
	newIsDeleted    bool
}

func (t TeacherAction) NewIsDeleted() bool {
	return t.newIsDeleted
}

func (t TeacherAction) NewWage() TeacherWage {
	return t.newWage
}

func (t TeacherAction) NewJob() TeacherJob {
	return t.newJob
}

func (t TeacherAction) NewName() string {
	return t.newName
}

func (t TeacherAction) TeacherID() uuid.UUID {
	return t.teacherID
}

func (t TeacherAction) RevertsActionID() *uuid.UUID {
	return t.revertsActionID
}

func (t TeacherAction) ActorName() string {
	return t.actorName
}

func (t TeacherAction) Id() uuid.UUID {
	return t.id
}

func NewTeacherAction(id uuid.UUID, actorName string, revertsActionID *uuid.UUID, teacherID uuid.UUID, newName string, newJob TeacherJob, newWage TeacherWage, newIsDeleted bool) *TeacherAction {
	return &TeacherAction{id: id, actorName: actorName, revertsActionID: revertsActionID, teacherID: teacherID, newName: newName, newJob: newJob, newWage: newWage, newIsDeleted: newIsDeleted}
}
