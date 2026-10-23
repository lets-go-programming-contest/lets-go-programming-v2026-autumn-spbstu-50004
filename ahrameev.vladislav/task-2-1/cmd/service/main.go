package main

import "fmt"

func updateRange(L, R int, op string, val int) (int, int) {
	if op == ">=" && val > L {
		L = val
	} else if op == "<=" && val < R {
		R = val
	}
	return L, R
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var k int
		fmt.Scan(&k)

		L, R := 15, 30

		for j := 0; j < k; j++ {
			var op string
			var val int
			fmt.Scan(&op, &val)

			L, R = updateRange(L, R, op, val)

			if L <= R {
				fmt.Println(L)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
