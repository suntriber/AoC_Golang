package main

import (
	"AoC_Golang/myutils"
	"fmt"
	"unicode"
)

func main() {
	input := myutils.ReadStrings("input.txt")
	fmt.Println((part1(input)))
}

func part1(input []string) int {
	nums := []int{}
	for y, line := range input {
		for x := 0; x < len(line); x++ {
			tmp := ""
			if unicode.IsDigit(rune(input[y][x])) {
				b := false
				for {
					if x <= len(input[y])-1 && unicode.IsDigit(rune(input[y][x])) {
						tmp += string(input[y][x])
						if shouldAdd(input, x, y) {
							b = true
						}
					} else if x >= len(line) || !unicode.IsDigit(rune(input[y][x])) {
						if b {
							nums = append(nums, myutils.NumbersFromString(tmp)[0])
						}
						break
					}
					x++
				}
			}
		}
	}
	// fmt.Println(nums)
	return myutils.Sum(nums)
}

func shouldAdd(input []string, x int, y int) bool {
	// above [y-1][x]
	if above(input, x, y) {
		return true
	}
	// below [y+1][x]
	if below(input, x, y) {
		return true
	}
	// left  [y][x-1]
	if left(input, x, y) {
		return true
	}
	// right [y][x+1]
	if right(input, x, y) {
		return true
	}
	// top left [y-1][x-1]
	if topLeft(input, x, y) {
		return true
	}
	// top right [y-1][x+1]
	if topRight(input, x, y) {
		return true
	}
	// bottom left [y+1][x-1]
	if bottomLeft(input, x, y) {
		return true
	}
	// bottom right [y+1][x+1]
	if bottomRight(input, x, y) {
		return true
	}
	return false
}

// above [y-1][x]
func above(input []string, x int, y int) bool {
	if y <= 0 {
		return false
	}
	return string(input[y-1][x]) != "." && !unicode.IsDigit(rune(input[y-1][x]))
}

// below [y+1][x]
func below(input []string, x int, y int) bool {
	if y >= len(input)-1 {
		return false
	}
	return string(input[y+1][x]) != "." && !unicode.IsDigit(rune(input[y+1][x]))
}

// left  [y][x-1]
func left(input []string, x int, y int) bool {
	if x <= 0 {
		return false
	}
	return string(input[y][x-1]) != "." && !unicode.IsDigit(rune(input[y][x-1]))
}

// right [y][x+1]
func right(input []string, x int, y int) bool {
	if x >= len(input[y])-1 {
		return false
	}
	return string(input[y][x+1]) != "." && !unicode.IsDigit(rune(input[y][x+1]))
}

// top left [y-1][x-1]
func topLeft(input []string, x int, y int) bool {
	if y <= 0 || x <= 0 {
		return false
	}
	return string(input[y-1][x-1]) != "." && !unicode.IsDigit(rune(input[y-1][x-1]))
}

// top right [y-1][x+1]
func topRight(input []string, x int, y int) bool {
	if y <= 0 || x >= len(input[y])-1 {
		return false
	}
	return string(input[y-1][x+1]) != "." && !unicode.IsDigit(rune(input[y-1][x+1]))
}

// bottom left [y+1][x-1]
func bottomLeft(input []string, x int, y int) bool {
	if y >= len(input)-1 || x <= 0 {
		return false
	}
	return string(input[y+1][x-1]) != "." && !unicode.IsDigit(rune(input[y+1][x-1]))
}

// bottom right [y+1][x+1]
func bottomRight(input []string, x int, y int) bool {
	if y >= len(input)-1 || x >= len(input[y])-1 {
		return false
	}
	return string(input[y+1][x+1]) != "." && !unicode.IsDigit(rune(input[y+1][x+1]))
}
