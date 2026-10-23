package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	minDefaultTemp = 15
	maxDefaultTemp = 30
	fieldsPerLine  = 2
)

func readInt(scanner *bufio.Scanner) int {
	if !scanner.Scan() {
		return 0
	}

	val, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		return 0
	}

	return val
}

func processDepartmentRequests(scanner *bufio.Scanner, numEmployees int) {
	minTemp := minDefaultTemp
	maxTemp := maxDefaultTemp

	for i := 0; i < numEmployees; i++ {
		if !scanner.Scan() {
			break
		}

		reqFields := strings.Fields(scanner.Text())
		if len(reqFields) < fieldsPerLine {
			continue
		}

		operatorName := reqFields[0]

		tempVal, parseErr := strconv.Atoi(reqFields[1])
		if parseErr != nil {
			continue
		}

		switch operatorName {
		case ">=":
			if tempVal > minTemp {
				minTemp = tempVal
			}
		case "<=":
			if tempVal < maxTemp {
				maxTemp = tempVal
			}
		}

		if minTemp <= maxTemp {
			fmt.Println(minTemp)
		} else {
			fmt.Println(-1)
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	numDepts := readInt(scanner)

	for i := 0; i < numDepts; i++ {
		numEmployees := readInt(scanner)

		processDepartmentRequests(scanner, numEmployees)
	}
}
