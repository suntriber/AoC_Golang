package main

import (
	utils "AoC_Golang/myutils"
	"fmt"
)

func main() {
	fmt.Println("Welcome to day 6")
	data := utils.ReadStrings(("input.txt"))[0]
	part1And2(data, 1)
	part1And2(data, 2)
}

func hasRepeating(s string) bool {
	for _, c := range s {
		count := 0
		for _, c2 := range s {
			if c2 == c {
				count++

			}
		}
		if count > 1 {
			return true
		}
	}
	return false
}

func part1And2(data string, part int) {
	j := 0
	if part == 1 {
		j = 4
	} else if part == 2 {
		j = 14
	}
	s := data[:j]
	for i := j; i < len(data); i++ {
		s += string(data[i])
		s = s[1:]
		if !hasRepeating(s) {
			fmt.Printf("Part %d : %d\n", part, i+1)
			break
		}
	}
}
