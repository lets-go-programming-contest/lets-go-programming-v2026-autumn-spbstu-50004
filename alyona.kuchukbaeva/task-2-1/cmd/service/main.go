package main

import "fmt"

const (
	minTemp = 15
	maxTemp = 30
)

type conditioner struct {
	low  int
	high int
}

func newConditioner() conditioner {
	return conditioner{low: minTemp, high: maxTemp}
}

func (c *conditioner) apply(sign string, value int) {
	switch sign {
	case ">=":
		c.low = max(c.low, value)
	case "<=":
		c.high = min(c.high, value)
	}
}

func (c *conditioner) optimal() (int, bool) {
	if c.low > c.high {
		return 0, false
	}

	return c.low, true
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

		cond := newConditioner()

		for range employees {
			var (
				sign  string
				value int
			)

			if _, err := fmt.Scan(&sign, &value); err != nil {
				return
			}

			cond.apply(sign, value)

			if temp, ok := cond.optimal(); ok {
				fmt.Println(temp)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
