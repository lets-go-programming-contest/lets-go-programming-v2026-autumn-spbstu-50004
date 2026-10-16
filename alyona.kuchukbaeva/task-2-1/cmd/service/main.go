package main

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

}
