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

			switch operation {
			case ">=":
				if targetTemp > currentMinTemp {
					currentMinTemp = targetTemp
				}
			case "<=":
				if targetTemp < currentMaxTemp {
					currentMaxTemp = targetTemp
				}
			}

			if currentMinTemp <= currentMaxTemp {
				fmt.Println(currentMinTemp)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
