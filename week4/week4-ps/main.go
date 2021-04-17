package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	targetValues := TargetValuesFromFilepath("./_6ec67df2804ff4b58ab21c12edcb21f8_algo1-programming_prob-2sum.txt")
	if targetValues != nil {
		fmt.Println("Populated")
	}
}

func TargetValuesFromFilepath(
	filepath string) []int {
	targetValues := make([]int, 0)

	contentBytes, _ := ioutil.ReadFile(filepath)
	for _, intStr := range strings.Split(string(contentBytes), "\n") {
		num, _ := strconv.Atoi(intStr)
		targetValues = append(targetValues, num)
	}

	return targetValues
}
