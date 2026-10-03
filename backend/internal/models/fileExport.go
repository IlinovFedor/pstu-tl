package models

import "uuid"

type FileExport struct {
	id        uuid.UUID
	name      string
	isDeleted bool
	data      []byte
}

func NewFileExport(id uuid.UUID, name string, isDeleted bool, data []byte) *FileExport {
	return &FileExport{id: id, name: name, isDeleted: isDeleted, data: data}
}

func (e *FileExport) SetIsDeleted(isDeleted bool) {
	e.isDeleted = isDeleted
}

func (e FileExport) Id() uuid.UUID {
	return e.id
}

func (e FileExport) Name() string {
	return e.name
}

func (e FileExport) IsDeleted() bool {
	return e.isDeleted
}

func (e FileExport) Data() []byte {
	return e.data
}

type FileExportAction struct {
	id              uuid.UUID
	actorName       string
	revertsActionID *uuid.UUID
	newIsDeleted    bool
}

func (f FileExportAction) Id() uuid.UUID {
	return f.id
}

func (f FileExportAction) ActorName() string {
	return f.actorName
}

func (f FileExportAction) RevertsActionID() *uuid.UUID {
	return f.revertsActionID
}

func (f FileExportAction) NewIsDeleted() bool {
	return f.newIsDeleted
}

func NewFileExportAction(id uuid.UUID, actorName string, revertsActionID *uuid.UUID, newIsDeleted bool) *FileExportAction {
	return &FileExportAction{id: id, actorName: actorName, revertsActionID: revertsActionID, newIsDeleted: newIsDeleted}
}
