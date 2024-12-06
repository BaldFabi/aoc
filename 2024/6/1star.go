package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)

func findObstaclesAndStart(m []string) (map[int][]int, []int, int) {
	obstacles := make(map[int][]int)
	start := []int{0, 0}
	direction := UP

	for l, line := range m {
		for c, char := range line {
			if char == '#' {
				obstacles[l] = append(obstacles[l], c)
			}

			switch char {
			case '^':
				direction = UP
			case 'v':
				direction = DOWN
			case '<':
				direction = LEFT
			case '>':
				direction = RIGHT
			}

			if slices.Contains([]rune{'^', 'v', '<', '>'}, char) {
				start = []int{l, c}
			}
		}
	}

	return obstacles, start, direction
}

func findUniqueSpots(m []string, obstacles map[int][]int, start []int, direction int) map[int][]int {
	uniqueSpots := map[int][]int{
		start[0]: []int{start[1]},
	}
	currentPosition := start

	mapWidth := len(m[0])
	mapHeight := len(m)

	for {
		switch direction {
		case UP:
			currentPosition[0] -= 1
		case DOWN:
			currentPosition[0] += 1
		case LEFT:
			currentPosition[1] -= 1
		case RIGHT:
			currentPosition[1] += 1
		}

		if o, ok := obstacles[currentPosition[0]]; ok {
			if slices.Contains(o, currentPosition[1]) {
				switch direction {
				case UP:
					direction = RIGHT
					currentPosition[0] += 1
				case DOWN:
					direction = LEFT
					currentPosition[0] -= 1
				case LEFT:
					direction = UP
					currentPosition[1] += 1
				case RIGHT:
					direction = DOWN
					currentPosition[1] -= 1
				}

				continue
			}
		}

		if currentPosition[0] < 0 || currentPosition[0] > mapWidth ||
			currentPosition[1] < 0 || currentPosition[1] > mapHeight {
			break
		}

		if !slices.Contains(uniqueSpots[currentPosition[0]], currentPosition[1]) {
			uniqueSpots[currentPosition[0]] = append(uniqueSpots[currentPosition[0]], currentPosition[1])
		}
	}

	return uniqueSpots
}

func main() {
	lines := strings.Split(input, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	obstacles, start, direction := findObstaclesAndStart(lines)
	uniqueSpots := findUniqueSpots(lines, obstacles, start, direction)

	var uS int
	for _, v := range uniqueSpots {
		uS += len(v)
	}
	fmt.Println(uS)
}
