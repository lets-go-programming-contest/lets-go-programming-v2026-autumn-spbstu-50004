package main

import "fmt"

func main() {
	var a, b int32
	var op string

	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Invalid first operand")
	}
	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Println("Invalid second operand")
	}
	_, err = fmt.Scan(&op)
	if err != nil {
		fmt.Println("Invalid operation")
	}

	switch op {
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
			fmt.Println(float32(a) / float32(b))
		}
	default:
		fmt.Println("Invalid operation")
	}
}
