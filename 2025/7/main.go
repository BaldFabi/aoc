package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

//go:embed sample.txt
var sample string

func rows(fileContent string) []string {
	return strings.Split(fileContent, "\n")
}

func iterateRows(fileContent string, f func(int, int, string)) {
	rows := rows(fileContent)
	rowLen := len(rows[0])
	for i, row := range rows {
		if row == "" {
			continue
		}

		f(rowLen, i, row)
	}
}

func main() {
	star1(sample)
	star2(sample)
}

// 1770 ^

func star1(data string) {
	var beamSplits int
	var startI int

	iterateRows(data, func(rowLen int, rI int, row string) {
		for cI, char := range row {
			if rI == 0 {
				if char == 'S' {
					startI = cI
					break
				}
			}

			if char == '.' {
				continue
			}

		}
	})

	fmt.Println("Star 1:", beamSplits)
}

func star2(data string) {
	fmt.Println("Star 2:")
}
