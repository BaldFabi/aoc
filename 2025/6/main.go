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
	rs := strings.Split(fileContent, "\n")
	return rs[:len(rs)-1]
}

func iterateRows(_rows []string, fileContent string, f func(row string)) {
	for _, row := range _rows {
		if row == "" {
			continue
		}

		f(row)
	}
}

func main() {
	star1(input)
	star2(input)
}

func star1(data string) {
	_rows := rows(data)
	length := len(_rows)
	signs := strings.Fields(_rows[length-1])
	_rows = _rows[:length-1]

	sums := make([]int, len(signs))

	iterateRows(_rows, sample, func(row string) {
		numbers := strings.Fields(row)

		for i, number := range numbers {
			numberInt, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}

			switch signs[i] {
			case "+":
				sums[i] += numberInt
			case "*":
				if sums[i] == 0 {
					sums[i] = 1
				}
				sums[i] *= numberInt
			}
		}
	})

	var sum int
	for _, s := range sums {
		sum += s
	}

	fmt.Println("Star 1:", sum)
}

func star2(data string) {
	_rows := rows(data)
	length := len(_rows)
	signs := _rows[length-1]
	signs += "e"
	_rows = _rows[:length-1]

	var currentSign rune
	var blockSum int
	var sum int

	for i, sign := range signs {
		var rowNumber string

		if sign != ' ' {
			if blockSum != 0 {
				sum += blockSum
			}

			if sign == 'e' {
				break
			}

			currentSign = sign
			blockSum = 0
			if currentSign == '*' {
				blockSum = 1
			}
		}

		for _, row := range _rows {
			number := string(row[i])
			if strings.TrimSpace(number) == "" {
				continue
			}

			rowNumber += number
		}

		if rowNumber == "" {
			continue
		}

		rowNumberInt, err := strconv.Atoi(rowNumber)
		if err != nil {
			panic(err)
		}

		switch currentSign {
		case '+':
			blockSum += rowNumberInt
		case '*':
			blockSum *= rowNumberInt
		}
	}

	fmt.Println("Star 2:", sum)
}
