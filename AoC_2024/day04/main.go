package main

import (
	"AoC_Golang/AoC_2024/helpers"
	"fmt"
	"strings"
)

const (
	WORD = "XMAS"
)

func main() {
	fmt.Println("day 4..")
	data, err := helpers.ReadStrings("input.txt")
	if err != nil {
		panic(err)
	}

	GRIDSIZE := len(data)

	// test data == 18
	// horizontal : 5
	// vertical   : 3
	// diagonal   : 10

	occurencesHorizontal := 0
	occurencesVertical := 0
	// check horizontal and vertical
	for i := 0; i < GRIDSIZE; i++ {
		s := data[i]
		occurencesHorizontal += strings.Count(s, WORD)
		occurencesHorizontal += strings.Count(reverseString(s), WORD)

		vertical := ""
		for j := 0; j < GRIDSIZE; j++ {
			vertical += string(data[j][i])
		}
		occurencesVertical += strings.Count(vertical, WORD)
		occurencesVertical += strings.Count(reverseString(vertical), WORD)
	}

	occurencesDiagonal := 0
	// Diagonal check
	for i := 0; i < GRIDSIZE; i++ {
		// Top-left to bottom-right
		diagonal1 := ""
		diagonal2 := ""
		for j := 0; j < GRIDSIZE-i; j++ {
			diagonal1 += string(data[j][j+i])
			if i != 0 {
				diagonal2 += string(data[j+i][j])
			}
		}
		occurencesDiagonal += strings.Count(diagonal1, WORD)
		occurencesDiagonal += strings.Count(reverseString(diagonal1), WORD)
		occurencesDiagonal += strings.Count(diagonal2, WORD)
		occurencesDiagonal += strings.Count(reverseString(diagonal2), WORD)

		// Top-right to bottom-left
		diagonal3 := ""
		diagonal4 := ""
		for j := 0; j < GRIDSIZE-i; j++ {
			diagonal3 += string(data[j][GRIDSIZE-1-j-i])
			if i != 0 {
				diagonal4 += string(data[j+i][GRIDSIZE-1-j])
			}
		}
		occurencesDiagonal += strings.Count(diagonal3, WORD)
		occurencesDiagonal += strings.Count(reverseString(diagonal3), WORD)
		occurencesDiagonal += strings.Count(diagonal4, WORD)
		occurencesDiagonal += strings.Count(reverseString(diagonal4), WORD)
	}
	sum := occurencesVertical + occurencesDiagonal + occurencesHorizontal

	fmt.Printf("Part 1: %d\n", sum)

	/*
		.M.S......
		..A..MSMS.
		.M.S.MAA..
		..A.ASMSM.
		.M.S.M....
		..........
		S.S.S.S.S.
		.A.A.A.A..
		M.M.M.M.M.
		..........
	*/

	part2 := 0
	for i := 1; i < GRIDSIZE-1; i++ {
		for j := 1; j < GRIDSIZE-1; j++ {
			if data[i][j] == 'A' {
				if isXmas(data, i, j) {
					// fmt.Printf("(%d, %d)\n", i, j)
					part2++
				}
			}
		}
	}
	fmt.Printf("Part 2: %d\n", part2)
}

func isXmas(d []string, i, j int) bool {
	/*
		top left     : i - 1, j - 1
		bottom left  : i + 1, j - 1

		top right    : i - 1, j + 1
		bottom right : i + 1, j + 1
	*/

	// M  S
	//  A
	// M  S
	if d[i-1][j-1] == 'M' && d[i-1][j+1] == 'S' {
		if d[i+1][j-1] == 'M' && d[i+1][j+1] == 'S' {
			return true
		}
	}

	// S  M
	//  A
	// S  M
	if d[i-1][j-1] == 'S' && d[i-1][j+1] == 'M' {
		if d[i+1][j-1] == 'S' && d[i+1][j+1] == 'M' {
			return true
		}
	}

	// M  M
	//  A
	// S  S
	if d[i-1][j-1] == 'M' && d[i-1][j+1] == 'M' {
		if d[i+1][j-1] == 'S' && d[i+1][j+1] == 'S' {
			return true
		}
	}

	// S  S
	//  A
	// M  M
	if d[i-1][j-1] == 'S' && d[i-1][j+1] == 'S' {
		if d[i+1][j-1] == 'M' && d[i+1][j+1] == 'M' {
			return true
		}
	}

	return false
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
