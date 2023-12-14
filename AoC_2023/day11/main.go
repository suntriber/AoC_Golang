package main

import (
	"AoC_Golang/myutils"
	"fmt"
	"strings"
)

func main() {
	input := myutils.ReadStrings("test.txt")
	fmt.Printf("Part1: %d\n", part1(input))
}

func part1(input []string) int {

	expandedInput := expandVertically(expandHorizontally(input))
	for _, l := range expandedInput {
		fmt.Println(l)
	}

	// number galaxies
	// determine pairs
	// calc distance between pairs
	// return sum distance

	return 0
}

func expandHorizontally(input []string) []string {
	addingLine := strings.Repeat(".", len(input[0]))
	r := []string{}

	for _, line := range input {
		r = append(r, line)
		if strings.Count(line, ".") == len(line) {
			r = append(r, addingLine)
		}
	}

	return r
}

func expandVertically(input []string) []string {
	lineLength := len(input[0])
	verticalLength := len(input)
	r := make([]string, verticalLength)

	for x := 0; x < lineLength; x++ {
		tmp := ""
		for y := 0; y < verticalLength; y++ {
			tmp += string(input[y][x])
		}
		if strings.Count(tmp, ".") == verticalLength {
			// should add
			for y := 0; y < verticalLength; y++ {
				r[y] += string(tmp[y])
			}
		}
		for y := 0; y < verticalLength; y++ {
			r[y] += string(tmp[y])
		}
	}

	return r
}
