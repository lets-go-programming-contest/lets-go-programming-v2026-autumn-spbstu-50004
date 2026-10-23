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
	return (*h)[i] > (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x any) {
	heapValue, isInteger := x.(int)
	if !isInteger {
		return
	}

	*h = append(*h, heapValue)
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	if !scanner.Scan() {
		return
	}

	dishesCount, err := strconv.Atoi(scanner.Text())
	if err != nil || dishesCount <= 0 {
		return
	}

	maxHeap := &IntHeap{}
	heap.Init(maxHeap)

	for range dishesCount {
		if !scanner.Scan() {
			break
		}

		val, parseErr := strconv.Atoi(scanner.Text())
		if parseErr != nil {
			continue
		}

		heap.Push(maxHeap, val)
	}

	if !scanner.Scan() {
		return
	}

	targetRank, err := strconv.Atoi(scanner.Text())
	if err != nil || targetRank <= 0 {
		return
	}

	var result int

	for range targetRank {
		if maxHeap.Len() == 0 {
			break
		}

		popped, isInteger := heap.Pop(maxHeap).(int)
		if !isInteger {
			return
		}

		result = popped
	}

	fmt.Println(result)
}
