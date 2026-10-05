package models

import (
	"uuid"
)

type Year struct {
	name               int
	wageNormalHours    int
	extHourlyWageTier1 int
	extHourlyWageTier2 int
	extHourlyWageTier3 int
	wageTolerance      int
	isDeleted          bool
}

func NewYear(name int, wageNormalHours int, extHourlyWageTier1 int, extHourlyWageTier2 int, extHourlyWageTier3 int, wageTolerance int, isDeleted bool) *Year {
	return &Year{name: name, wageNormalHours: wageNormalHours, extHourlyWageTier1: extHourlyWageTier1, extHourlyWageTier2: extHourlyWageTier2, extHourlyWageTier3: extHourlyWageTier3, wageTolerance: wageTolerance, isDeleted: isDeleted}
}

func (y *Year) IsDeleted() bool {
	return y.isDeleted
}

func (y *Year) SetIsDeleted(isDeleted bool) {
	y.isDeleted = isDeleted
}

func (y *Year) Name() int {
	return y.name
}

func (y *Year) SetName(name int) {
	y.name = name
}

func (y *Year) WageNormalHours() int {
	return y.wageNormalHours
}

func (y *Year) SetWageNormalHours(wageNormalHours int) {
	y.wageNormalHours = wageNormalHours
}

func (y *Year) ExtHourlyWageTier1() int {
	return y.extHourlyWageTier1
}

func (y *Year) SetExtHourlyWageTier1(extHourlyWageTier1 int) {
	y.extHourlyWageTier1 = extHourlyWageTier1
}

func (y *Year) ExtHourlyWageTier2() int {
	return y.extHourlyWageTier2
}

func (y *Year) SetExtHourlyWageTier2(extHourlyWageTier2 int) {
	y.extHourlyWageTier2 = extHourlyWageTier2
}

func (y *Year) ExtHourlyWageTier3() int {
	return y.extHourlyWageTier3
}

func (y *Year) SetExtHourlyWageTier3(extHourlyWageTier3 int) {
	y.extHourlyWageTier3 = extHourlyWageTier3
}

func (y *Year) WageTolerance() int {
	return y.wageTolerance
}

func (y *Year) SetWageTolerance(wageTolerance int) {
	y.wageTolerance = wageTolerance
}

// YearAction модель изменения года. Получается только из репозитория, append-only
type YearAction struct {
	id                    uuid.UUID
	actorName             string
	revertsActionID       *uuid.UUID
	name                  int
	newWageNormalHours    int
	newExtHourlyWageTier1 int
	newExtHourlyWageTier2 int
	newExtHourlyWageTier3 int
	newWageTolerance      int
	newIsDeleted          bool
}

func NewYearAction(id uuid.UUID, actorName string, revertsActionID *uuid.UUID, newWageNormalHours int, newExtHourlyWageTier1 int, newExtHourlyWageTier2 int, newExtHourlyWageTier3 int, newWageTolerance int, newIsDeleted bool) *YearAction {
	return &YearAction{id: id, actorName: actorName, revertsActionID: revertsActionID, newWageNormalHours: newWageNormalHours, newExtHourlyWageTier1: newExtHourlyWageTier1, newExtHourlyWageTier2: newExtHourlyWageTier2, newExtHourlyWageTier3: newExtHourlyWageTier3, newWageTolerance: newWageTolerance, newIsDeleted: newIsDeleted}
}

func (y YearAction) NewIsDeleted() bool {
	return y.newIsDeleted
}

func (y YearAction) NewWageTolerance() int {
	return y.newWageTolerance
}

func (y YearAction) NewExtHourlyWageTier3() int {
	return y.newExtHourlyWageTier3
}

func (y YearAction) NewExtHourlyWageTier2() int {
	return y.newExtHourlyWageTier2
}

func (y YearAction) NewExtHourlyWageTier1() int {
	return y.newExtHourlyWageTier1
}

func (y YearAction) NewWageNormalHours() int {
	return y.newWageNormalHours
}

func (y YearAction) Name() int {
	return y.name
}

func (y YearAction) RevertsActionID() *uuid.UUID {
	return y.revertsActionID
}

func (y YearAction) ActorName() string {
	return y.actorName
}

func (y YearAction) Id() uuid.UUID {
	return y.id
}
