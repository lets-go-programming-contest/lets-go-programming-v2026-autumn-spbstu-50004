package main

import "fmt"

func updateRange(left, right int, operation string, value int) (int, int) {
	switch operation {
	case "<=":
		if value < right {
			right = value
		}
	case ">=":
		if value > left {
			left = value
		}
	case "<":
		if value-1 < right {
			right = value - 1
		}
	case ">":
		if value+1 > left {
			left = value + 1
		}
	}

	return left, right
}

func main() {
	var departmentsCount int
	if _, err := fmt.Scan(&departmentsCount); err != nil {
		return
	}

	for range departmentsCount {
		var employeesCount int
		if _, err := fmt.Scan(&employeesCount); err != nil {
			return
		}

		left, right := 15, 30

		for range employeesCount {
			var operation string

			var value int

			if _, err := fmt.Scan(&operation, &value); err != nil {
				return
			}

			left, right = updateRange(left, right, operation, value)

			if left <= right {
				fmt.Println(left)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
