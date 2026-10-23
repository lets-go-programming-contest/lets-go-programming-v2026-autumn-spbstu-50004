package main

import (
	"container/heap"
	"fmt"
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
	value, ok := x.(int)
	if !ok {
		fmt.Println("value is not int")

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

func main() {
	intHeap := &IntHeap{}
	heap.Init(intHeap)

	count := 0

	_, err := fmt.Scanf("%d", &count)
	if err != nil {
		fmt.Println(err)

		return
	}

	for range count {
		num := 0

		_, err = fmt.Scanf("%d", &num)
		if err != nil {
			fmt.Println(err)

			return
		}

		heap.Push(intHeap, num)
	}

	position := 0

	_, err = fmt.Scanf("%d", &position)
	if err != nil {
		fmt.Println(err)

		return
	}

	if position > count {
		fmt.Println("The position is too big if compare with count")

		return
	}

	for range position - 1 {
		heap.Pop(intHeap)
	}

	result := heap.Pop(intHeap)

	fmt.Println(result)
}
