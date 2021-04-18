package main

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"
)

func main() {
    numbers := NumberMapFromFilepath("./_6ec67df2804ff4b58ab21c12edcb21f8_algo1-programming_prob-2sum.txt")

    countOfTargetValuesFormedByDistinctXAndY := 0

    for targetValue := -10000; targetValue <= 10000; targetValue++ {
        //for targetValue := -9967; targetValue <= -9967; targetValue++ {
        isEverythingDistinct := false
        for num := range numbers {
            x := num
            y := targetValue - x

            _, xExists := numbers[x]
            _, yExists := numbers[y]

            bothExist := xExists && yExists
            bothDifferent := x != y

            if !bothExist || !bothDifferent {
                continue
            } else {
                isEverythingDistinct = true
                break
            }
        }

        if isEverythingDistinct {
            countOfTargetValuesFormedByDistinctXAndY++
        }
    }

    fmt.Println(
        fmt.Sprintf(
            "countOfTargetValuesFormedByDistinctXAndY = %d", countOfTargetValuesFormedByDistinctXAndY))
}

func NumberMapFromFilepath(
    filepath string) map[int]int {
    numbers := make(map[int]int)

    contentBytes, _ := ioutil.ReadFile(filepath)
    for _, intStr := range strings.Split(string(contentBytes), "\n") {
        num, _ := strconv.Atoi(intStr)
        numbers[num] += 1
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
