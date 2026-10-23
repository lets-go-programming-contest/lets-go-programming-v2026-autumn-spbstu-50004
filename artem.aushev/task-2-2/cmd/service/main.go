package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x interface{}) {
	val, ok := x.(int)
	if !ok {
		return
	}

	*h = append(*h, val)
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}

func readNextToken(scanner *bufio.Scanner) (string, bool) {
	for {
		if len(scanner.Bytes()) > 0 {
			fields := strings.Fields(scanner.Text())
			if len(fields) > 0 {
				// В реальном потоковом сканировании токенов удобнее использовать 
				// scanner.Split(bufio.ScanWords), но для сохранения структуры 
				// воспользуемся стандартным построчным чтением с правильным разбором.
			}
		}
		if !scanner.Scan() {
			return "", false
		}
	}
}

func parseTargetCount(scanner *bufio.Scanner) (int, bool) {
	if !scanner.Scan() {
		return 0, false
	}

	itemCountStr := strings.TrimSpace(scanner.Text())

	itemCount, err := strconv.Atoi(itemCountStr)
	if err != nil || itemCount <= 0 {
		return 0, false
	}

	return itemCount, true
}

func fillHeap(scanner *bufio.Scanner, numberHeap *IntHeap, itemCount int) {
	wordsRead := 0

	for wordsRead < itemCount && scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		for _, field := range fields {
			if wordsRead >= itemCount {
				break
			}

			val, parseErr := strconv.Atoi(field)
			if parseErr == nil {
				heap.Push(numberHeap, val)
				wordsRead++
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	itemCount, ok := parseTargetCount(scanner)
	if !ok {
		return
	}

	numberHeap := &IntHeap{}
	heap.Init(numberHeap)

	fillHeap(scanner, numberHeap, itemCount)

	if !scanner.Scan() {
		return
	}

	kStr := strings.TrimSpace(scanner.Text())

	k, parseErr := strconv.Atoi(kStr)
	if parseErr != nil || k <= 0 {
		return
	}

	for range k - 1 {
		if numberHeap.Len() == 0 {
			break
		}

		heap.Pop(numberHeap)
	}

	if numberHeap.Len() > 0 {
		fmt.Println((*numberHeap)[0])
	}
}
