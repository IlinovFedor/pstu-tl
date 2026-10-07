package models

import (
	"time"
	"uuid"
)

type FileExport struct {
	id        uuid.UUID
	yearName  int
	deletedAt *time.Time
}

func (f *FileExport) SetDeletedAt(deletedAt *time.Time) {
	f.deletedAt = deletedAt
}

func (f FileExport) Id() uuid.UUID {
	return f.id
}

func (f FileExport) YearName() int {
	return f.yearName
}

func (f FileExport) DeletedAt() *time.Time {
	return f.deletedAt
}

func NewFileExport(id uuid.UUID, yearName int, deletedAt *time.Time) *FileExport {
	return &FileExport{id: id, yearName: yearName, deletedAt: deletedAt}
}

type FileExportAction struct {
	id              uuid.UUID
	actorName       string
	revertsActionID *uuid.UUID
	fileID          *uuid.UUID
	newDeletedAt    *time.Time
}

func (f FileExportAction) NewDeletedAt() *time.Time {
	return f.newDeletedAt
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

func NewFileExportAction(id uuid.UUID, actorName string, revertsActionID *uuid.UUID, fileID *uuid.UUID, newDeletedAt *time.Time) *FileExportAction {
	return &FileExportAction{id: id, actorName: actorName, revertsActionID: revertsActionID, fileID: fileID, newDeletedAt: newDeletedAt}
}
