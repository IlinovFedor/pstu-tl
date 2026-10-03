package models

type ModelError int

const (
	ModelErrorBadRequest = iota
	ModelErrorUnauthorized
	ModelErrorForbidden
	ModelErrorUnknownYear
	ModelErrorUnknownTeacher
	ModelErrorUnknownGroup
	ModelErrorUnknownDiscipline
	ModelErrorUnknownUser
	ModelErrorUnknownFile
	ModelErrorUnknownAction
	ModelErrorCannotRevertAction
	ModelErrorFileAlreadyAdded
)

func (e ModelError) Code() string {
	switch e {
	case ModelErrorBadRequest:
		return "errorCodeBadRequest"
	case ModelErrorUnauthorized:
		return "errorCodeUnauthorized"
	case ModelErrorForbidden:
		return "errorCodeForbidden"
	case ModelErrorUnknownYear:
		return "errorCodeUnknownYear"
	case ModelErrorUnknownTeacher:
		return "errorCodeUnknownTeacher"
	case ModelErrorUnknownGroup:
		return "errorCodeUnknownGroup"
	case ModelErrorUnknownDiscipline:
		return "errorCodeUnknownDiscipline"
	case ModelErrorUnknownUser:
		return "errorCodeUnknownUser"
	case ModelErrorUnknownFile:
		return "errorCodeUnknownFile"
	case ModelErrorUnknownAction:
		return "errorCodeUnknownAction"
	case ModelErrorCannotRevertAction:
		return "errorCodeCannotRevertAction"
	case ModelErrorFileAlreadyAdded:
		return "errorCodeFileAlreadyAdded"
	default:
		return "unknown"
	}
}
