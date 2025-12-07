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

func rows(fileContent string) []string {
	return strings.Split(fileContent, "\n")
}

func iterateRows(fileContent string, f func(row string)) {
	for _, row := range rows(fileContent) {
		if row == "" {
			continue
		}

		f(row)
	}
}

func main() {
	star1()
	star2()
}

func star1() {
	joltage := 0

	iterateRows(input, func(row string) {
		var firstHighestNumber rune
		var secondHighestNumber rune

		for i := 0; i < len(row)-1; i++ {
			char := rune(row[i])
			if char > firstHighestNumber {
				firstHighestNumber = char
				secondHighestNumber = 0
				//highestIndex = i
			} else if char > secondHighestNumber {
				secondHighestNumber = char
			}
		}

		lastNumber := rune(row[len(row)-1])

		if secondHighestNumber < lastNumber {
			secondHighestNumber = lastNumber
		}

		j, err := strconv.Atoi(fmt.Sprintf("%c%c", firstHighestNumber, secondHighestNumber))
		if err != nil {
			panic(err)
		}
		joltage += j
	})

	fmt.Println("Star 1:", joltage)
}

func star2() {
	fmt.Println("Star 2:")
}
