package models

import (
	"uuid"
)

type FileReport struct {
	id                  uuid.UUID
	isDeleted           bool
	associatedTeacherId uuid.UUID
}

func (f *FileReport) SetIsDeleted(isDeleted bool) {
	f.isDeleted = isDeleted
}

func (f FileReport) AssociatedTeacherId() uuid.UUID {
	return f.associatedTeacherId
}

func (f FileReport) IsDeleted() bool {
	return f.isDeleted
}

func (f FileReport) Id() uuid.UUID {
	return f.id
}

func NewFileReport(id uuid.UUID, isDeleted bool, associatedTeacherId uuid.UUID) *FileReport {
	return &FileReport{id: id, isDeleted: isDeleted, associatedTeacherId: associatedTeacherId}
}

type FileReportAction struct {
	id              uuid.UUID
	actorName       string
	revertsActionID *uuid.UUID
	fileID          uuid.UUID
	newIsDeleted    bool
}

func (f FileReportAction) NewIsDeleted() bool {
	return f.newIsDeleted
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

func NewFileReportAction(id uuid.UUID, actorName string, revertsActionID *uuid.UUID, fileID uuid.UUID, newIsDeleted bool) *FileReportAction {
	return &FileReportAction{id: id, actorName: actorName, revertsActionID: revertsActionID, fileID: fileID, newIsDeleted: newIsDeleted}
}
