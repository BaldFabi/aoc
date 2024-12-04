package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

const searchWord = "MAS"
const wordLength = len(searchWord) - 1

func toInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func main() {
	matrix := strings.Split(input, "\n")
	if matrix[len(matrix)-1] == "" {
		matrix = matrix[:len(matrix)-1]
	}

	var sum int
	for r, row := range matrix {
		for c, cell := range row {
			if byte(cell) == searchWord[1] {
				if r == 0 || r == len(matrix)-1 ||
					c == 0 || c == len(row)-1 {
					continue
				}

				upLeft := matrix[r-1][c-1]
				upRight := matrix[r-1][c+1]
				downLeft := matrix[r+1][c-1]
				downRight := matrix[r+1][c+1]

				sum += toInt(upLeft == searchWord[0] && upRight == searchWord[0] &&
					downLeft == searchWord[2] && downRight == searchWord[2])
				sum += toInt(upLeft == searchWord[0] && downLeft == searchWord[0] &&
					upRight == searchWord[2] && downRight == searchWord[2])
				sum += toInt(upRight == searchWord[0] && downRight == searchWord[0] &&
					upLeft == searchWord[2] && downLeft == searchWord[2])
				sum += toInt(downLeft == searchWord[0] && downRight == searchWord[0] &&
					upLeft == searchWord[2] && upRight == searchWord[2])
			}
		}
	}

	fmt.Println(sum)
}
