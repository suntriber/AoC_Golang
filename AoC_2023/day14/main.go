package main

import (
	"AoC_Golang/myutils"
	"fmt"
	"strings"
)

const (
	ROLLING_ROCK = 'O'
	CUBE_ROCK    = '#'
	DOT          = '.'
	WEST         = "west"
	EAST         = "east"
	NORTH        = "north"
	SOUTH        = "south"
	MAX_CYCLES   = 1000000000
)

func main() {
	input := myutils.ReadStrings("test.txt")
	fmt.Printf("Part1: %d\n", part1(input))
	fmt.Printf("Part2: %d\n", part2(input))
}

func part1(input []string) int {
	sum := 0
	s := tilt(input)
	for i := 0; i < len(s); i++ {
		sum += strings.Count(s[i], "O") * (len(s) - i)
	}
	return sum
}

// tilt tilts a slice in the northern direction
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

// tiltString tilts all rounded rocks left (west)
func tiltString(s string) string {
	tmp := ""
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

func part2(input []string) int {
	cycleDirections := []string{NORTH, WEST, SOUTH, EAST}
	s := input
	for i := 0; i < 1000000000; i++ {
		for _, dir := range cycleDirections {
			s = tiltDirection(s, dir)
		}
		// fmt.Printf("\nAfter %d cycle..\n", i+1)
		// for _, l := range s {
		// 	fmt.Println(l)
		// }
	}

	sum := 0
	for i := 0; i < len(s); i++ {
		sum += strings.Count(s[i], "O") * (len(s) - i)
	}
	return sum
}

func tiltDirection(input []string, dir string) []string {
	output := []string{}
	switch dir {
	case EAST:
		for _, line := range input {
			output = append(output, tiltStringEast(line))
		}
	case WEST:
		for _, line := range input {
			output = append(output, tiltString(line))
		}
	case NORTH:
		output = tilt(input)
	case SOUTH:
		output = tiltSouth(input)
	default:
		panic(fmt.Sprintf("panic!! len(input): %d with dir %s", len(input), dir))
	}

	return output
}

func tiltStringEast(s string) string {
	tmp := ""
	rollCnt := 0
	dotCnt := 0
	for _, c := range s {
		if c == ROLLING_ROCK {
			rollCnt++
		} else if c == DOT {
			dotCnt++
		} else if c == CUBE_ROCK {
			for i := 0; i < dotCnt; i++ {
				tmp += string(DOT)
			}
			for i := 0; i < rollCnt; i++ {
				tmp += string(ROLLING_ROCK)
			}
			tmp += string(CUBE_ROCK)
			rollCnt = 0
			dotCnt = 0
		}
	}
	for i := 0; i < dotCnt; i++ {
		tmp += string(DOT)
	}
	for i := 0; i < rollCnt; i++ {
		tmp += string(ROLLING_ROCK)
	}
	return tmp
}

func tiltSouth(input []string) []string {
	s := []string{}
	for i := 0; i < len(input[0]); i++ {
		tmp := ""
		for j := 0; j < len(input); j++ {
			tmp += string(input[j][i])
		}
		s = append(s, tiltStringEast(tmp))
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
