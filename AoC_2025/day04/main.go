package main

import (
	"AoC_Golang/AoC_2025/helpers"
	"fmt"
)

const (
	EMPTY  = '.'
	PAPER  = '@'
	TURNED = 'x'
)

func main() {
	data, err := helpers.ReadStrings2D("input.txt")
	if err != nil {
		panic(err)
	}

	part1 := 0

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[0]); j++ {
			if data[i][j] == PAPER {
				count := findNumberAdjacentPaper(i, j, data)
				if count < 4 {
					part1++
				}
			}
		}
	}
	fmt.Println("Part 1:", part1)

	stillGoing := true
	part2 := 0

	for stillGoing {
		count := turnPaperToTurnedIncludingCount(data)
		if count == 0 {
			stillGoing = false
		} else {
			part2 += count
		}
	}

	fmt.Println("Part 2:", part2)
}

func findNumberAdjacentPaper(x, y int, grid [][]rune) int {
	count := 0
	// Check all 8 directions
	directions := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1}}

	for _, dir := range directions {
		nx, ny := x+dir[0], y+dir[1]
		if nx >= 0 && ny >= 0 && nx < len(grid) && ny < len(grid[0]) {
			if grid[nx][ny] == PAPER {
				count++
			}
		}
	}
	return count
}

func turnPaperToTurnedIncludingCount(grid [][]rune) int {
	total := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == PAPER {
				count := findNumberAdjacentPaper(i, j, grid)
				if count < 4 {
					grid[i][j] = TURNED
					total++
				}
			}
		}
	}
	return total
}

// func printGrid(grid [][]rune) {
// 	for i := 0; i < len(grid); i++ {
// 		for j := 0; j < len(grid[0]); j++ {
// 			fmt.Print(string(grid[i][j]))
// 		}
// 		fmt.Println()
// 	}
// }
