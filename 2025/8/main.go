package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed sample.txt
var sample string

func rows(fileContent string) [][]int {
	lines := strings.Split(fileContent, "\n")
	rowsCoords := make([][]int, len(lines)-1)
	for i, line := range lines {
		if line == "" {
			continue
		}
		rowsCoords[i] = toCoords(line)
	}
	return rowsCoords
}

func main() {
	star1(sample)
	star2(sample)
}

func star1(data string) {
	var distances []float64
	distancesWithCoords := make(map[float64]struct {
		coords1 int
		coords2 int
	})
	rows := rows(data)

	for i, coords1 := range rows {
		for j := i + 1; j < len(rows); j++ {
			coords2 := rows[j]
			d := distance(coords1, coords2)

			if _, exists := distancesWithCoords[d]; exists {
				panic("duplicate distance found")
			}

			distances = append(distances, d)
			distancesWithCoords[d] = struct {
				coords1 int
				coords2 int
			}{
				coords1: i,
				coords2: j,
			}
		}
	}

	slices.Sort(distances)

	var junctionBoxes []int

	for i := 0; i < len(rows)-1; i++ {

	}

	fmt.Println("Star 1:")
}

func star2(data string) {
	fmt.Println("Star 2:")
}

func distance(coords1, coords2 []int) float64 {
	if len(coords1) != 3 || len(coords2) != 3 {
		panic("coordinates must have length 3")
	}

	return math.Sqrt(
		math.Pow(float64(coords1[0])-float64(coords2[0]), 2) +
			math.Pow(float64(coords1[1])-float64(coords2[1]), 2) +
			math.Pow(float64(coords1[2])-float64(coords2[2]), 2))
}

func toCoords(s string) []int {
	parts := strings.Split(s, ",")
	if len(parts) != 3 {
		panic("invalid coordinate string: " + s)
	}

	coords := make([]int, len(parts))
	for i, part := range parts {
		coord, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}
		coords[i] = coord

	}
	return coords
}
