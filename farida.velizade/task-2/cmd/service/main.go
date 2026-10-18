package main

import "fmt"

const (
	defaultLow  = 15
	defaultHigh = 30
)

type climateControl struct {
	low  int
	high int
}

func newClimateControl() climateControl {
	return climateControl{low: defaultLow, high: defaultHigh}
}

func (c *climateControl) update(sign string, value int) int {
	switch sign {
	case ">=":
		if value > c.low {
			c.low = value
		}
	case "<=":
		if value < c.high {
			c.high = value
		}
	}

	if c.low > c.high {
		return -1
	}

	return c.low
}

func main() {
	var departments int
	if _, err := fmt.Scan(&departments); err != nil {
		return
	}

	for range departments {
		var employees int
		if _, err := fmt.Scan(&employees); err != nil {
			return
		}

		ctrl := newClimateControl()

		for range employees {
			var (
				sign  string
				value int
			)

			if _, err := fmt.Scan(&sign, &value); err != nil {
				return
			}

			fmt.Println(ctrl.update(sign, value))
		}
	}
}
