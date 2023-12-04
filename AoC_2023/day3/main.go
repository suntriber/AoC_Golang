package main

import (
	"AoC_Golang/myutils"
	"fmt"
	"unicode"
)

func main() {
	input := myutils.ReadStrings("input.txt")
	fmt.Println(input[0])
	part1(input)
}

func part1(input []string) int {
	// check edges
	nums := []string{}
	tmp := ""
	for i := 0; i < len(input[0]); i++ {
		if unicode.IsDigit(rune(input[0][i])) {
			for {
				if unicode.IsDigit(rune(input[0][i])) {
					tmp += string(input[0][i])
				} else if string(input[0][i]) == "." {
					nums = append(nums, tmp)
					tmp = ""
					break
				}
				i++
			}
		}
	}
	fmt.Println(nums)
	return 0
	// check inner
}
