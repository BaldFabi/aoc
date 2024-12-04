package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type level struct {
	level    []string
	increase bool
}

func toInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func check(increase bool, n []string) bool {
	for i, t := range n {
		if i == len(n)-1 {
			break
		}

		this, err := strconv.Atoi(t)
		if err != nil {
			panic(err)
		}

		next, err := strconv.Atoi(n[i+1])
		if err != nil {
			panic(err)
		}

		if this == next ||
			(increase && this > next || next-this > 3) ||
			(!increase && this < next || this-next > 3) {
			return false
		}
	}

	return true
}

func main() {
	var safe int
	var unsafes []level

	for _, l := range strings.Split(input, "\n") {
		if l == "" {
			continue
		}

		n := strings.Split(l, " ")
		var increase bool

		first, err := strconv.Atoi(n[0])
		if err != nil {
			panic(err)
		}

		last, err := strconv.Atoi(n[len(n)-1])
		if err != nil {
			panic(err)
		}

		if first < last {
			increase = true
		}

		if check(increase, n) {
			safe += 1
		} else {
			unsafes = append(unsafes, level{
				level:    n,
				increase: increase,
			})
		}
	}

	for _, u := range unsafes {
		for i := 0; i < len(u.level); i++ {
			l := make([]string, len(u.level))
			copy(l, u.level)
			copy(l[i:], l[i+1:])
			l = l[:len(l)-1]
			if check(u.increase, l) {
				safe += 1
				break
			}
		}
	}

	fmt.Println(safe)
}
