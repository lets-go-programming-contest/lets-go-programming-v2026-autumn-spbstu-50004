package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func readNumbers(scanner *bufio.Scanner, count int) []int {
	nums := make([]int, 0, count)

	for range count {
		if !scanner.Scan() {
			break
		}

		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			continue
		}

		nums = append(nums, val)
	}

	return nums
}

func readSingleInt(scanner *bufio.Scanner) (int, bool) {
	if !scanner.Scan() {
		return 0, false
	}

	val, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return 0, false
	}

	return val, true
}

func processTask(scanner *bufio.Scanner) {
	itemCount, ok := readSingleInt(scanner)
	if !ok || itemCount <= 0 {
		return
	}

	nums := readNumbers(scanner, itemCount)

	k, ok := readSingleInt(scanner)
	if !ok || k <= 0 || k > len(nums) {
		return
	}

	sort.Ints(nums)

	fmt.Println(nums[k-1])
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	processTask(scanner)
}
