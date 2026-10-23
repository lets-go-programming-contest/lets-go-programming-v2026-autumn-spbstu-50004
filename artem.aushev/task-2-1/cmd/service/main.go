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

func processRequests(scanner *bufio.Scanner, numDepts, numEmployees int) {
	minTemp := make([]int, numDepts)
	maxTemp := make([]int, numDepts)
	for i := range minTemp {
		minTemp[i] = minDefaultTemp
		maxTemp[i] = maxDefaultTemp
	}

	totalRequests := numDepts * numEmployees
	for reqIdx := range totalRequests {
		if !scanner.Scan() {
			break
		}
		reqFields := strings.Fields(scanner.Text())
		if len(reqFields) < fieldsPerLine {
			continue
		}

		op := reqFields[0]
		tempVal, parseErr := strconv.Atoi(reqFields[1])
		if parseErr != nil {
			continue
		}

		deptIndex := reqIdx / numEmployees

		switch op {
		case ">=":
			if tempVal > minTemp[deptIndex] {
				minTemp[deptIndex] = tempVal
			}
		case "<=":
			if tempVal < maxTemp[deptIndex] {
				maxTemp[deptIndex] = tempVal
			}
		}

		if minTemp[deptIndex] <= maxTemp[deptIndex] {
			fmt.Println(minTemp[deptIndex])
		} else {
			fmt.Println(-1)
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	numDepts := readInt(scanner)
	numEmployees := readInt(scanner)

	processRequests(scanner, numDepts, numEmployees)
}
