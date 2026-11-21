package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(firstIndex, secondIndex int) bool {
	return (*h)[firstIndex] > (*h)[secondIndex]
}

func (h *IntHeap) Swap(firstIndex, secondIndex int) {
	(*h)[firstIndex], (*h)[secondIndex] = (*h)[secondIndex], (*h)[firstIndex]
}

func (h *IntHeap) Push(element any) {
	integerValue, isInteger := element.(int)
	if !isInteger {
		return
	}

	*h = append(*h, integerValue)
}

func (h *IntHeap) Pop() any {
	oldHeapState := *h
	heapLength := len(oldHeapState)
	lastElement := oldHeapState[heapLength-1]
	*h = oldHeapState[0 : heapLength-1]

	return lastElement
}

func main() {
	var totalDishesCount int
	if _, inputError := fmt.Scan(&totalDishesCount); inputError != nil {
		return
	}

	customMaxHeap := &IntHeap{}

	for range totalDishesCount {
		var currentDishValue int
		if _, readError := fmt.Scan(&currentDishValue); readError != nil {
			return
		}

		*customMaxHeap = append(*customMaxHeap, currentDishValue)
	}

	heap.Init(customMaxHeap)

	var targetRankIndex int
	if _, rankError := fmt.Scan(&targetRankIndex); rankError != nil {
		return
	}

	var finalResult int

	for range targetRankIndex {
		poppedValue, conversionOk := heap.Pop(customMaxHeap).(int)
		if !conversionOk {
			return
		}

		finalResult = poppedValue
	}

	fmt.Println(finalResult)
}
