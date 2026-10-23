package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	current := *h
	length := len(current)
	result := current[length-1]
	*h = current[0 : length-1]

	return result
}

func main() {
	h := &IntHeap{}
	heap.Init(h)

	count := 0

	_, error := fmt.Scanf("%d", &count)
	if error != nil {
		fmt.Println(error)

		return
	}

	for i := 0; i < count; i++ {
		num := 0

		_, error := fmt.Scanf("%d", &num)
		if error != nil {
			fmt.Println(error)

			return
		}

		heap.Push(h, num)
	}

	k := 0
	_, err := fmt.Scanf("%d", &k)
	if err != nil {
		fmt.Println(err)

		return
	}

	if k > count {
		fmt.Println("The position is too big if compare with count")

		return
	}

	for i := 0; i < k-1; i++ {
		heap.Pop(h)
	}

	result := heap.Pop(h)

	fmt.Println(result)
}
