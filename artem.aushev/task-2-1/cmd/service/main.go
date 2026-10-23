package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	minTemp := make([]int, n)
	maxTemp := make([]int, n)
	for i := 0; i < n; i++ {
		minTemp[i] = 15
		maxTemp[i] = 30
	}

	total := n * k
	for i := 0; i < total; i++ {
		var op string
		var t int
		fmt.Scan(&op, &t)

		dept := i / k

		if op == ">=" {
			if t > minTemp[dept] {
				minTemp[dept] = t
			}
		} else if op == "<=" {
			if t < maxTemp[dept] {
				maxTemp[dept] = t
			}
		}

		if minTemp[dept] <= maxTemp[dept] {
			fmt.Println(minTemp[dept])
		} else {
			fmt.Println(-1)
		}
	}
}
