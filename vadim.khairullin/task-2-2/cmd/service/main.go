package main

import (
	"container/heap"
	"fmt"

	myheap "github.com/BinaryKenny/task-2-2/cmd/service/internal/myheap"
)

func main() {
	intHeap := &myheap.IntHeap{}
	heap.Init(intHeap)

	count := 0

	_, err := fmt.Scanf("%d", &count)
	if err != nil {
		fmt.Println(err.Error() + ": the format of count is wrong")

		return
	}

	for range count {
		num := 0

		_, err = fmt.Scanf("%d", &num)
		if err != nil {
			fmt.Println(err.Error() + ": the format of number is wrong")

			return
		}

		heap.Push(intHeap, num)
	}

	position := 0

	_, err = fmt.Scanf("%d", &position)
	if err != nil {
		fmt.Println(err.Error() + ": the format of position is wrong")

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
