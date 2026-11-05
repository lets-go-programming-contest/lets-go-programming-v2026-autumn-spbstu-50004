package main

import (
	"container/heap"
	"fmt"
	"os"

	"ahrameev.vladislav/task-2-2/pkg/maxheap"
)

func main() {
	var dishesCount int
	if _, err := fmt.Scan(&dishesCount); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading dishes count: %v\n", err)

		return
	}

	heapData := make(maxheap.MaxHeap, 0, dishesCount)

	for range dishesCount {
		var value int

		if _, err := fmt.Scan(&value); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading value: %v\n", err)

			return
		}

		heapData = append(heapData, value)
	}

	var position int
	if _, err := fmt.Scan(&position); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading position: %v\n", err)

		return
	}

	heap.Init(&heapData)

	var result int

	for range position {
		item, ok := heap.Pop(&heapData).(int)
		if !ok {
			fmt.Fprintf(os.Stderr, "Error: heap is empty\n")

			return
		}

		result = item
	}

	fmt.Println(result)
}
