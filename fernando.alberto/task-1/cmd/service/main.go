package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func readInt(scanner *bufio.Scanner, errorMsg string) (int, bool) {
	scanner.Scan()
	valueStr := strings.TrimSpace(scanner.Text())
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		fmt.Println(errorMsg)
		return 0, false
	}
	return value, true
}

func readOperation(scanner *bufio.Scanner) string {
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func calculate(first, second int, op string) (int, bool) {
	switch op {
	case "+":
		return first + second, true
	case "-":
		return first - second, true
	case "*":
		return first * second, true
	case "/":
		if second == 0 {
			fmt.Println("Division by zero")
			return 0, false
		}
		return first / second, true
	default:
		fmt.Println("Invalid operation")
		return 0, false
	}
}

func main() {
}
