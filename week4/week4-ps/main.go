package main

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"
)

func main() {
    numbers := NumberMapFromFilepath("./_6ec67df2804ff4b58ab21c12edcb21f8_algo1-programming_prob-2sum.txt")
    usedNumberExistenceCount := make(map[int]int)

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
            }

            _, xUsed := usedNumberExistenceCount[x]
            _, yUsed := usedNumberExistenceCount[y]

            //bothUnused := !xUsed && !yUsed

            xLessThanUsed := !xUsed
            yLessThanUsed := !yUsed

            if xUsed {
                xLessThanUsed = usedNumberExistenceCount[x] < numbers[x]
            }

            if yUsed {
                yLessThanUsed = usedNumberExistenceCount[y] < numbers[y]
            }

            bothXYLessThanUsed := xLessThanUsed && yLessThanUsed

            if bothXYLessThanUsed {
                isEverythingDistinct = true
                usedNumberExistenceCount[x] += 1
                usedNumberExistenceCount[y] += 1
            } else {
                isEverythingDistinct = false
                break
            }
        }

        if isEverythingDistinct {
            countOfTargetValuesFormedByDistinctXAndY++
        }
        usedNumberExistenceCount = make(map[int]int)
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
    //

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
