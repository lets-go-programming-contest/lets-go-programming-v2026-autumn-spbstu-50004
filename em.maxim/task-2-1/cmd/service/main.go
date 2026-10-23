package main

import (
	"fmt"
)

type Department struct {
	lowerBound int
	upperBound int
}

func CreateDepartment() Department {
	return Department{
		lowerBound: 15,
		upperBound: 30,
	}
}

func (d *Department) ApplyConstraint(sign string, val int) int {
	switch sign {
	case ">=":
		if val > d.lowerBound {
			d.lowerBound = val
		}
	case "<=":
		if val < d.upperBound {
			d.upperBound = val
		}
	}

	if d.lowerBound > d.upperBound {
		return -1
	}

	return d.lowerBound
}

func main() {
	var numDepts int
	if _, err := fmt.Scan(&numDepts); err != nil {
		fmt.Println("Invalid departments count")
		return
	}

	for i := 0; i < numDepts; i++ {
		var numStaff int
		if _, err := fmt.Scan(&numStaff); err != nil {
			fmt.Println("Invalid employees count")
			return
		}

		dept := CreateDepartment()

		for j := 0; j < numStaff; j++ {
			var op string
			var temp int

			if _, err := fmt.Scan(&op, &temp); err != nil {
				fmt.Println("Invalid temperature constraint")
				return
			}

			fmt.Println(dept.ApplyConstraint(op, temp))
		}
	}
}
