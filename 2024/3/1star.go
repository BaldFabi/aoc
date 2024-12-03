package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed input.txt
var input string

func main() {
	r := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := r.FindAllStringSubmatch(input, -1)

	var sum int
	for _, m := range matches {
		f, err := strconv.Atoi(m[1])
		if err != nil {
			panic(err)
		}

		s, err := strconv.Atoi(m[2])
		if err != nil {
			panic(err)
		}

		sum += f * s
	}

	fmt.Println(sum)
}
