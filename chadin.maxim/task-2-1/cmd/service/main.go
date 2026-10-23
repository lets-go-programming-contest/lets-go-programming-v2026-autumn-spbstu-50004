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

func (c *Conditioner) Apply(op string, value int) {
	switch op {
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
	var n int

	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	for range n {
		var k int

		if _, err := fmt.Scan(&k); err != nil {
			return
		}

		conditioner := NewConditioner()

		for range k {
			var (
				op    string
				value int
			)

			if _, err := fmt.Scan(&op, &value); err != nil {
				return
			}

			conditioner.Apply(op, value)

			fmt.Println(conditioner.Optimal())
		}
	}
}
