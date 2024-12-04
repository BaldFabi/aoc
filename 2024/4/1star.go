package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

const searchWord = "XMAS"
const wordLength = len(searchWord) - 1

func toInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func searchForward(matrix []string, r, c int) bool {
	if c+wordLength > len(matrix[r])-1 {
		return false
	}

	for i, char := range searchWord {
		if byte(char) != matrix[r][c+i] {
			return false
		}
	}

	return true
}

func searchBackward(matrix []string, r, c int) bool {
	if c-wordLength < 0 {
		return false
	}

	for i, char := range searchWord {
		if byte(char) != matrix[r][c-i] {
			return false
		}
	}

	return true
}

func searchDown(matrix []string, r, c int) bool {
	if r+wordLength > len(matrix)-1 {
		return false
	}

	for i, char := range searchWord {
		if byte(char) != matrix[r+i][c] {
			return false
		}
	}

	return true
}

func searchUp(matrix []string, r, c int) bool {
	if r-wordLength < 0 {
		return false
	}

	for i, char := range searchWord {
		if byte(char) != matrix[r-i][c] {
			return false
		}
	}

	return true
}

func searchDiagonalDownRight(matrix []string, r, c int) bool {
	if r+wordLength > len(matrix)-1 || c+wordLength > len(matrix[r])-1 {
		return false
	}

	for i, char := range searchWord {
		if byte(char) != matrix[r+i][c+i] {
			return false
		}
	}

	return true
}

func searchDiagonalDownLeft(matrix []string, r, c int) bool {
	if r+wordLength > len(matrix)-1 || c-wordLength < 0 {
		return false
	}

	for i, char := range searchWord {
		if byte(char) != matrix[r+i][c-i] {
			return false
		}
	}

	return true
}

func searchDiagonalUpRight(matrix []string, r, c int) bool {
	if r-wordLength < 0 || c+wordLength > len(matrix[r])-1 {
		return false
	}

	for i, char := range searchWord {
		if byte(char) != matrix[r-i][c+i] {
			return false
		}
	}

	return true
}

func searchDiagonalUpLeft(matrix []string, r, c int) bool {
	if r-wordLength < 0 || c-wordLength < 0 {
		return false
	}

	for i, char := range searchWord {
		if byte(char) != matrix[r-i][c-i] {
			return false
		}
	}

	return true
}

func main() {
	matrix := strings.Split(input, "\n")
	if matrix[len(matrix)-1] == "" {
		matrix = matrix[:len(matrix)-1]
	}

	var sum int
	for r, row := range matrix {
		for c, cell := range row {
			if byte(cell) == searchWord[0] {
				sum += toInt(searchForward(matrix, r, c))
				sum += toInt(searchBackward(matrix, r, c))
				sum += toInt(searchDown(matrix, r, c))
				sum += toInt(searchUp(matrix, r, c))
				sum += toInt(searchDiagonalDownRight(matrix, r, c))
				sum += toInt(searchDiagonalDownLeft(matrix, r, c))
				sum += toInt(searchDiagonalUpRight(matrix, r, c))
				sum += toInt(searchDiagonalUpLeft(matrix, r, c))
			}
		}
	}

	fmt.Println(sum)
}
