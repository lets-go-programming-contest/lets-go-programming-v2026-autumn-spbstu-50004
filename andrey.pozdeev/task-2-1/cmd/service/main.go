package main

const (
	minLimitTemp = 15
	maxLimitTemp = 30
	opGreaterEq  = ">="
	opLessEq     = "<="
	invalidTemp  = -1
)

type climateControl struct {
	minTemp int
	maxTemp int
}

func newClimateControl() climateControl {
	return climateControl{
		minTemp: minLimitTemp,
		maxTemp: maxLimitTemp,
	}
}

func (cc *climateControl) adjust(operator string, target int) int {
	switch operator {
	case opGreaterEq:
		if target > cc.minTemp {
			cc.minTemp = target
		}
	case opLessEq:
		if target < cc.maxTemp {
			cc.maxTemp = target
		}
	}

	if cc.minTemp > cc.maxTemp {
		return invalidTemp
	}

	return cc.minTemp
}

func main() {}
