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
	fmt.Println("Star 1:")
}

func star2() {
	fmt.Println("Star 2:")
}
