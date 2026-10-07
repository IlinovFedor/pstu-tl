package models

import (
	"time"
	"uuid"
)

type FileReport struct {
	id        uuid.UUID
	deletedAt *time.Time
	teacherID uuid.UUID
}

func (f *FileReport) SetDeletedAt(deletedAt *time.Time) {
	f.deletedAt = deletedAt
}

func (f FileReport) TeacherID() uuid.UUID {
	return f.teacherID
}

func (f FileReport) DeletedAt() *time.Time {
	return f.deletedAt
}

func (f FileReport) Id() uuid.UUID {
	return f.id
}

func NewFileReport(id uuid.UUID, deletedAt *time.Time, teacherID uuid.UUID) *FileReport {
	return &FileReport{id: id, deletedAt: deletedAt, teacherID: teacherID}
}

type FileReportAction struct {
	id              uuid.UUID
	actorName       string
	revertsActionID *uuid.UUID
	fileID          uuid.UUID
	newDeletedAt    *time.Time
}

func (f FileReportAction) NewDeletedAt() *time.Time {
	return f.newDeletedAt
}

func (f FileReportAction) FileID() uuid.UUID {
	return f.fileID
}

func (f FileReportAction) RevertsActionID() *uuid.UUID {
	return f.revertsActionID
}

func (f FileReportAction) ActorName() string {
	return f.actorName
}

func (f FileReportAction) Id() uuid.UUID {
	return f.id
}

func NewFileReportAction(id uuid.UUID, actorName string, revertsActionID *uuid.UUID, fileID uuid.UUID, newDeletedAt *time.Time) *FileReportAction {
	return &FileReportAction{id: id, actorName: actorName, revertsActionID: revertsActionID, fileID: fileID, newDeletedAt: newDeletedAt}
}
