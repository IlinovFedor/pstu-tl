package models

import (
	"time"
	"uuid"
)

type FileImport struct {
	id        uuid.UUID
	yearName  int
	name      string
	hash      string
	deletedAt *time.Time
}

func (f *FileImport) SetDeletedAt(deletedAt *time.Time) {
	f.deletedAt = deletedAt
}

func (f FileImport) Id() uuid.UUID {
	return f.id
}

func (f FileImport) YearName() int {
	return f.yearName
}

func (f FileImport) Name() string {
	return f.name
}

func (f FileImport) Hash() string {
	return f.hash
}

func (f FileImport) DeletedAt() *time.Time {
	return f.deletedAt
}

func NewFileImport(id uuid.UUID, yearName int, name string, hash string) *FileImport {
	return &FileImport{id: id, yearName: yearName, name: name, hash: hash}
}

type FileImportAction struct {
	id              uuid.UUID
	actorName       string
	revertsActionID *uuid.UUID
	fileID          uuid.UUID
	newDeletedAt    *time.Time
}

func (f FileImportAction) Id() uuid.UUID {
	return f.id
}

func (f FileImportAction) ActorName() string {
	return f.actorName
}

func (f FileImportAction) RevertsActionID() *uuid.UUID {
	return f.revertsActionID
}

func (f FileImportAction) FileImportID() uuid.UUID {
	return f.fileID
}

func (f FileImportAction) NewDeletedAt() *time.Time {
	return f.newDeletedAt
}

func NewFileImportAction(id uuid.UUID, actorName string, revertsActionID *uuid.UUID, fileID uuid.UUID, newDeletedAt *time.Time) *FileImportAction {
	return &FileImportAction{id: id, actorName: actorName, revertsActionID: revertsActionID, fileID: fileID, newDeletedAt: newDeletedAt}
}
