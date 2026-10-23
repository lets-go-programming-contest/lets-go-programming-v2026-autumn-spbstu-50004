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

var (
	departmentNumberErr = errors.New("invalid department number")
	employeesNumberErr  = errors.New("invalid employees number")
	temperatureInputErr = errors.New("invalid temperature input")
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
	}

	if c.low > c.high {
		return NoTemp, nil
	}

	return c.low, nil
}

func main() {
	var departamentsCount int

	_, err := fmt.Scan(&departamentsCount)
	if err != nil {
		fmt.Println(departmentNumberErr)

		return
	}

	for range departamentsCount {
		var employeesCount int
		if _, err := fmt.Scan(&employeesCount); err != nil {
			fmt.Println(employeesNumberErr)

			return
		}

		curCond := newConditioner()

		for range employeesCount {
			var (
				operator string
				temp     int
			)

			if _, err := fmt.Scan(&operator, &temp); err != nil {
				fmt.Println(temperatureInputErr)

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
