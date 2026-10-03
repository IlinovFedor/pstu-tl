package models

type Calculation struct {
	fullTimeHoursTerm1     int
	fullTimeHoursTerm2     int
	postgraduateHoursTerm1 int
	postgraduateHoursTerm2 int
	calculatedWage         int
	partTimeHoursTerm1     int
	partTimeHoursTerm2     int
	pastTimeConstYearly    int
	wageDiffHours          int
}

func NewCalculation(fullTimeHoursTerm1 int, fullTimeHoursTerm2 int, postgraduateHoursTerm1 int, postgraduateHoursTerm2 int, calculatedWage int, partTimeHoursTerm1 int, partTimeHoursTerm2 int, pastTimeConstYearly int, wageDiffHours int) *Calculation {
	return &Calculation{fullTimeHoursTerm1: fullTimeHoursTerm1, fullTimeHoursTerm2: fullTimeHoursTerm2, postgraduateHoursTerm1: postgraduateHoursTerm1, postgraduateHoursTerm2: postgraduateHoursTerm2, calculatedWage: calculatedWage, partTimeHoursTerm1: partTimeHoursTerm1, partTimeHoursTerm2: partTimeHoursTerm2, pastTimeConstYearly: pastTimeConstYearly, wageDiffHours: wageDiffHours}
}

func (c Calculation) FullTimeHoursTerm1() int {
	return c.fullTimeHoursTerm1
}

func (c Calculation) FullTimeHoursTerm2() int {
	return c.fullTimeHoursTerm2
}

func (c Calculation) PostgraduateHoursTerm1() int {
	return c.postgraduateHoursTerm1
}

func (c Calculation) PostgraduateHoursTerm2() int {
	return c.postgraduateHoursTerm2
}

func (c Calculation) CalculatedWage() int {
	return c.calculatedWage
}

func (c Calculation) PartTimeHoursTerm1() int {
	return c.partTimeHoursTerm1
}

func (c Calculation) PartTimeHoursTerm2() int {
	return c.partTimeHoursTerm2
}

func (c Calculation) PastTimeConstYearly() int {
	return c.pastTimeConstYearly
}

func (c Calculation) WageDiffHours() int {
	return c.wageDiffHours
}
