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

	// should be refactored to use logic of star 2
	iterateRows(input, func(row string) {
		var firstHighestNumber rune
		var secondHighestNumber rune

		for i := 0; i < len(row)-1; i++ {
			char := rune(row[i])
			if char > firstHighestNumber {
				firstHighestNumber = char
				secondHighestNumber = 0
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
	joltage := 0

	iterateRows(input, func(row string) {
		var index int
		var j string

		for i := 11; i >= 0; i-- {
			newRow := row[index : len(row)-i]
			newIndex, highestNumber := findHighestNumber(newRow)
			index += newIndex
			j += highestNumber

			fmt.Println("New row:", newRow, "Index:", index, "Highest number:", highestNumber)
		}

		joltageRow, err := strconv.Atoi(j)
		if err != nil {
			panic(err)
		}
		joltage += joltageRow

		fmt.Println("Row joltage:", j, joltageRow)
	})

	fmt.Println("Star 2:", joltage)
}

func findHighestNumber(row string) (int, string) {
	var highestNumber rune
	var index int

	for i, number := range row {
		if number > highestNumber {
			highestNumber = number
			index = i
		}
	}

	return index + 1, string(highestNumber)
}
