package main

import "fmt"

func main() {
	var (
		firstNum, secondNum int
		action              string
	)

	_, err := fmt.Scan(&firstNum)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Scan(&secondNum)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, err = fmt.Scan(&action)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch action {
	case "+":
		fmt.Println(firstNum + secondNum)
	case "-":
		fmt.Println(firstNum - secondNum)
	case "*":
		fmt.Println(firstNum * secondNum)
	case "/":
		if secondNum == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(firstNum / secondNum)
	default:
		fmt.Println("Invalid operation")
	}
}
