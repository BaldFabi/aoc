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
	star1(input)
	star2(input)
}

func star1(data string) {
	var beamSplits int
	rows := rows(data)

	iterateRows(data, func(rowLen, i int, row string) {
		if i == 0 {
			return
		}

		for cI, char := range row {
			if cI == 0 || cI == len(row)-1 {
				continue
			}

			above := rows[i-1][cI]

			if char == '.' {
				if above == '|' || above == 'S' {
					rows[i] = rows[i][:cI] + "|" + rows[i][cI+1:]
				}

				continue
			}

			if above == '|' {
				beamSplits++

				rows[i] = rows[i][:cI-1] + "|" + rows[i][cI:]
				rows[i] = rows[i][:cI+1] + "|" + rows[i][cI+2:]
			}
		}
	})

	fmt.Println("Star 1:", beamSplits)
}

func star2(data string) {
	var cols []int
	var sum int

	iterateRows(data, func(rowLen, i int, row string) {
		if i%2 == 1 {
			return
		}

		for cI, char := range row {
			if char == '.' {
				continue
			} else if char == 'S' {
				cols = make([]int, len(row))
				cols[cI] = 1
				break
			}

			cols[cI+1] += cols[cI]
			cols[cI-1] += cols[cI]
			cols[cI] = 0
		}
	})

	for _, c := range cols {
		sum += c
	}

	fmt.Println("Star 2:", sum)
}
