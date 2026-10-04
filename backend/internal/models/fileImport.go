package models

import (
	"uuid"
)

type FileImport struct {
	id        uuid.UUID
	yearName  int
	name      string
	hash      string
	isDeleted bool
}

func (f *FileImport) SetIsDeleted(isDeleted bool) {
	f.isDeleted = isDeleted
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

func (f FileImport) IsDeleted() bool {
	return f.isDeleted
}

func NewFileImport(yearName int, name string, hash string) *FileImport {
	return &FileImport{id: uuid.NewV7(), yearName: yearName, name: name, hash: hash}
}

type FileImportAction struct {
	id              uuid.UUID
	actorName       string
	revertsActionID *uuid.UUID
	fileID          uuid.UUID
	newIsDeleted    bool
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

func (f FileImportAction) NewIsDeleted() bool {
	return f.newIsDeleted
}

func NewFileImportAction(id uuid.UUID, actorName string, revertsActionID *uuid.UUID, fileID uuid.UUID, newIsDeleted bool) *FileImportAction {
	return &FileImportAction{id: id, actorName: actorName, revertsActionID: revertsActionID, fileID: fileID, newIsDeleted: newIsDeleted}
}
