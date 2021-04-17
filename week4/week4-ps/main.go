package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	numbers := NumberMapFromFilepath("./_6ec67df2804ff4b58ab21c12edcb21f8_algo1-programming_prob-2sum.txt")
	usedNumbers := make(map[int]uint8)

	countOfTargetValuesFormedByDistinctXAndY := 0

	//for targetValue := -10000; targetValue <= 10000; targetValue++ {
	for targetValue := -9967; targetValue <= -9967; targetValue++ {
		for num := range numbers {
			x := num
			y := targetValue - x

			_, xExists := numbers[x]
			_, yExists := numbers[y]

			bothExist := xExists && yExists
			bothDifferent := x != y

			if !bothExist || !bothDifferent {
				continue
			}

			_, xUsed := usedNumbers[x]
			_, yUsed := usedNumbers[y]

			bothUnused := !xUsed && !yUsed

			if bothExist && bothDifferent && bothUnused {
				usedNumbers[x] = 1
				usedNumbers[y] = 1
			}
		}
	}

	fmt.Println(
		fmt.Sprintf(
			"countOfTargetValuesFormedByDistinctXAndY = %d", countOfTargetValuesFormedByDistinctXAndY))

	if numbers != nil {
		fmt.Println("Populated")
	}
}

func NumberMapFromFilepath(
	filepath string) map[int]uint8 {
	numbers := make(map[int]uint8)

	contentBytes, _ := ioutil.ReadFile(filepath)
	for _, intStr := range strings.Split(string(contentBytes), "\n") {
		num, _ := strconv.Atoi(intStr)
		numbers[num] = 1
	}

	return numbers
}

func NumberArrayFromFilepath(
	filepath string) []int {
	numbers := make([]int, 0, 1000000)

	contentBytes, _ := ioutil.ReadFile(filepath)
	for _, intStr := range strings.Split(string(contentBytes), "\n") {
		num, _ := strconv.Atoi(intStr)
		numbers = append(numbers, num)
	}

	return numbers
}
