package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var sum int
	var left, right []int

	rows := strings.Split(input, "\n")
	for _, row := range rows {
		cols := strings.Split(row, " ")
		for i, col := range cols {
			if col == "" {
				continue
			}

			if v, err := strconv.Atoi(col); err == nil {
				if i == 0 {
					left = append(left, v)
				} else {
					right = append(right, v)
				}
			} else {
				panic(err)
			}
		}
	}

	for _, l := range left {
		var count int

		for _, r := range right {
			if l == r {
				count++
			}
		}

		sum += count * l
	}

	fmt.Println(sum)
}
