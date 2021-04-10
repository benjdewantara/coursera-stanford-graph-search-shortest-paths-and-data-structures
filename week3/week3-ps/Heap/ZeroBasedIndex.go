package Heap

type ZeroBasedIndex struct {
	zeroBasedIndex int
}

func (z ZeroBasedIndex) GetParentIndex() int {
	normalizedIndx := z.zeroBasedIndex + 1

	if (normalizedIndx % 2) == 1 {
		normalizedParentIndx := (normalizedIndx - 1) / 2
		return normalizedParentIndx - 1
	}

	normalizedParentIndx := normalizedIndx / 2
	return normalizedParentIndx - 1
}

func (z ZeroBasedIndex) GetLeftChildIndex() int {
	normalizedIndx := z.zeroBasedIndex + 1
	return (normalizedIndx * 2) - 1
}

func (z ZeroBasedIndex) GetRightChildIndex() int {
	normalizedIndx := z.zeroBasedIndex + 1
	return ((normalizedIndx * 2) + 1) - 1
}

func (z ZeroBasedIndex) IsLeftChildItself() bool {
	normalizedIndx := z.zeroBasedIndex + 1

	if (normalizedIndx % 2) == 1 {
		return false
	}

	return true
}
