package main

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

func (c *conditioner) regulate(op string, temp int) int {
	switch op {
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
		return NoTemp
	}

	return c.low
}

func main() {
	var departamentsCount int

	_, err := fmt.Scan(&departamentsCount)
	if err != nil {
		fmt.Println("Invalid department number")

		return
	}

	for range departamentsCount {
		var employeesCount int
		if _, err := fmt.Scan(&employeesCount); err != nil {
			fmt.Println("Invalid employees number")

			return
		}


}
