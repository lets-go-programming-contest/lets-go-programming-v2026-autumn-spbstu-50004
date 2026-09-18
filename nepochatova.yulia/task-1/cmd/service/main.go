package main

import "fmt"

func main() {
	var a, b int
	var operation string

	fmt.Print("Enter first number ")
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	fmt.Print("Enter operation ")
	_, err = fmt.Scan(&operation)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	fmt.Print("Enter second number ")
	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Println("Invalid output")
		return
	}

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
			return
		} else {
			fmt.Println(float64(a) / float64(b))
			return
		}

	default:
		fmt.Println("Invalid operation")
	}
}
