package main

import "fmt"

const (
	defaultMinTemp = 15
	defaultMaxTemp = 30
)

type temperatureBounds struct {
	min int
	max int
}

func (t *temperatureBounds) update(operation string, targetTemp int) {
	switch operation {
	case ">=":
		if targetTemp > t.min {
			t.min = targetTemp
		}
	case "<=":
		if targetTemp < t.max {
			t.max = targetTemp
		}
	}
}

func processDepartment() {
	var employeesCount int
	if _, err := fmt.Scan(&employeesCount); err != nil {
		fmt.Println("Invalid employees count")
		return
	}

	bounds := temperatureBounds{
		min: defaultMinTemp,
		max: defaultMaxTemp,
	}

	for range employeesCount {
		var (
			operation  string
			targetTemp int
		)
		if _, err := fmt.Scan(&operation, &targetTemp); err != nil {
			fmt.Println("Invalid employee data")
			break
		}

		bounds.update(operation, targetTemp)

		if bounds.min <= bounds.max {
			fmt.Println(bounds.min)
		} else {
			fmt.Println(-1)
		}
	}
}

func main() {
	var departmentsCount int
	if _, err := fmt.Scan(&departmentsCount); err != nil {
		fmt.Println("Invalid departments count")
		return
	}

	for range departmentsCount {
		processDepartment()
	}
}
