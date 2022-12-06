package main

import (
	utils "AoC_Golang/myutils"
	"fmt"
)

func main() {
	fmt.Println("Welcome to day 6")
	data := utils.ReadStrings(("input.txt"))[0]
	part1(data)
	part2(data)
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

func part1(data string) {
	s := data[:4]
	for i := 4; i < len(data); i++ {
		s += string(data[i])
		s = s[1:]
		if !hasRepeating(s) {
			fmt.Printf("Part 1 : %d\n", i+1)
			break
		}
	}
}

func part2(data string) {
	s := data[:14]
	for i := 14; i < len(data); i++ {
		s += string(data[i])
		s = s[1:]
		if !hasRepeating(s) {
			fmt.Printf("Part 2 : %d\n", i+1)
			break
		}
	}
}
