package main

import (
	"errors"
	"fmt"
)

const (
	MaxTemp = 30
	MinTemp = 15
	lte     = "<="
	gte     = ">="
	NoTemp  = -1
)

type conditioner struct {
	low  int
	high int
}

func newConditioner() *conditioner {
	return &conditioner{high: MaxTemp, low: MinTemp}
}

func (c *conditioner) regulate(operator string, temp int) (int, error) {
	switch operator {
	case lte:
		if temp < c.high {
			c.high = temp
		}
	case gte:
		if temp > c.low {
			c.low = temp
		}
	default:
		return NoTemp, errors.New("invalid operator")
	}

	if c.low > c.high {
		return NoTemp, errors.New("no valid temperature")
	}

	return c.low, nil
}

func main() {
	var departamentsCount int

	_, err := fmt.Scan(&departamentsCount)
	if err != nil {
		fmt.Println(errors.New("invalid department number"))
		return
	}

	for range departamentsCount {
		var employeesCount int
		if _, err := fmt.Scan(&employeesCount); err != nil {
			fmt.Println(errors.New("invalid employees number"))
			return
		}

		curCond := newConditioner()

		for range employeesCount {
			var (
				operator string
				temp     int
			)

			if _, err := fmt.Scan(&operator, &temp); err != nil {
				fmt.Println(errors.New("invalid temperature input"))
				return
			}

			result, err := curCond.regulate(operator, temp)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(result)
		}
	}
}
