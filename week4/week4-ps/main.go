package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	numbers := NumbersFromFilepath("./_6ec67df2804ff4b58ab21c12edcb21f8_algo1-programming_prob-2sum.txt")

	if numbers != nil {
		fmt.Println("Populated")
	}
}

func NumbersFromFilepath(
	filepath string) map[int]uint8 {
	numbers := make(map[int]uint8)

	contentBytes, _ := ioutil.ReadFile(filepath)
	for _, intStr := range strings.Split(string(contentBytes), "\n") {
		num, _ := strconv.Atoi(intStr)
		numbers[num] = 1
	}

	return numbers
}
