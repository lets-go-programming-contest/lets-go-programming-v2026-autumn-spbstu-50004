package main

import "fmt"

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

			if op == ">=" {
				if val > L {
					L = val
				}
			} else if op == "<=" {
				if val < R {
					R = val
				}
			}

			if L <= R {
				fmt.Println(L)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
