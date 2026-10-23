package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	itemCountStr := strings.TrimSpace(scanner.Text())
	itemCount, err := strconv.Atoi(itemCountStr)
	if err != nil || itemCount <= 0 {
		return
	}

	if !scanner.Scan() {
		return
	}
	itemsStr := strings.Fields(scanner.Text())

	h := &IntHeap{}
	heap.Init(h)

	for i := 0; i < itemCount && i < len(itemsStr); i++ {
		val, parseErr := strconv.Atoi(itemsStr[i])
		if parseErr != nil {
			continue
		}
		heap.Push(h, val)
	}

	if !scanner.Scan() {
		return
	}
	kStr := strings.TrimSpace(scanner.Text())
	k, parseErr := strconv.Atoi(kStr)
	if parseErr != nil || k <= 0 {
		return
	}

	for i := 1; i < k && h.Len() > 0; i++ {
		heap.Pop(h)
	}

	if h.Len() > 0 {
		fmt.Println((*h)[0])
	}
}
