package main

import (
	"fmt"
	"strconv"
)

func main() {
	var op1Str, operator, op2Str string

	if _, err := fmt.Scan(&op1Str); err != nil {
		return
	}
	op1, err := strconv.Atoi(op1Str)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	if _, err := fmt.Scan(&operator); err != nil {
		return
	}
	if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
		fmt.Println("Invalid operation")
		return
	}

	if _, err := fmt.Scan(&op2Str); err != nil {
		return
	}
	op2, err := strconv.Atoi(op2Str)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	switch operator {
	case "+":
		fmt.Println(op1 + op2)
	case "-":
		fmt.Println(op1 - op2)
	case "*":
		fmt.Println(op1 * op2)
	case "/":
		if op2 == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(op1 / op2)
	}
}
