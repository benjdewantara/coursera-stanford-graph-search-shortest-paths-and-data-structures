package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	ReadFilepath("./_6ec67df2804ff4b58ab21c12edcb21f8_algo1-programming_prob-2sum.txt")
}

func ReadFilepath(
	filepath string) int {

	sumMedians := 0

	contentBytes, _ := ioutil.ReadFile(filepath)
	for intIndx, intStr := range strings.Split(string(contentBytes), "\r\n") {
		_ = intIndx
		_ = intStr
	}

	return sumMedians
}
