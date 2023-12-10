package main

import (
	"AoC_Golang/myutils"
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	input := myutils.ReadStrings("input.txt")
	fmt.Printf("Part1: %d\n", part1(input))
	fmt.Printf("Part2: %d\n", part2(input))
	// fmt.Println(getNumber(input, 7, 11))
	// fmt.Println(getNumber(input, 8, 11))
	// fmt.Println(getNumber(input, 9, 11))
	// fmt.Println(getNumber(input, 7, 12))
	// fmt.Println(getNumber(input, 8, 12))
	// fmt.Println(getNumber(input, 9, 12))
	// fmt.Println(findAdjacentNumbers(input, []GearNumbers{{XIndexOfStar: 6, YIndexOfStar: 12}}))
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

type GearNumbers struct {
	XIndexOfStar    int
	YIndexOfStar    int
	NumbersAdjacent []int
}

func part2(input []string) int {
	gears := findAdjacentNumbers(input, findStars(input))
	sum := 0
	for _, gear := range gears {
		sum += (gear.NumbersAdjacent[0] * gear.NumbersAdjacent[1])
	}
	return sum
}

func findStars(input []string) []GearNumbers {
	gn := []GearNumbers{}

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if string(input[y][x]) == "*" {
				gn = append(gn, GearNumbers{XIndexOfStar: x, YIndexOfStar: y})
			}
		}
	}
	return gn
}

func findAdjacentNumbers(input []string, gears []GearNumbers) []GearNumbers {
	gn := []GearNumbers{}
	for _, gear := range gears {
		nums := []int{}

		if unicode.IsDigit(rune(input[gear.YIndexOfStar-1][gear.XIndexOfStar])) { // above
			nums = append(nums, getNumber(input, gear.XIndexOfStar, gear.YIndexOfStar-1))
		}
		if !unicode.IsDigit(rune(input[gear.YIndexOfStar-1][gear.XIndexOfStar])) &&
			unicode.IsDigit(rune(input[gear.YIndexOfStar-1][gear.XIndexOfStar-1])) { // up left && !up.isDigit
			nums = append(nums, getNumber(input, gear.XIndexOfStar-1, gear.YIndexOfStar-1))
		}
		if !unicode.IsDigit(rune(input[gear.YIndexOfStar-1][gear.XIndexOfStar])) &&
			unicode.IsDigit(rune(input[gear.YIndexOfStar-1][gear.XIndexOfStar+1])) { // up right && !up.isDigit
			nums = append(nums, getNumber(input, gear.XIndexOfStar+1, gear.YIndexOfStar-1))
		}
		if unicode.IsDigit(rune(input[gear.YIndexOfStar+1][gear.XIndexOfStar])) { // below
			nums = append(nums, getNumber(input, gear.XIndexOfStar, gear.YIndexOfStar+1))
		}
		if !unicode.IsDigit(rune(input[gear.YIndexOfStar+1][gear.XIndexOfStar])) &&
			unicode.IsDigit(rune(input[gear.YIndexOfStar+1][gear.XIndexOfStar-1])) { // below left && !down.isDigit
			nums = append(nums, getNumber(input, gear.XIndexOfStar-1, gear.YIndexOfStar+1))
		}
		if !unicode.IsDigit(rune(input[gear.YIndexOfStar+1][gear.XIndexOfStar])) &&
			unicode.IsDigit(rune(input[gear.YIndexOfStar+1][gear.XIndexOfStar+1])) { // below right && !down.isDigit
			nums = append(nums, getNumber(input, gear.XIndexOfStar+1, gear.YIndexOfStar+1))
		}
		if unicode.IsDigit(rune(input[gear.YIndexOfStar][gear.XIndexOfStar-1])) { // left
			nums = append(nums, getNumber(input, gear.XIndexOfStar-1, gear.YIndexOfStar))
		}
		if unicode.IsDigit(rune(input[gear.YIndexOfStar][gear.XIndexOfStar+1])) { // right
			nums = append(nums, getNumber(input, gear.XIndexOfStar+1, gear.YIndexOfStar))
		}

		if len(nums) == 2 {
			g := GearNumbers{
				XIndexOfStar:    gear.XIndexOfStar,
				YIndexOfStar:    gear.YIndexOfStar,
				NumbersAdjacent: nums,
			}
			gn = append(gn, g)
		}
	}
	return gn
}

func getNumber(input []string, x int, y int) int {
	// fmt.Println("incoming point: x:", x, "y:", y)
	var startIndex int
	var endIndex int
	i := x
	for {
		// fmt.Printf("finding startIndex i: %d\n", i)
		i--
		if i < 0 || !unicode.IsDigit(rune(input[y][i])) {
			break
		}
	}
	startIndex = i + 1
	// fmt.Println("startIndex = ", startIndex)
	i = x
	for {
		// fmt.Printf("Finding endIndex x: %d\n", i)
		i++
		if i > len(input[y])-1 || !unicode.IsDigit(rune(input[y][i])) {
			break
		}
	}
	endIndex = i
	// fmt.Println("endIndex = ", endIndex)
	tmp := ""
	for i := startIndex; i < endIndex; i++ {
		tmp += string(input[y][i])
	}
	n, err := strconv.Atoi(tmp)
	if err != nil {
		panic(err)
	}
	return n
}
