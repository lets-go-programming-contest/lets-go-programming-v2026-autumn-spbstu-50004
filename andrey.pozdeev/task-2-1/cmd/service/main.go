package main

import "fmt"

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

func main() {
	var totalDepts int
	if _, err := fmt.Scan(&totalDepts); err != nil {
		fmt.Println("Invalid departments count:", err)

		return
	}

	for range totalDepts {
		var totalStaff int
		if _, err := fmt.Scan(&totalStaff); err != nil {
			fmt.Println("Invalid employees count:", err)

			return
		}

		ctrl := newClimateControl()

		for range totalStaff {
			var (
				condOperator string
				wantedTemp   int
			)

			if _, err := fmt.Scan(&condOperator, &wantedTemp); err != nil {
				fmt.Println("Invalid temperature constraint:", err)

				return
			}

			fmt.Println(ctrl.adjust(condOperator, wantedTemp))
		}
	}
}
