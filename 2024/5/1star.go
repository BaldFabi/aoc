package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func numberIsValid(rules [][]int, n []int) bool {
	var validCombinations int
	l := len(n) - 1

	for i := 0; i < l; i++ {
		for _, r := range rules {
			if n[i] == r[0] && n[i+1] == r[1] {
				validCombinations += 1
				break
			}
		}
	}

	return validCombinations == l
}

func main() {
	rules := [][]int{}
	numbers := [][]int{}

	var i int
	for _, l := range strings.Split(input, "\n") {
		if l == "" {
			continue
		}

		if strings.Contains(l, "|") {
			s := strings.Split(l, "|")
			rules = append(rules, []int{stringToInt(s[0]), stringToInt(s[1])})
		}

		if strings.Contains(l, ",") {
			s := strings.Split(l, ",")
			numbers = append(numbers, []int{})

			for _, n := range s {
				numbers[i] = append(numbers[i], stringToInt(n))
			}

			i += 1
		}
	}

	var sum int
	for _, n := range numbers {
		if numberIsValid(rules, n) {
			sum += n[len(n)/2]
		}
	}

	fmt.Println(sum)
}
