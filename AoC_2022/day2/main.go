package main

import (
	utils "AoC_Golang/myutils"
	"fmt"
)

func main() {
	// A X = ROCK (1)
	// B Y = PAPER (2)
	// C Z = SCISSORS (3)
	input := utils.ReadStrings("part1.txt")
	part1(input)
	part2(input)

}

func part1(input []string) {
	var totalScore int
	for i := 0; i < len(input); i++ {
		totalScore += calculateScorePart1(input[i])
	}
	fmt.Printf("Part 1: %d\n", totalScore)
}

func calculateScorePart1(game string) int {
	switch game {
	case "A X":
		return 1 + 3 // rock draw
	case "A Y":
		return 2 + 6 // paper vs rock win
	case "A Z":
		return 3 + 0 // scissors vs rock loose
	case "B X":
		return 1 + 0 // rock vs paper loose
	case "B Y":
		return 2 + 3 // paper draw
	case "B Z":
		return 3 + 6 // scissors vs paper win
	case "C X":
		return 1 + 6 // rock vs scissors win
	case "C Y":
		return 2 + 0 // paper vs scissors loose
	case "C Z":
		return 3 + 3 // scissors draw
	default:
		return 0
	}
}

func part2(input []string) {
	var totalScore int
	for i := 0; i < len(input); i++ {
		totalScore += calculateScorePart2(input[i])
	}
	fmt.Printf("Part 2: %d\n", totalScore)
}

func calculateScorePart2(game string) int {
	// X = LOOSE
	// Y = DRAW
	// Z = WIN
	// A = ROCK (1)
	// B = PAPER (2)
	// C = SCISSORS (3)
	switch game {
	case "A X":
		return 3 + 0
	case "A Y":
		return 1 + 3
	case "A Z":
		return 2 + 6
	case "B X":
		return 1 + 0
	case "B Y":
		return 2 + 3
	case "B Z":
		return 3 + 6
	case "C X":
		return 2 + 0
	case "C Y":
		return 3 + 3
	case "C Z":
		return 1 + 6
	default:
		return 0
	}
}
