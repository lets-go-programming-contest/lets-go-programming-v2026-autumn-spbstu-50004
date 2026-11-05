package maxheap

type MaxHeap []int

func (h *MaxHeap) Len() int {
	if h == nil {
		return 0
	}

	return len(*h)
}

func (h *MaxHeap) Less(i, j int) bool {
	if h == nil || i >= len(*h) || j >= len(*h) {
		return false
	}

	return (*h)[i] > (*h)[j]
}

func (h *MaxHeap) Swap(i, j int) {
	if h == nil || i >= len(*h) || j >= len(*h) {
		return
	}

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
	if h == nil || len(*h) == 0 {
		return nil
	}

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}
