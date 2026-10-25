package myheap

import (
	"fmt"
)

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	if (*h).Len() > i && (*h).Len() > j {
		return (*h)[i] > (*h)[j]
	}

	fmt.Println("Out of bound")

	return false
}

func (h *IntHeap) Swap(i, j int) {
	if (*h).Len() > i && (*h).Len() > j {
		(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
	} else {
		fmt.Println("Out of bound")
	}
}

func (h *IntHeap) Push(x any) {
	value, ok := x.(int)
	if !ok {
		return
	}

	*h = append(*h, value)
}

func (h *IntHeap) Pop() any {
	current := *h
	length := len(current)
	result := current[length-1]
	*h = current[0 : length-1]

	return result
}
