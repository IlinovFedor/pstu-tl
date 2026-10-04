package models

import (
	"uuid"
)

type FileExport struct {
	id        uuid.UUID
	yearName  int
	isDeleted bool
}

func (f *FileExport) SetIsDeleted(isDeleted bool) {
	f.isDeleted = isDeleted
}

func (f FileExport) Id() uuid.UUID {
	return f.id
}

func (f FileExport) YearName() int {
	return f.yearName
}

func (f FileExport) IsDeleted() bool {
	return f.isDeleted
}

func NewFileExport(id uuid.UUID, yearName int, isDeleted bool) *FileExport {
	return &FileExport{id: id, yearName: yearName, isDeleted: isDeleted}
}

type FileExportAction struct {
	id              uuid.UUID
	actorName       string
	revertsActionID *uuid.UUID
	fileID          *uuid.UUID
	newIsDeleted    bool
}

func (f FileExportAction) NewIsDeleted() bool {
	return f.newIsDeleted
}

func (f FileExportAction) FileExportID() *uuid.UUID {
	return f.fileID
}

func (f FileExportAction) RevertsActionID() *uuid.UUID {
	return f.revertsActionID
}

func (f FileExportAction) ActorName() string {
	return f.actorName
}

func (f FileExportAction) Id() uuid.UUID {
	return f.id
}

func NewFileExportAction(id uuid.UUID, actorName string, revertsActionID *uuid.UUID, fileID *uuid.UUID, newIsDeleted bool) *FileExportAction {
	return &FileExportAction{id: id, actorName: actorName, revertsActionID: revertsActionID, fileID: fileID, newIsDeleted: newIsDeleted}
}
