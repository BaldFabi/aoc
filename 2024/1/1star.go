package main

import (
	_ "embed"
	"fmt"
	"sort"
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

	sort.IntSlice(left).Sort()
	sort.IntSlice(right).Sort()

	for i, l := range left {
		if right[i] > l {
			sum += right[i] - l
		} else {
			sum += l - right[i]
		}
	}

	fmt.Println(sum)
}
