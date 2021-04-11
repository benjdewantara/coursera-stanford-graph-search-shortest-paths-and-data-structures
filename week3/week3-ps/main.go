package main

import (
	"./Heap"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	maxHeap := Heap.Heap{
		IsMinHeap: false,
		Indices:   make([]int, 0),
		Values:    make([]int, 0),
	}

	minHeap := Heap.Heap{
		IsMinHeap: true,
		Indices:   make([]int, 0),
		Values:    make([]int, 0),
	}

	sumMedians := PopulateHeapsSumMedians(
		"_6ec67df2804ff4b58ab21c12edcb21f8_Median.txt",
		&minHeap,
		&maxHeap)

	fmt.Println(fmt.Sprintf("sumMedians = %d", sumMedians))
	fmt.Println("Hell on earth")
}

func PopulateHeapsSumMedians(
	filepath string,
	minHeap *Heap.Heap,
	maxHeap *Heap.Heap) int {

	sumMedians := 0

	contentBytes, _ := ioutil.ReadFile(filepath)
	for intIndx, intStr := range strings.Split(string(contentBytes), "\r\n") {
		if len(intStr) == 0 {
			continue
		}

		num, _ := strconv.Atoi(intStr)

		if len(minHeap.Values) == 0 {
			minHeap.Insert(intIndx, num)
			sumMedians += (pickMedian(minHeap, maxHeap))
			sumMedians = sumMedians % 10000
			continue
		} else if len(maxHeap.Values) == 0 {
			maxHeap.Insert(intIndx, num)
			sumMedians += pickMedian(minHeap, maxHeap)
			sumMedians = sumMedians % 10000
			continue
		}

		if num <= maxHeap.Values[0] {
			maxHeap.Insert(intIndx, num)
		} else {
			minHeap.Insert(intIndx, num)
		}

		isMinHeapHoldingMore := len(minHeap.Values) >= len(maxHeap.Values)+2
		isMaxHeapHoldingMore := len(maxHeap.Values) >= len(minHeap.Values)+2
		if isMinHeapHoldingMore {
			indx, value := minHeap.ExtractRoot()
			maxHeap.Insert(indx, value)
		} else if isMaxHeapHoldingMore {
			indx, value := maxHeap.ExtractRoot()
			minHeap.Insert(indx, value)
		}

		median := pickMedian(minHeap, maxHeap)
		sumMedians += median
		sumMedians = sumMedians % 10000
	}

	return sumMedians
}

func pickMedian(
	minHeap *Heap.Heap,
	maxHeap *Heap.Heap) int {
	if len(minHeap.Values) >= len(maxHeap.Values)+1 {
		return minHeap.Values[0]
	} else if len(maxHeap.Values) >= len(minHeap.Values)+1 {
		return maxHeap.Values[0]
	} else {
		return maxHeap.Values[0]
	}
}
