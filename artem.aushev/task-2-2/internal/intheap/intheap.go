package intheap

type IntHeap []int

func (h *IntHeap) Len() int {
	if h == nil {
		return 0
	}

	return len(*h)
}

func (h *IntHeap) Less(firstIndex, secondIndex int) bool {
	if h == nil || firstIndex < 0 || firstIndex >= len(*h) || secondIndex < 0 || secondIndex >= len(*h) {
		return false
	}

	return (*h)[firstIndex] > (*h)[secondIndex]
}

func (h *IntHeap) Swap(firstIndex, secondIndex int) {
	if h == nil || firstIndex < 0 || firstIndex >= len(*h) || secondIndex < 0 || secondIndex >= len(*h) {
		return
	}

	(*h)[firstIndex], (*h)[secondIndex] = (*h)[secondIndex], (*h)[firstIndex]
}

func (h *IntHeap) Push(element any) {
	if h == nil {
		return
	}

	integerValue, isInteger := element.(int)
	if !isInteger {
		return
	}

	*h = append(*h, integerValue)
}

func (h *IntHeap) Pop() any {
	if h == nil || len(*h) == 0 {
		return nil
	}

	oldHeapState := *h
	heapLength := len(oldHeapState)
	lastElement := oldHeapState[heapLength-1]
	*h = oldHeapState[0 : heapLength-1]

	return lastElement
}
