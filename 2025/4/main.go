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

func main() {
	star1(input)
	star2(input)
}

func star1(data string) {
	_rows := rows(data)
	_rows = _rows[:len(_rows)-1]
	colLen := len(_rows)
	rowsLen := len(_rows)
	roll := byte('@')

	var count int

	for rI, row := range _rows {
		for cI := range row {
			if row[cI] != roll {
				continue
			}

			var rowBefore string
			var rowAfter string

			if rI == 0 {
				rowBefore = strings.Repeat(".", colLen)
			} else {
				rowBefore = _rows[rI-1]
			}

			if rI == rowsLen-1 {
				rowAfter = strings.Repeat(".", colLen)
			} else {
				rowAfter = _rows[rI+1]
			}

			var rollsFound int

			indexes := []int{cI - 1, cI, cI + 1}

			for _, i := range indexes {
				if i < 0 || i > len(row)-1 {
					continue
				}

				if rowBefore[i] == roll {
					rollsFound++
				}
				if i != cI && row[i] == roll {
					rollsFound++
				}
				if rowAfter[i] == roll {
					rollsFound++
				}
			}

			if rollsFound < 4 {
				count++
			}
		}
	}

	fmt.Println("Star 1:", count)
}

func star2(data string) {
	_rows := rows(data)
	_rows = _rows[:len(_rows)-1]
	colLen := len(_rows)
	rowsLen := len(_rows)
	roll := byte('@')

	var count int

	for {
		oldCount := count
		for rI, row := range _rows {
			for cI := range row {
				if row[cI] != roll {
					continue
				}

				var rowBefore string
				var rowAfter string

				if rI == 0 {
					rowBefore = strings.Repeat(".", colLen)
				} else {
					rowBefore = _rows[rI-1]
				}

				if rI == rowsLen-1 {
					rowAfter = strings.Repeat(".", colLen)
				} else {
					rowAfter = _rows[rI+1]
				}

				var rollsFound int

				indexes := []int{cI - 1, cI, cI + 1}

				for _, i := range indexes {
					if i < 0 || i > len(row)-1 {
						continue
					}

					if rowBefore[i] == roll {
						rollsFound++
					}
					if i != cI && row[i] == roll {
						rollsFound++
					}
					if rowAfter[i] == roll {
						rollsFound++
					}
				}

				if rollsFound < 4 {
					rowBytes := []byte(_rows[rI])
					rowBytes[cI] = '.'
					_rows[rI] = string(rowBytes)
					count++
				}
			}
		}

		if count == oldCount {
			break
		}
	}

	fmt.Println("Star 2:", count)
}
