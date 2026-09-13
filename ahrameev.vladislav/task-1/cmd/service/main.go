package main

import "fmt"

func main() {
	var a, b int
	var operator string

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&operator)

	switch operator {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		fmt.Println(a / b)
	}
}
