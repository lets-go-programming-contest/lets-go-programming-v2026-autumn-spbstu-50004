package main

import "fmt"

func main() {
	var a, b int
	var operation string

	fmt.Print("Enter first number ")
	fmt.Scan(&a)

	fmt.Print("Enter operation ")
	fmt.Scan(&operation)

	fmt.Print("Enter second number ")
	fmt.Scan(&b)

	switch operation {

	case "+":
		fmt.Println(a + b)

	case "-":
		fmt.Println(a - b)

	case "*":
		fmt.Println(a * b)

	case "/":
		if b == 0 {
			fmt.Println("Division by zero")
		} else {
			fmt.Println(a / b)
		}

	default:
		fmt.Println("Invalid operation")
	}
}
