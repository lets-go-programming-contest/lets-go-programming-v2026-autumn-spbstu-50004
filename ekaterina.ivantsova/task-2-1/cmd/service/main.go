package main

import "fmt"

const (
	DefaultMinTemp = 15
	DefaultMaxTemp = 30
)

func updateTemperatureBounds(operation string, targetTemp int, currentMin, currentMax int) (int, int) {
	switch operation {
	case ">=":
		if targetTemp > currentMin {
			currentMin = targetTemp
		}
	case "<=":
		if targetTemp < currentMax {
			currentMax = targetTemp
		}
	}
	return currentMin, currentMax
}

func processDepartment() {
	var employeesCount int
	if _, err := fmt.Scan(&employeesCount); err != nil {
		return
	}

	var (
		currentMinTemp = DefaultMinTemp
		currentMaxTemp = DefaultMaxTemp
	)

	for range employeesCount {
		var (
			operation  string
			targetTemp int
		)
		if _, err := fmt.Scan(&operation, &targetTemp); err != nil {
			break
		}

		currentMinTemp, currentMaxTemp = updateTemperatureBounds(operation, targetTemp, currentMinTemp, currentMaxTemp)

		if currentMinTemp <= currentMaxTemp {
			fmt.Println(currentMinTemp)
		} else {
			fmt.Println(-1)
		}
	}
}

func main() {
	var departmentsCount int
	if _, err := fmt.Scan(&departmentsCount); err != nil {
		return
	}

	for range departmentsCount {
		processDepartment()
	}
}
