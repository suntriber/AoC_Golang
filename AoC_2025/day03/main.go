package main

import (
	"AoC_Golang/AoC_2024/helpers"
	"fmt"
	"strconv"
)

func main() {
	data, err := helpers.ReadStrings("input.txt")
	if err != nil {
		panic(err)
	}

	// fmt.Println(data)
	// data = []string{"811111111111119"}
	// data = []string{"818181911112111"}
	// fmt.Println(len(data[0]))
	// data = []string{"987654321111111"}

	// b, i := findHighestRune(data[1], 0)
	// b2, i2 := findHighestRune(data[1], i+1)

	// fmt.Println(string(b), i)
	// fmt.Println(string(b2), i2)

	part1 := 0

	for _, line := range data {
		n1, i1 := findHighestRune(line, 0, false)
		n2, _ := findHighestRune(line, i1+1, true)

		s := fmt.Sprintf("%c%c", n1, n2)
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		part1 += num
	}
	fmt.Println("Part 1:", part1)

	part2 := 0
	for _, line := range data {
		s := largestNumberSubsequence(line, 12)
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		part2 += num
	}
	fmt.Println("Part 2:", part2)

}

func findHighestRune(s string, start int, wantLast bool) (byte, int) {
	lastIndex := len(s) - 1
	highest := s[start]
	index := start

	for i, r := range s[start:] {
		if r > rune(highest) {
			if start+i == lastIndex && !wantLast {
				break
			} else {
				highest = byte(r)
				index = start + i
			}
		}
	}
	return highest, index
}

func largestNumberSubsequence(s string, n int) string {
	removals := len(s) - n
	stack := []rune{}

	for _, c := range s {
		for removals > 0 && len(stack) > 0 && stack[len(stack)-1] < c {
			stack = stack[:len(stack)-1]
			removals--
		}
		stack = append(stack, c)
	}

	// If removals still remain, remove from the end
	if removals > 0 {
		stack = stack[:len(stack)-removals]
	}

	return string(stack)
}
