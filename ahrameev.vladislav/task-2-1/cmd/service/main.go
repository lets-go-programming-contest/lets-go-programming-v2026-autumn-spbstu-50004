package main

import (
	"fmt"
	"os"
)

const (
	initialLeft  = 15
	initialRight = 30
)

func updateRange(left, right int, operation string, val int) (int, int) {
	switch operation {
	case "<=":
		if val < right {
			right = val
		}
	case ">=":
		if val > left {
			left = val
		}
	case "<":
		if val-1 < right {
			right = val - 1
		}
	case ">":
		if val+1 > left {
			left = val + 1
		}
	}

	return left, right
}

func main() {
	var departmentsCount int
	if _, err := fmt.Scan(&departmentsCount); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading departments count: %v\n", err)

		return
	}

	for range departmentsCount {
		var employeesCount int
		if _, err := fmt.Scan(&employeesCount); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading employees count: %v\n", err)

			return
		}

		left, right := initialLeft, initialRight

		for range employeesCount {
			var operation string

			var val int
			if _, err := fmt.Scan(&operation, &val); err != nil {
				fmt.Fprintf(os.Stderr, "Error reading operation: %v\n", err)

				return
			}

			left, right = updateRange(left, right, operation, val)

			if left <= right {
				fmt.Println(left)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
