package main

import "fmt"

const (
	DefaultMinTemp = 15
	DefaultMaxTemp = 30
)

func main() {
	var departmentsCount int
	if _, err := fmt.Scan(&departmentsCount); err != nil {
		return
	}

	for range departmentsCount {
		var employeesCount int
		if _, err := fmt.Scan(&employeesCount); err != nil {
			break
		}

		for range employeesCount {
			var (
				operation  string
				targetTemp int
			)
			if _, err := fmt.Scan(&operation, &targetTemp); err != nil {
				break
			}
		}
	}
}
