package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
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

func processTask(scanner *bufio.Scanner) {
	if !scanner.Scan() {
		return
	}

	itemCount, err := strconv.Atoi(scanner.Text())
	if err != nil || itemCount <= 0 {
		return
	}

	numberHeap := &IntHeap{}
	heap.Init(numberHeap)

	for range itemCount {
		if !scanner.Scan() {
			break
		}

		val, parseErr := strconv.Atoi(scanner.Text())
		if parseErr != nil {
			continue
		}

		heap.Push(numberHeap, val)
	}

	if !scanner.Scan() {
		return
	}

	k, parseErr := strconv.Atoi(scanner.Text())
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	processTask(scanner)
}
