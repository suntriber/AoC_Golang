package main

import (
	"AoC_Golang/myutils"
	"fmt"
)

const (
	STARTNODE = "AAA"
	ENDNODE   = "ZZZ"
	RIGHT     = "R"
	LEFT      = "L"
)

func main() {
	input := myutils.ReadStrings("input.txt")
	// fmt.Printf("Part1: %d\n", part1(input))
	fmt.Printf("Part2: %d\n", part2(input))
	fmt.Println(getStartingPositions(input))
	// fmt.Println(makeOneStep(2, input, "L"))
	// fmt.Println(allWins(input, []int{4, 8}))
	// fmt.Println(allWins(input, []int{4, 9}))
}

func part1(input []string) int {
	commands := input[0]
	steps := 0
	currIndex := 2

	for i := 2; i < len(input); i++ {
		n, _, _ := getNodeAndElements(input[i])
		if n == STARTNODE {
			currIndex = i
			break
		}
	}

	for {
		for _, c := range commands {
			node, eleA, eleB := getNodeAndElements(input[currIndex])
			if node == ENDNODE {
				return steps
			}
			var nodeToGo string
			if string(c) == RIGHT {
				nodeToGo = eleB
			} else {
				nodeToGo = eleA
			}
			for {
				currIndex++
				if currIndex >= len(input) {
					currIndex = 2
				}
				newNode, _, _ := getNodeAndElements(input[currIndex])
				if newNode == nodeToGo {
					steps++
					break
				}
			}
		}
	}
}

func getNodeAndElements(s string) (node, eleA, eleB string) {
	node = s[0:3]
	eleA = s[7:10]
	eleB = s[12:15]
	return
}

func part2(input []string) int {
	commands := input[0]
	totalSteps := 0
	startingPositions := getStartingPositions(input)
	for {
		for _, c := range commands {
			for i := 0; i < len(startingPositions); i++ {
				newPlace := makeOneStep(startingPositions[i], input, string(c))
				startingPositions[i] = newPlace
			}
			totalSteps++
			if allWins(input, startingPositions) {
				return totalSteps
			}
		}
	}
}

func allWins(input []string, positions []int) bool {
	win := true
	for i := 0; i < len(positions); i++ {
		if string(input[positions[i]][2]) != "Z" {
			win = false
		}
	}
	return win
}

func makeOneStep(place int, input []string, cmd string) (newPlace int) {
	for {
		_, eleA, eleB := getNodeAndElements(input[place])
		fmt.Printf("[%d] (%s, %s)\n", place, eleA, eleB)
		place++
		if place >= len(input) {
			place = 2
		}
		nodeToGo := ""
		if cmd == RIGHT {
			nodeToGo = eleB
		} else {
			nodeToGo = eleA
		}
		for {
			newNode, _, _ := getNodeAndElements(input[place])
			if newNode == nodeToGo {
				return place
			}
			place++
			if place >= len(input) {
				place = 2
			}
		}
	}
}

func getStartingPositions(input []string) []int {
	indexes := []int{}
	for i := 2; i < len(input); i++ {
		if string(input[i][2]) == "A" {
			indexes = append(indexes, i)
		}
	}
	return indexes
}
