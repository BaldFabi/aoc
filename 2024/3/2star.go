package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	r := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	splits := strings.Split(input, "don't()")

	matches := r.FindAllStringSubmatch(splits[0], -1)
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

	for i := 1; i < len(splits); i++ {
		nSplits := strings.Split(splits[i], "do")

		for n, s := range nSplits {
			if n == 0 {
				continue
			}

			matches := r.FindAllStringSubmatch(s, -1)

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
		}
	}

	fmt.Println(sum)
}
