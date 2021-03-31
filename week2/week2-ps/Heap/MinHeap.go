package Heap

type MinHeap struct {
	Indices []int
	Values  []int
}

func (h *MinHeap) heapify() {
	for i := 0; i < len(h.Indices); i++ {
		parentIndx := i
		leftChildIndx := (2 * (i + 1)) - 1
		rightChildIndx := (2 * (i + 1)) + 1 - 1
	}
}
