package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	departmentsCount, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	for i := 0; i < departmentsCount; i++ {
		if !scanner.Scan() {
			return
		}
		employeesCount, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

		minAllowedTemp := 15
		maxAllowedTemp := 30

		for j := 0; j < employeesCount; j++ {
			if !scanner.Scan() {
				return
			}

			inputParts := strings.Fields(scanner.Text())
			if len(inputParts) != 2 {
				continue
			}

			operator := inputParts[0]
			desiredTemp, _ := strconv.Atoi(inputParts[1])

			if operator == ">=" {
				if desiredTemp > minAllowedTemp {
					minAllowedTemp = desiredTemp
				}
			} else if operator == "<=" {
				if desiredTemp < maxAllowedTemp {
					maxAllowedTemp = desiredTemp
				}
			}

			if minAllowedTemp <= maxAllowedTemp {
				fmt.Println(minAllowedTemp)
			} else {
				fmt.Println("-1")
			}
		}
	}
}
