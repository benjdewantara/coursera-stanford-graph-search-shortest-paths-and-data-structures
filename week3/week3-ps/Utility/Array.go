package Utility

type IntArr []int

func (arr IntArr) IndexOf(elem int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == elem {
			return i
		}
	}

	return -1
}
