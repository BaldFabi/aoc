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

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func numberIsValid(rules map[int][]int, n []int) (bool, int) {
	var validCombinations int
	l := len(n) - 1

	for i := 0; i < l; i++ {
		for _, r := range rules[n[i]] {
			if n[i+1] == r {
				validCombinations += 1
				break
			}
		}
	}

	valid := validCombinations == l

	return valid, n[len(n)/2]
}

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func sw(n1, n2 int, n []int) []int {
	n1Index := slices.Index(n, n1)
	n2Index := slices.Index(n, n2)
	n[n1Index] = n2
	n[n2Index] = n1
	return n
}

func fixNumber(rules map[int][]int, numbers []int) []int {
	l := len(numbers)

	for i := 0; i < l-1; i++ {
		this := numbers[i]
		next := numbers[i+1]

		r1, ok1 := rules[this]
		r2, ok2 := rules[next]

		if !ok1 {
			numbers = sw(this, numbers[l-1], numbers)
			continue
		}

		if !ok2 {
			numbers = sw(next, numbers[l-1], numbers)
			continue
		}

		if !slices.Contains(r1, next) && slices.Contains(r2, this) {
			numbers = sw(this, next, numbers)
			continue
		}
	}

	if ok, _ := numberIsValid(rules, numbers); !ok {
		numbers = fixNumber(rules, numbers)
	}

	return numbers
}

func main() {
	rules := map[int][]int{}
	numbers := [][]int{}

	var i int
	for _, l := range strings.Split(input, "\n") {
		if l == "" {
			continue
		}

		if strings.Contains(l, "|") {
			s := strings.Split(l, "|")
			rules[stringToInt(s[0])] = append(rules[stringToInt(s[0])], stringToInt(s[1]))
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
		valid, _ := numberIsValid(rules, n)
		if !valid {
			n = fixNumber(rules, n)
			sum += n[len(n)/2]
		}
	}

	fmt.Println(sum)
}
