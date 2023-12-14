package main

import (
	"AoC_Golang/myutils"
	"fmt"
	"strings"
)

func main() {
	input := myutils.ReadStrings("input.txt")
	fmt.Printf("Part1: %d\n", part1(input))
}

func part1(input []string) int {
	sum := 0
	s := tilt(input)
	for i := 0; i < len(s); i++ {
		sum += strings.Count(s[i], "O") * (len(s) - i)
	}
	return sum
}

func tilt(input []string) []string {
	s := []string{}
	for i := 0; i < len(input[0]); i++ {
		tmp := ""
		for j := 0; j < len(input); j++ {
			tmp += string(input[j][i])
		}
		s = append(s, tiltString(tmp))
	}

	output := []string{}
	for i := 0; i < len(s[0]); i++ {
		tmp := ""
		for j := 0; j < len(s); j++ {
			tmp += string(s[j][i])
		}
		output = append(output, tmp)
	}
	return output
}

func tiltString(s string) string {
	tmp := ""
	ROLLING_ROCK := 'O'
	CUBE_ROCK := '#'
	DOT := '.'
	rollCnt := 0
	dotCnt := 0
	for _, c := range s {
		if c == ROLLING_ROCK {
			rollCnt++
		} else if c == DOT {
			dotCnt++
		} else if c == CUBE_ROCK {
			for i := 0; i < rollCnt; i++ {
				tmp += string(ROLLING_ROCK)
			}
			for i := 0; i < dotCnt; i++ {
				tmp += string(DOT)
			}
			tmp += string(CUBE_ROCK)
			rollCnt = 0
			dotCnt = 0
		}
	}
	for i := 0; i < rollCnt; i++ {
		tmp += string(ROLLING_ROCK)
	}
	for i := 0; i < dotCnt; i++ {
		tmp += string(DOT)
	}
	return tmp
}
