package main

import "fmt"

const (
	minTemp = 15
	maxTemp = 30
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var k int
		fmt.Scan(&k)

		lo, hi := minTemp, maxTemp

		for j := 0; j < k; j++ {
			var op string
			var value int
			fmt.Scan(&op, &value)

			switch op {
			case ">=":
				if value > lo {
					lo = value
				}
			case "<=":
				if value < hi {
					hi = value
				}
			}

			fmt.Println(lo)
		}
	}
}
