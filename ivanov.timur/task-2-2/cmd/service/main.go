package main

import "fmt"

type badHeapError string

func (e badHeapError) Error() string {
	return string(e)
}

const ErrBadHeapPushType badHeapError = "IntHeap.Push expects int"

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

func (h *IntHeap) Push(x any) error {
	value, ok := x.(int)
	if !ok {
		return ErrBadHeapPushType
	}

	*h = append(*h, value)

	return nil
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func (h *IntHeap) Sort() {
	length := len(*h)
	for i := range length - 1 {
		for j := range length - 1 - i {
			if (*h)[j] > (*h)[j+1] {
				h.Swap(j, j+1)
			}
		}
	}
}

func main() {
	var amount int

	_, err := fmt.Scan(&amount)
	if err != nil {
		fmt.Println(err)

		return
	}

	intHeap := IntHeap{}

	for range amount {
		var value int

		_, err := fmt.Scan(&value)
		if err != nil {
			fmt.Println(err)

			return
		}

		if err := intHeap.Push(value); err != nil {
			fmt.Println(err)

			return
		}
	}

	var position int

	_, err = fmt.Scan(&position)
	if err != nil {
		fmt.Println(err)

		return
	}

	intHeap.Sort()

	if position <= 0 || position > intHeap.Len() {
		fmt.Println(-1)

		return
	}

	fmt.Println(intHeap[intHeap.Len()-position])
}
