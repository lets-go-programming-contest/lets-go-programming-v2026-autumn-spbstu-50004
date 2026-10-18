package main

import "fmt"

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

		for range employees {
			var (
				sign  string
				value int
			)

			if _, err := fmt.Scan(&sign, &value); err != nil {
				return
			}

			fmt.Println(sign, value)
		}
	}
}
