package main

import (
	"fmt"
	"strconv"
)

func main() {
	var firstOperandStr, secondOperandStr, operator string

	if _, err := fmt.Scan(&firstOperandStr); err != nil {
		return
	}
	firstOperand, err := strconv.Atoi(firstOperandStr)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	if _, err := fmt.Scan(&secondOperandStr); err != nil {
		return
	}
	secondOperand, err := strconv.Atoi(secondOperandStr)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	if _, err := fmt.Scan(&operator); err != nil {
		return
	}

	if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
		fmt.Println("Invalid operation")
		return
	}
	if operator == "/" && secondOperand == 0 {
		fmt.Println("Division by zero")
		return
	}

	switch operator {
	case "+":
		fmt.Println(firstOperand + secondOperand)
	case "-":
		fmt.Println(firstOperand - secondOperand)
	case "*":
		fmt.Println(firstOperand * secondOperand)
	case "/":
		fmt.Println(firstOperand / secondOperand)
	}
}
