package main

import "fmt"

func main() {
	var (
		first    int
		second   int
		operator string
	)
	_, err := fmt.Scan(&first)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	_, err = fmt.Scan(&second)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	_, err = fmt.Scan(&operator)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}
}
