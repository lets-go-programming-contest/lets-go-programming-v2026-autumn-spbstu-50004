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
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var k int
		fmt.Scan(&k)

		conditioner := NewConditioner()

		for j := 0; j < k; j++ {
			var op string
			var value int
			fmt.Scan(&op, &value)

			conditioner.Apply(op, value)
			fmt.Println(conditioner.Optimal())
		}
	}
}
