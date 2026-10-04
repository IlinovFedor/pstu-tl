package models

import (
	"uuid"
)

// RecordType тип занятия (лекция, практика, ...)
type RecordType int

const (
	RecordTypeLecture  RecordType = iota // Лек
	RecordTypePractice                   // Пр
	RecordTypeLab                        // Лаб
	RecordTypeKsr                        // КСР

	RecordTypeCourseProject // КП
	RecordTypeCourseWork    // КР
	RecordTypeControlWork   // КонтрРаб

	RecordTypePass       // Зач
	RecordTypeGradedPass // ДифЗач
	RecordTypeExam       // Экз
	RecordTypeStateExam  // ГосЭкз
	RecordTypeGek        // ГЭК

	RecordTypeThesisSupervision // РукВКР
	RecordTypeReviewing         // Рец

	RecordTypeStudyPractice      // УчПр
	RecordTypeWorkPractice       // ПрПр
	RecordTypePreDiplomaPractice // ПредПр
)

// SubgroupType тип подгруппы: 0 - обе, 1 - первая, 2 - вторая
type SubgroupType int

const (
	SubgroupTypeBoth SubgroupType = iota
	SubgroupTypeFirst
	SubgroupTypeSecond
)

// GroupsSubgroups пара группа - номер подгруппы
type GroupsSubgroups struct {
	subgroupId             uuid.UUID
	subgroupSequenceNumber SubgroupType
}

type Term int

const (
	TermFirst Term = iota + 1
	TermSecond
)

// TeacherHours тройка: айди преподавателя - количество часов - является ли доп нагрузкой
type TeacherHours struct {
	teacherId uuid.UUID
	hours     int
	isExtra   bool
}

// Record базовая модель дисциплины
type Record struct {
	id            uuid.UUID
	importFileID  uuid.UUID
	groups        []GroupsSubgroups
	disciplineId  uuid.UUID
	term          Term
	recordType    RecordType
	goalHours     int
	teachersHours []TeacherHours
	isError       bool
	isDeleted     bool
}

func NewRecord(importFileID uuid.UUID, groups []GroupsSubgroups, disciplineId uuid.UUID, term Term, recordType RecordType, goalHours int) *Record {
	return &Record{id: uuid.NewV7(), importFileID: importFileID, groups: groups, disciplineId: disciplineId, term: term, recordType: recordType, goalHours: goalHours}
}

func (r *Record) SetIsError(isError bool) {
	r.isError = isError
}

func (r *Record) SetTeachers(teachersHours []TeacherHours) {
	r.teachersHours = teachersHours
}

func (r Record) Id() uuid.UUID {
	return r.id
}

func (r Record) ImportFileID() uuid.UUID {
	return r.importFileID
}

func (r Record) Groups() []GroupsSubgroups {
	return r.groups
}

func (r Record) DisciplineId() uuid.UUID {
	return r.disciplineId
}

func (r Record) Term() Term {
	return r.term
}

func (r Record) RecordType() RecordType {
	return r.recordType
}

func (r Record) GoalHours() int {
	return r.goalHours
}

func (r Record) CurrentHours() int {
	var currentHours int
	for _, teacher := range r.teachersHours {
		currentHours += teacher.hours
	}
	return currentHours
}

func (r Record) Teachers() []TeacherHours {
	return r.teachersHours
}

func (r Record) IsError() bool {
	return r.isError
}

// RecordAction модель изменения записи. Получается только из репозитория, append-only
type RecordAction struct {
	id              uuid.UUID
	actorName       string
	revertsActionID *uuid.UUID
	recordId        uuid.UUID
	teacherId       uuid.UUID
	newHours        *int
	newIsExtra      bool
	newIsDeleted    bool
}

func (r RecordAction) Id() uuid.UUID {
	return r.id
}

func (r RecordAction) ActorName() string {
	return r.actorName
}

func (r RecordAction) RevertsActionID() *uuid.UUID {
	return r.revertsActionID
}

func (r RecordAction) RecordId() uuid.UUID {
	return r.recordId
}

func (r RecordAction) TeacherId() uuid.UUID {
	return r.teacherId
}

func (r RecordAction) NewHours() *int {
	return r.newHours
}

func (r RecordAction) NewIsExtra() bool {
	return r.newIsExtra
}

func (r RecordAction) NewIsDeleted() bool {
	return r.newIsDeleted
}

func NewRecordAction(id uuid.UUID, actorName string, revertsActionID *uuid.UUID, recordId uuid.UUID, teacherId uuid.UUID, newHours *int, newIsExtra bool, newIsDeleted bool) *RecordAction {
	return &RecordAction{id: id, actorName: actorName, revertsActionID: revertsActionID, recordId: recordId, teacherId: teacherId, newHours: newHours, newIsExtra: newIsExtra, newIsDeleted: newIsDeleted}
}
