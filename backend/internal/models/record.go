package models

import (
	"fmt"
	"strings"
	"time"
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

func NewRecordType(name string) (RecordType, error) {
	name = strings.ToLower(name)
	name = strings.TrimSpace(name)
	switch name {
	case "лек":
		return RecordTypeLecture, nil
	case "пр":
		return RecordTypePractice, nil
	case "лаб":
		return RecordTypeLab, nil
	case "кср":
		return RecordTypeKsr, nil
	case "кп":
		return RecordTypeCourseProject, nil
	case "кр":
		return RecordTypeCourseWork, nil
	case "контрраб":
		return RecordTypeControlWork, nil
	case "зач":
		return RecordTypePass, nil
	case "дифзач":
		return RecordTypeGradedPass, nil
	case "экз":
		return RecordTypeExam, nil
	case "госэкз":
		return RecordTypeStateExam, nil
	case "гэк":
		return RecordTypeGek, nil
	case "руквкр":
		return RecordTypeThesisSupervision, nil
	case "рец":
		return RecordTypeReviewing, nil
	case "учпр":
		return RecordTypeStudyPractice, nil
	case "прпр":
		return RecordTypeWorkPractice, nil
	case "предпр":
		return RecordTypePreDiplomaPractice, nil
	default:
		return 0, ErrUnexpectedRecordType
	}
}

func (r RecordType) String() (string, error) {
	switch r {
	case RecordTypeLecture:
		return "лек", nil
	case RecordTypePractice:
		return "пр", nil
	case RecordTypeLab:
		return "лаб", nil
	case RecordTypeKsr:
		return "кср", nil
	case RecordTypeCourseProject:
		return "кп", nil
	case RecordTypeCourseWork:
		return "кр", nil
	case RecordTypeControlWork:
		return "контрраб", nil
	case RecordTypePass:
		return "зач", nil
	case RecordTypeGradedPass:
		return "дифзач", nil
	case RecordTypeExam:
		return "экз", nil
	case RecordTypeStateExam:
		return "госэкз", nil
	case RecordTypeGek:
		return "гэк", nil
	case RecordTypeThesisSupervision:
		return "руквкр", nil
	case RecordTypeReviewing:
		return "рец", nil
	case RecordTypeStudyPractice:
		return "учпр", nil
	case RecordTypeWorkPractice:
		return "прпр", nil
	case RecordTypePreDiplomaPractice:
		return "предпр", nil
	default:
		return "", ErrUnexpectedRecordType
	}
}

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

func NewGroupsSubgroups(groupID uuid.UUID, subgroup SubgroupType) GroupsSubgroups {
	return GroupsSubgroups{subgroupId: groupID, subgroupSequenceNumber: subgroup}
}

func (g GroupsSubgroups) GroupID() uuid.UUID {
	return g.subgroupId
}

func (g GroupsSubgroups) Subgroup() SubgroupType {
	return g.subgroupSequenceNumber
}

// Term семестр. Строковое представление совпадает с названием листа в xlsx
type Term int

const (
	TermFirst Term = iota + 1
	TermSecond
)

func NewTermFromString(s string) (Term, error) {
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)
	switch s {
	case "осенний семестр":
		return TermFirst, nil
	case "весенний семестр":
		return TermSecond, nil
	default:
		return 0, ErrUnexpectedTerm
	}
}

func (t Term) String() string {
	switch t {
	case TermFirst:
		return "Осенний семестр"
	case TermSecond:
		return "Весенний семестр"
	default:
		return fmt.Sprintf("Term(%d)", int(t))
	}
}

// TeacherHours тройка: айди преподавателя - количество часов - является ли доп нагрузкой
type TeacherHours struct {
	teacherId uuid.UUID
	hours     int
	isExtra   bool
}

// Record базовая модель дисциплины
type Record struct {
	id             uuid.UUID
	importFileID   uuid.UUID
	groups         []GroupsSubgroups
	studentsAmount int
	disciplineId   uuid.UUID
	term           Term
	recordType     RecordType
	goalHours      int
	teachersHours  []TeacherHours
	isError        bool
	deletedAt      *time.Time
}

func (r *Record) StudentsAmount() int {
	return r.studentsAmount
}

func (r *Record) SetStudentsAmount(studentsAmount int) {
	r.studentsAmount = studentsAmount
}

func NewRecord(id uuid.UUID, importFileID uuid.UUID, groups []GroupsSubgroups, studentsAmount int, disciplineId uuid.UUID, term Term, recordType RecordType, goalHours int, teachersHours []TeacherHours, isError bool, deletedAt *time.Time) *Record {
	return &Record{id: id, importFileID: importFileID, groups: groups, studentsAmount: studentsAmount, disciplineId: disciplineId, term: term, recordType: recordType, goalHours: goalHours, teachersHours: teachersHours, isError: isError, deletedAt: deletedAt}
}

func (r *Record) SetId(id uuid.UUID) {
	r.id = id
}

func (r *Record) SetDisciplineId(disciplineId uuid.UUID) {
	r.disciplineId = disciplineId
}

func (r *Record) SetGroups(groups []GroupsSubgroups) {
	r.groups = groups
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
	newDeletedAt    *time.Time
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

func (r RecordAction) NewDeletedAt() *time.Time {
	return r.newDeletedAt
}

func NewRecordAction(id uuid.UUID, actorName string, revertsActionID *uuid.UUID, recordId uuid.UUID, teacherId uuid.UUID, newHours *int, newIsExtra bool, newDeletedAt *time.Time) *RecordAction {
	return &RecordAction{id: id, actorName: actorName, revertsActionID: revertsActionID, recordId: recordId, teacherId: teacherId, newHours: newHours, newIsExtra: newIsExtra, newDeletedAt: newDeletedAt}
}
