package main

import (
	"AoC_Golang/AoC_2025/helpers"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data, err := helpers.ReadString("input.txt")
	if err != nil {
		panic(err)
	}
	data2 := strings.Split(strings.Trim(data, "\n"), ",")
	// fmt.Println(data2)

	// data2 = []string{"95-115"}

	part1 := 0

	for _, line := range data2 {
		line = strings.TrimSpace(line)
		invalids := getInvalidIDs(line, 1)

		for _, n := range invalids {
			part1 += n
		}
	}
	fmt.Printf("Part 1: %d\n", part1)

	part2 := 0
	for _, line := range data2 {
		line = strings.TrimSpace(line)
		invalids := getInvalidIDs(line, 2)

		for _, n := range invalids {
			part2 += n
		}
	}
	fmt.Printf("Part 2: %d\n", part2)

}

func getRange(s string) (int, int) {
	ss := strings.Split(s, "-")
	n1, err := strconv.Atoi(ss[0])
	if err != nil {
		panic(err)
	}
	n2, err := strconv.Atoi(ss[1])
	if err != nil {
		panic(err)
	}

	return n1, n2
}

func getInvalidIDs(s string, part int) []int {

	min, max := getRange(s)
	nums := []int{}

	for i := min; i <= max; i++ {
		if part == 2 {
			nums = append(nums, hasRepeatedSequencePart2(strconv.Itoa(i)))
			continue
		}
		nums = append(nums, hasRepeatedSequence(strconv.Itoa(i)))
	}

	return nums
}

func hasRepeatedSequence(s string) int {

	ss := s[:len(s)/2]
	sss := s[len(s)/2:]

	if ss == sss {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return n
	}
	return 0
}

func hasRepeatedSequencePart2(s string) int {

	// repeated sequence is made of some sequence of digits repeated **at least twice**
	// only that sequence is considered invalid
	// e.g. 1212 is invalid (12 repeated twice)
	// but 123123 is invalid (123 repeated twice)
	// but 1111 is invalid (1 repeated four times)
	// 110 is not invalid (1 repeated twice but 0 is different)

	// Check all possible pattern lengths that could divide the string
	for l := 1; l <= len(s)/2; l++ {
		// Only check if the pattern length divides evenly into the string length
		if len(s)%l == 0 {
			pattern := s[:l]
			isValid := true

			// Check if the entire string is made of this pattern repeated
			for i := l; i < len(s); i += l {
				if s[i:i+l] != pattern {
					isValid = false
					break
				}
			}

			if isValid {
				n, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}
				return n
			}
		}
	}
	return 0
}
