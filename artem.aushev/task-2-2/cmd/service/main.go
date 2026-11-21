package main

import (
	"container/heap"
	"fmt"
	"log"
	"os"

	"task-2-2/internal/intheap"
)

func main() {
	var totalDishesCount int
	if _, inputError := fmt.Scan(&totalDishesCount); inputError != nil {
		log.Println(fmt.Errorf("failed to read total dishes count: %w", inputError))
		return
	}

	customMaxHeap := &intheap.IntHeap{}

	for range totalDishesCount {
		var currentDishValue int
		if _, readError := fmt.Scan(&currentDishValue); readError != nil {
			log.Println(fmt.Errorf("failed to read dish value: %w", readError))
			return
		}

		*customMaxHeap = append(*customMaxHeap, currentDishValue)
	}

	heap.Init(customMaxHeap)

	var targetRankIndex int
	if _, rankError := fmt.Scan(&targetRankIndex); rankError != nil {
		log.Println(fmt.Errorf("failed to read target rank index: %w", rankError))
		return
	}

	if targetRankIndex <= 0 || targetRankIndex > customMaxHeap.Len() {
		log.Println(fmt.Errorf("target rank index %d is out of bounds for heap length %d", targetRankIndex, customMaxHeap.Len()))
		return
	}

	var finalResult int

	for range targetRankIndex {
		poppedValue, conversionOk := heap.Pop(customMaxHeap).(int)
		if !conversionOk {
			log.Println(fmt.Errorf("failed to convert popped heap element to int"))
			return
		}

		finalResult = poppedValue
	}

	fmt.Println(finalResult)
}
