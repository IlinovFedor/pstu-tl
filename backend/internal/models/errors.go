package models

import (
	"errors"
	"fmt"
	"strings"
)

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
	ErrUnknownGroupType   = errors.New("cannot parse group type")

	// Ошибки парсинга входных XLSX

	ErrUnparsableGroupCell           = errors.New("cannot parse group cell")
	ErrUnparsableStudentsAmount      = errors.New("cannot parse record students amount")
	ErrUnparsableYearOfEnrollment    = errors.New("cannot parse year of enrollment")
	ErrUnparsableGroupSequenceNumber = errors.New("cannot parse group sequence number")
	ErrUnparsableSubgroupNumber      = errors.New("cannot parse record subgroup number")
	ErrUnknownRecordType             = errors.New("unknown record type")
	ErrUnparsableGoalHours           = errors.New("cannot parse record hours")
	ErrEmptyDisciplineCol            = errors.New("empty discipline")
)

type CellError struct {
	err   error
	term  Term
	cell  string
	value string
}

func NewCellError(err error, term Term, cell string, value string) CellError {
	return CellError{err: err, term: term, cell: cell, value: value}
}

func (c CellError) Error() string {
	return fmt.Sprintf("term: %s cell: %s value: %s: %v", c.term, c.cell, c.value, c.err)
}

// Unwrap нужен, чтобы errors.Is(err, ErrUnparsableGoalHours) и т.п. работали.
func (c CellError) Unwrap() error { return c.err }

func (c CellError) Value() string { return c.value }
func (c CellError) Cell() string  { return c.cell }
func (c CellError) Term() Term    { return c.term }
func (c CellError) Err() error    { return c.err }

type ParseError struct {
	cells []CellError
}

func NewParseError(cells []CellError) *ParseError {
	return &ParseError{cells: cells}
}

func (p *ParseError) AddErr(cellError CellError) {
	p.cells = append(p.cells, cellError)
}

func (p *ParseError) Cells() []CellError {
	return p.cells
}

func (p *ParseError) SetCells(cells []CellError) {
	p.cells = cells
}

func (p *ParseError) HasErrors() bool {
	return p != nil && len(p.cells) > 0
}

func (p *ParseError) Error() string {
	msgs := make([]string, len(p.cells))
	for i, c := range p.cells {
		msgs[i] = c.Error()
	}
	return fmt.Sprintf("%d errors:\n%s", len(p.cells), strings.Join(msgs, "\n"))
}

func (p *ParseError) Unwrap() []error {
	errs := make([]error, len(p.cells))
	for i, c := range p.cells {
		errs[i] = c
	}
	return errs
}
