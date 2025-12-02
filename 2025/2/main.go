package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed sample.txt
var sample string

func ranges(fileContent string) [][]int {
	ranges := strings.Split(fileContent, ",")
	var result [][]int

	for _, r := range ranges {
		rangeStrings := strings.Split(r, "-")
		from, err := strconv.Atoi(strings.TrimSpace(rangeStrings[0]))
		if err != nil {
			panic(err)
		}

		to, err := strconv.Atoi(strings.TrimSpace(rangeStrings[1]))
		if err != nil {
			panic(err)
		}

		result = append(result, []int{from, to})
	}

	return result
}

func iterateRange(fileContent string, f func(int, int)) {
	for _, r := range ranges(fileContent) {
		f(r[0], r[1])
	}
}

func main() {
	star1()
	star2()
}

func star1() {
	invalidSum := 0

	iterateRange(input, func(from, to int) {
		for i := from; i <= to; i++ {
			iStr := strconv.Itoa(i)
			length := len(iStr)
			if length%2 != 0 {
				continue
			}

			//fmt.Println(i, iStr, iStr[0:length/2], iStr[length/2:])
			if iStr[0:length/2] == iStr[length/2:] {
				invalidSum += i
			}
		}

	})

	fmt.Println("Star 1:", invalidSum)
}

func star2() {
	invalidSum := 0
	fmt.Println("Star 2:", invalidSum)
}
