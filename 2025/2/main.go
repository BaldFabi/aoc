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

	iterateRange(input, func(from, to int) {
		for i := from; i <= to; i++ {
			iStr := strconv.Itoa(i)
			length := len(iStr)
			if length == 1 {
				continue
			}

			halfLength := length / 2

			// --- 11111

			identical := true
			for n := 1; n < length; n++ {
				if iStr[0] != iStr[n] {
					identical = false
					break
				}
			}

			if identical {
				invalidSum += i
				continue
			}

			// --- 1212

			firstHalf := iStr[0:halfLength]
			secondHalf := iStr[halfLength:]

			if firstHalf == secondHalf {
				invalidSum += i
				continue
			}

			// --- 121212 | 824824824

			for n := 2; n < halfLength; n++ {
				if length%n != 0 {
					continue
				}

				identical := true
				firstSequence := iStr[0:n]

				tempHalfLength := halfLength
				if n > 2 {
					tempHalfLength -= 1
				}

				for m := n - (n%2 + 1); m < tempHalfLength; m++ {
					index := m * n
					thisSequence := iStr[index : index+n]
					if firstSequence != thisSequence {
						identical = false
						break
					}
				}

				if identical {
					invalidSum += i
				}

			}
		}
	})

	fmt.Println("Star 2:", invalidSum)
}
