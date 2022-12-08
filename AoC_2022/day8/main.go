package main

import (
	utils "AoC_Golang/myutils"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Welcome to day 8")
	data := utils.ReadStrings("input.txt")

	byteToInt := func(b byte) int {
		i, err := strconv.Atoi(string(b))
		if err != nil {
			panic(err)
		}
		return i
	}

	isVisible := func(i, j int) bool {
		tree := byteToInt(data[i][j])
		left, right, up, down := true, true, true, true
		// check left
		for k := j - 1; k >= 0; k-- {
			if tree <= byteToInt(data[i][k]) {
				left = false
				break
			}
		}
		// check right
		for k := j + 1; k < len(data[i]); k++ {
			if tree <= byteToInt(data[i][k]) {
				right = false
				break
			}
		}
		// check down
		for k := i - 1; k >= 0; k-- {
			if tree <= byteToInt(data[k][j]) {
				down = false
				break
			}
		}
		// check up
		for k := i + 1; k < len(data); k++ {
			if tree <= byteToInt(data[k][j]) {
				up = false
				break
			}
		}
		return left || right || up || down
	}

	part1 := len(data)*2 + len(data[0])*2 - 4
	for i := 1; i < len(data)-1; i++ {
		for j := 1; j < len(data[i])-1; j++ {
			if isVisible(i, j) {
				part1 += 1
			}
		}
	}
	fmt.Printf("Part 1 : %d\n", part1)
	fmt.Printf("Part 2 : %d\n", part2(data, byteToInt))
}

func part2(data []string, b2Int func(b byte) int) int {

	getViewingDistance := func(i, j int) int {
		tree := b2Int(data[i][j])
		dLeft, dRight, dDown, dUp := 0, 0, 0, 0
		// check left
		for k := j - 1; k >= 0; k-- {
			if tree <= b2Int(data[i][k]) {
				dLeft++
				break
			} else {
				dLeft++
			}
		}
		// check right
		for k := j + 1; k < len(data[i]); k++ {
			if tree <= b2Int(data[i][k]) {
				dRight++
				break
			} else {
				dRight++
			}
		}
		// check down
		for k := i - 1; k >= 0; k-- {
			if tree <= b2Int(data[k][j]) {
				dDown++
				break
			} else {
				dDown++
			}
		}
		// check up
		for k := i + 1; k < len(data); k++ {
			if tree <= b2Int(data[k][j]) {
				dUp++
				break
			} else {
				dUp++
			}
		}
		if dLeft == 0 {
			dLeft++
		}
		if dRight == 0 {
			dRight++
		}
		if dDown == 0 {
			dDown++
		}
		if dUp == 0 {
			dUp++
		}
		return dLeft * dRight * dDown * dUp
	}
	bestTree := 0
	for i := 1; i < len(data)-1; i++ {
		for j := 1; j < len(data[i])-1; j++ {
			if getViewingDistance(i, j) > bestTree {
				bestTree = getViewingDistance(i, j)
			}
		}
	}
	return bestTree
}
