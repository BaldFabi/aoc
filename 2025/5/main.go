package main

import (
	_ "embed"
	"fmt"
	"slices"
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
	var ranges [][]int
	var countFreshIDs int

	iterateRows(input, func(row string) {
		row = strings.TrimSpace(row)
		if row == "" {
			return
		}

		if strings.Contains(row, "-") {
			r := strings.Split(row, "-")
			start, err := strconv.Atoi(r[0])
			if err != nil {
				panic(err)
			}
			end, err := strconv.Atoi(r[1])
			if err != nil {
				panic(err)
			}

			ranges = append(ranges, []int{start, end})

			return
		}

		id, err := strconv.Atoi(row)
		if err != nil {
			panic(err)
		}

		for _, r := range ranges {
			if id >= r[0] && id <= r[1] {
				countFreshIDs++
				break
			}
		}
	})

	fmt.Println("Star 1:", countFreshIDs)
}

func star2() {
	var ranges [][]int
	var countFreshIDs int

	iterateRows(input, func(row string) {
		row = strings.TrimSpace(row)
		if row == "" {
			return
		}

		if !strings.Contains(row, "-") {
			return
		}

		r := strings.Split(row, "-")
		start, err := strconv.Atoi(r[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(r[1])
		if err != nil {
			panic(err)
		}

		ranges = append(ranges, []int{start, end})
	})

	//for {
	//consolidated := consolidateRanges(ranges)
	//if reflect.DeepEqual(consolidated, ranges) {
	//ranges = consolidated
	//break
	//}
	//ranges = consolidated
	//}
	fmt.Printf("Original ranges: %+v\n\n\n", ranges)

	ranges = clearRanges(ranges)

	fmt.Printf("Consolidated ranges: %+v\n", ranges)

	for _, r := range ranges {
		countFreshIDs += r[1] - r[0] + 1
	}

	fmt.Println("Star 2:", countFreshIDs)
}

func clearRanges(ranges [][]int) [][]int {
	slices.SortFunc(ranges, func(a, b []int) int {
		if a[0] < b[0] {
			return -1
		}
		return 1
	})

	var clearedRanges [][]int
	length := len(ranges)

	for i := 0; i < length-1; i++ {
		if ranges[i][1] < ranges[i+1][0] {
			clearedRanges = append(clearedRanges, ranges[i])
			continue
		}

		var start, end int

		start = ranges[i][0]
		end = ranges[i][1]

		for j := i + 1; j < length-1; j++ {
			if end > ranges[j][1] && end > ranges[j+1][0] {
				continue
			}

			end = ranges[j][1]
		}

		clearedRanges = append(clearedRanges, []int{start, end})
	}

	return clearedRanges
}
