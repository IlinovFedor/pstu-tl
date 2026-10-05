package models

import "errors"

var (
	ErrBadRequest         = errors.New("bad request")
	ErrUnauthorized       = errors.New("unauthorized")
	ErrForbidden          = errors.New("forbidden")
	ErrUnknownYear        = errors.New("unknown year")
	ErrUnknownTeacher     = errors.New("unknown teacher")
	ErrUnknownGroup       = errors.New("unknown group")
	ErrUnknownDiscipline  = errors.New("unknown discipline")
	ErrUnknownUser        = errors.New("unknown user")
	ErrUnknownFile        = errors.New("unknown file")
	ErrUnknownAction      = errors.New("unknown action")
	ErrCannotRevertAction = errors.New("cannot revert action")
	ErrDuplicateFileName  = errors.New("duplicate file name")
	ErrDuplicateFileData  = errors.New("duplicate file data")
	ErrInvalidFileData    = errors.New("invalid file data")
	ErrUnknownRecordType  = errors.New("unknown record type")
	ErrUnparsableGroup    = errors.New("cannot parse group name")
	ErrUnparsableHours    = errors.New("cannot parse record hours")
)
