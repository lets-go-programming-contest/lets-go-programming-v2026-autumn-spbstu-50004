package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Push(x interface{}) {
	val, ok := x.(int)
	if !ok {
		return
	}

	*h = append(*h, val)
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}

func main() {
	var dishesCount int
	if _, err := fmt.Scan(&dishesCount); err != nil {
		return
	}

	maxHeap := make(MaxHeap, 0, dishesCount)

	for range dishesCount {
		var value int

		if _, err := fmt.Scan(&value); err != nil {
			return
		}

		maxHeap = append(maxHeap, value)
	}

	var position int
	if _, err := fmt.Scan(&position); err != nil {
		return
	}

	heap.Init(&maxHeap)

	var result int

	for range position {
		item, ok := heap.Pop(&maxHeap).(int)
		if !ok {
			return
		}

		result = item
	}

	fmt.Println(result)
}
