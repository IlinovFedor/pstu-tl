package models

import "uuid"

type FileReport struct {
	id                  uuid.UUID
	name                string
	isDeleted           bool
	associatedTeacherId uuid.UUID
	data                []byte
}

func (f *FileReport) SetIsDeleted(isDeleted bool) {
	f.isDeleted = isDeleted
}

func (f FileReport) Id() uuid.UUID {
	return f.id
}

func (f FileReport) Name() string {
	return f.name
}

func (f FileReport) IsDeleted() bool {
	return f.isDeleted
}

func (f FileReport) AssociatedTeacherId() uuid.UUID {
	return f.associatedTeacherId
}

func (f FileReport) Data() []byte {
	return f.data
}

func NewFileReport(id uuid.UUID, name string, isDeleted bool, associatedTeacherId uuid.UUID, data []byte) *FileReport {
	return &FileReport{id: id, name: name, isDeleted: isDeleted, associatedTeacherId: associatedTeacherId, data: data}
}

type FileReportAction struct {
	id              uuid.UUID
	actorName       string
	revertsActionID *uuid.UUID
	newIsDeleted    bool
}

func (i FileReportAction) Id() uuid.UUID {
	return i.id
}

func (i FileReportAction) ActorName() string {
	return i.actorName
}

func (i FileReportAction) RevertsActionID() *uuid.UUID {
	return i.revertsActionID
}

func (i FileReportAction) NewIsDeleted() bool {
	return i.newIsDeleted
}

func NewFileReportAction(id uuid.UUID, actorName string, revertsActionID *uuid.UUID, newIsDeleted bool) *FileReportAction {
	return &FileReportAction{id: id, actorName: actorName, revertsActionID: revertsActionID, newIsDeleted: newIsDeleted}
}
