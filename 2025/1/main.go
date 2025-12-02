package main

import (
	_ "embed"
	"fmt"
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

var currentPosition = 0

func main() {
	currentPosition = 50
	star1()
	currentPosition = 50
	star2()
}

func star1() {
	hitZero := 0

	iterateRows(input, func(row string) {
		direction, clicks := getDirectionAndClicks(row)

		//fmt.Println("Current Position:", currentPosition, "Direction:", direction, "Clicks:", clicks)
		clicks = clicks % 100

		switch direction {
		case "L":
			if clicks > currentPosition {
				currentPosition = currentPosition + 100 - clicks
			} else {
				currentPosition -= clicks
			}
		case "R":
			if currentPosition+clicks > 99 {
				currentPosition = currentPosition - 100 + clicks
			} else {
				currentPosition += clicks
			}
		}

		if currentPosition == 0 {
			hitZero++
		}
	})

	fmt.Println("Star 1:", hitZero)
}

func star2() {
	hitZero := 0

	iterateRows(input, func(row string) {
		direction, clicks := getDirectionAndClicks(row)
		previousPosition := currentPosition

		//fmt.Println("Current Position:", currentPosition, "Direction:", direction, "Clicks:", clicks)
		hitZero += clicks / 100
		clicks = clicks % 100

		switch direction {
		case "L":
			if clicks > currentPosition {
				currentPosition = currentPosition + 100 - clicks
				if previousPosition != 0 {
					hitZero++
					return
				}
			} else {
				currentPosition -= clicks
			}
		case "R":
			if currentPosition+clicks > 99 {
				currentPosition = currentPosition - 100 + clicks
				if previousPosition != 0 {
					hitZero++
					return
				}
			} else {
				currentPosition += clicks
			}
		}

		if currentPosition == 0 {
			hitZero++
		}
	})

	fmt.Println("Star 2:", hitZero)
}

func getDirectionAndClicks(row string) (string, int) {
	direction := string(row[0])
	clicks, err := strconv.Atoi(row[1:])
	if err != nil {
		panic(err)
	}
	return direction, clicks
}
