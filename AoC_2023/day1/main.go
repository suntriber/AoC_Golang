package main

import (
	utils "AoC_Golang/myutils"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("aoc2023 day1")
	input := utils.ReadStrings("part1.txt")

	sum := 0

	for _, line := range input {
		tmp := ""
		for _, r := range line {
			if unicode.IsDigit(r) {
				tmp += string(r)
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				tmp += string(line[i])
				break
			}
		}
		n, _ := strconv.Atoi(tmp)
		sum += n
	}

	fmt.Printf("Part 1: %d\n", sum) // part 1

	sum2 := 0

	nums := map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9", "zero": "0"}

	for i := 0; i < len(input); i++ {
		tmp := ""
		for j := 0; j < len(input[i]); j++ {
			cond := true
			if unicode.IsDigit(rune(input[i][j])) {
				tmp += string(input[i][j])
				break
			}
			for k, v := range nums {
				if strings.HasPrefix(string(input[i][j]), k) {
					tmp += v
					cond = false
					break
				}
			}
			if !cond {
				break
			}
		}

		for j := len(input[i]) - 1; j >= 0; j-- {
			cond := true
			if unicode.IsDigit(rune(input[i][j])) {
				tmp += string(input[i][j])
				break
			}
			for k, v := range nums {
				if strings.HasPrefix(string(input[i][j]), k) {
					tmp += v
					cond = false
					break
				}
			}
			if !cond {
				break
			}
		}

		n, _ := strconv.Atoi(tmp)
		sum2 += n

	}

	fmt.Printf("Part 2: %d\n", sum2)

}
