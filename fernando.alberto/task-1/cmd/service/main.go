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

func main() {
}
