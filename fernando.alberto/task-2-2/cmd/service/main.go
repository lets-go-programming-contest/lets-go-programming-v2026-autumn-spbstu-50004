package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"

	myheap "github.com/AlbertoFernando10/task-2-2/internal/heap"
)

func kthLargest(arr []int, kValue int) int {
	minHeap := &myheap.IntHeap{}
	for _, v := range arr {
		if minHeap.Len() < kValue {
			heap.Push(minHeap, v)
		} else if v > (*minHeap)[0] {
			heap.Pop(minHeap)
			heap.Push(minHeap, v)
		}
	}

	return (*minHeap)[0]
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var arrLen int
	if _, err := fmt.Fscan(reader, &arrLen); err != nil {
		_, _ = fmt.Println("Error:", err)

		return
	}

	arr := make([]int, arrLen)
	for i := range arr {
		if _, err := fmt.Fscan(reader, &arr[i]); err != nil {
			_, _ = fmt.Println("Error:", err)

			return
		}
	}

	var kValue int
	if _, err := fmt.Fscan(reader, &kValue); err != nil {
		_, _ = fmt.Println("Error:", err)

		return
	}

	if kValue < 1 || kValue > arrLen {
		_, _ = fmt.Println(-1)

		return
	}

	_, _ = fmt.Println(kthLargest(arr, kValue))
}
