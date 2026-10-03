package models

import "uuid"

type FileImport struct {
	hash      string
	name      string
	isDeleted bool
	data      []byte
}

func NewFileImport(hash string, name string, isDeleted bool, data []byte) *FileImport {
	return &FileImport{hash: hash, name: name, isDeleted: isDeleted, data: data}
}

func (i *FileImport) Data() []byte {
	return i.data
}

func (i *FileImport) SetIsDeleted(isDeleted bool) {
	i.isDeleted = isDeleted
}

func (i FileImport) IsDeleted() bool {
	return i.isDeleted
}

func (i FileImport) Name() string {
	return i.name
}

func (i FileImport) Hash() string {
	return i.hash
}

type FileImportAction struct {
	id              uuid.UUID
	actorName       string
	revertsActionID *uuid.UUID
	newIsDeleted    bool
}

func (i FileImportAction) Id() uuid.UUID {
	return i.id
}

func (i FileImportAction) ActorName() string {
	return i.actorName
}

func (i FileImportAction) RevertsActionID() *uuid.UUID {
	return i.revertsActionID
}

func (i FileImportAction) NewIsDeleted() bool {
	return i.newIsDeleted
}

func NewFileImportAction(id uuid.UUID, actorName string, revertsActionID *uuid.UUID, newIsDeleted bool) *FileImportAction {
	return &FileImportAction{id: id, actorName: actorName, revertsActionID: revertsActionID, newIsDeleted: newIsDeleted}
}
