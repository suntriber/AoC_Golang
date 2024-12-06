package main

import (
	"AoC_Golang/AoC_2024/helpers"
	"fmt"
)

const (
	guardUp    = '^'
	guardDown  = 'v'
	guardLeft  = '<'
	guardRight = '>'
)

func main() {
	fmt.Println("day 6..")
	data, err := helpers.ReadStrings2D("input.txt")
	if err != nil {
		panic(err)
	}
	startX, startY := findGuardPosition(data)
	// distinctPositions := [][]int{}
	// distinctPositions = append(distinctPositions, []int{startX, startY})
	guardDirection := guardUp

	// boundaries, if we step higher than any of theses we are out of bounds
	maxLengthVertical := len(data) - 1
	maxLengthHorizontal := len(data[0]) - 1

	for inRange(maxLengthVertical, maxLengthHorizontal, startX, startY) {
		// check if next move is available
		switch guardDirection {
		case guardUp:
			if data[startY-1][startX] == '#' {
				guardDirection = turnGuardNinetyDegreesRight(guardDirection)
				continue
			}
		case guardDown:
			if data[startY+1][startX] == '#' {
				guardDirection = turnGuardNinetyDegreesRight(guardDirection)
				continue
			}
		case guardLeft:
			if data[startY][startX-1] == '#' {
				guardDirection = turnGuardNinetyDegreesRight(guardDirection)
				continue
			}
		case guardRight:
			if data[startY][startX+1] == '#' {
				guardDirection = turnGuardNinetyDegreesRight(guardDirection)
				continue
			}
		}
		// clear current position
		data[startY][startX] = 'X'

		// make move
		switch guardDirection {
		case guardUp:
			startY--
		case guardDown:
			startY++
		case guardLeft:
			startX--
		case guardRight:
			startX++
		}

		// enter new guard position
		data[startY][startX] = guardDirection

		// append new positions
		// distinctPositions = append(distinctPositions, []int{startX, startY})
		// printPattern(data)
		// fmt.Println()
	}
	fmt.Printf("Part 1: %d\n", coundDistinct(data)+1)
}

func coundDistinct(d [][]rune) int {
	sum := 0
	for i := 0; i < len(d); i++ {
		for j := 0; j < len(d[i]); j++ {
			if d[i][j] == 'X' {
				sum++
			}
		}
	}
	return sum
}

func printPattern(d [][]rune) {
	for _, l := range d {
		fmt.Println(string(l))
	}
}

func inRange(maxV, maxH, targetX, targetY int) bool {
	return targetY < maxV && targetX < maxH
}

// turnGuardNinetyDegreesRight
func turnGuardNinetyDegreesRight(r rune) rune {
	switch r {
	case guardUp:
		return guardRight
	case guardRight:
		return guardDown
	case guardDown:
		return guardLeft
	case guardLeft:
		return guardUp
	default:
		return '#'
	}
}

// returns the guard position coordinates x, y
func findGuardPosition(d [][]rune) (int, int) {

	for y := 0; y < len(d); y++ {
		for x := 0; x < len(d[y]); x++ {
			if d[y][x] == guardDown ||
				d[y][x] == guardUp ||
				d[y][x] == guardLeft ||
				d[y][x] == guardRight {
				return x, y
			}
		}
	}
	return -1, -1
}
