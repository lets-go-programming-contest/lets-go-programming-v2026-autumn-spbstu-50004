package main

import "fmt"

const (
	minTemp = 15
	maxTemp = 30
)

type Conditioner struct {
	lo int
	hi int
}

func NewConditioner() *Conditioner {
	return &Conditioner{lo: minTemp, hi: maxTemp}
}

func (c *Conditioner) Apply(operation string, value int) {
	switch operation {
	case ">=":
		if value > c.lo {
			c.lo = value
		}
	case "<=":
		if value < c.hi {
			c.hi = value
		}
	}
}

func (c *Conditioner) Optimal() int {
	if c.lo <= c.hi {
		return c.lo
	}

	return -1
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

		conditioner := NewConditioner()

		for range employeesCount {
			var (
				operation string
				value     int
			)

			if _, err := fmt.Scan(&operation, &value); err != nil {
				return
			}

			conditioner.Apply(operation, value)

			fmt.Println(conditioner.Optimal())
		}
	}
}
