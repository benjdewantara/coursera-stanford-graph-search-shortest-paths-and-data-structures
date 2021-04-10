package Heap

import (
	"../Utility"
)

type MinHeap struct {
	Indices []int
	Values  []int
}

func (h *MinHeap) Heapify() {
	var leftmostBoundary int
	for boundaryOrder := 0; boundaryOrder >= 0; boundaryOrder++ {
		leftmostBoundary = 1 << boundaryOrder
		rightmostBoundary := 1 << (boundaryOrder + 1)
		if leftmostBoundary <= len(h.Values) && len(h.Values) <= rightmostBoundary {
			break
		}
	}

	for indx := leftmostBoundary - 1; indx < len(h.Indices); indx++ {
		h.swapWithParent(indx)
	}
}

func (h *MinHeap) ExtractRoot() (int, int) {
	lastIndx := len(h.Values) - 1

	rootIndex := h.Indices[0]
	rootValue := h.Values[0]

	h.Indices[0] = h.Indices[lastIndx]
	h.Values[0] = h.Values[lastIndx]

	h.Indices = h.Indices[0:lastIndx]
	h.Values = h.Values[0:lastIndx]

	h.swapWithChild(0)

	return rootIndex, rootValue
}

func (h *MinHeap) Insert(index int, value int) {
	h.Values = append(h.Values, value)
	h.Indices = append(h.Indices, index)
	h.swapWithParent(len(h.Values) - 1)
}

func (h *MinHeap) UpdateIfLesser(index int, value int) {
	existingAtIndx := Utility.IntArr(h.Indices).IndexOf(index)
	if existingAtIndx < 0 {
		h.Insert(index, value)
		return
	}

	existingValue := h.Values[existingAtIndx]
	if value < existingValue {
		h.Values[existingAtIndx] = value
	}
}

func (h *MinHeap) Delete(indx int) {
	bottomMostIndex := len(h.Values) - 1

	h.swap(indx, bottomMostIndex)
	h.Values = h.Values[0:bottomMostIndex]
	h.Indices = h.Indices[0:bottomMostIndex]

	h.swapWithChild(indx)
}

func (h *MinHeap) swapWithParent(indx int) {
	zeroBasedIndex := ZeroBasedIndex{zeroBasedIndex: indx}

	parentIndx := zeroBasedIndex.GetParentIndex()
	childIndx := indx

	if indx < 0 || parentIndx < 0 {
		return
	}

	parentValue := h.Values[parentIndx]
	childValue := h.Values[childIndx]

	if childValue < parentValue {
		h.swap(parentIndx, childIndx)
	}

	h.swapWithParent(parentIndx)
}

func (h *MinHeap) swapWithChild(indx int) {
	zeroBasedIndex := ZeroBasedIndex{zeroBasedIndex: indx}
	parentIndx := indx

	leftChildIndx := zeroBasedIndex.GetLeftChildIndex()
	if leftChildIndx >= len(h.Values) {
		return
	}

	childIndx := leftChildIndx

	parentValue := h.Values[parentIndx]
	childValue := h.Values[childIndx]

	if childValue < parentValue {
		h.swap(parentIndx, childIndx)
		h.swapWithChild(childIndx)
	}

	rightChildIndx := zeroBasedIndex.GetRightChildIndex()
	if rightChildIndx >= len(h.Values) {
		return
	}

	childIndx = rightChildIndx

	parentValue = h.Values[parentIndx]
	childValue = h.Values[childIndx]

	if childValue < parentValue {
		h.swap(parentIndx, childIndx)
		h.swapWithChild(childIndx)
		return
	}
}

func (h *MinHeap) swap(fromIndx int, toIndx int) {
	valueBefore := h.Values[fromIndx]
	indexBefore := h.Indices[fromIndx]

	h.Values[fromIndx] = h.Values[toIndx]
	h.Indices[fromIndx] = h.Indices[toIndx]

	h.Values[toIndx] = valueBefore
	h.Indices[toIndx] = indexBefore
}
