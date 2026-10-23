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
	fields := strings.Fields(scanner.Text())
	if len(fields) < 2 {
		return
	}
	
	numDepts, err := strconv.Atoi(fields[0])
	if err != nil {
		return
	}
	
	numEmployees, err := strconv.Atoi(fields[1])
	if err != nil {
		return
	}

	minTemp := make([]int, numDepts)
	maxTemp := make([]int, numDepts)
	for i := range minTemp {
		minTemp[i] = 15
		maxTemp[i] = 30
	}

	totalRequests := numDepts * numEmployees
	for reqIdx := 0; reqIdx < totalRequests; reqIdx++ {
		if !scanner.Scan() {
			break
		}
		reqFields := strings.Fields(scanner.Text())
		if len(reqFields) < 2 {
			continue
		}
		
		op := reqFields[0]
		tempVal, err := strconv.Atoi(reqFields[1])
		if err != nil {
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
