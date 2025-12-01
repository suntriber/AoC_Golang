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
	INPUT      = "input.txt"
)

func main() {
	fmt.Println("day 6..")
	data2, err := helpers.ReadStrings2D(INPUT)
	if err != nil {
		panic(err)
	}
	startX, startY := findGuardPosition(data2)
	guardDirection := guardUp

	// boundaries, if we step higher than any of theses we are out of bounds
	maxLengthVertical := len(data2) - 1
	maxLengthHorizontal := len(data2[0]) - 1

	fmt.Println(maxLengthVertical)
	fmt.Println(maxLengthHorizontal)

	for inRange(maxLengthVertical, maxLengthHorizontal, startX, startY) {
		// check if next move is available
		switch guardDirection {
		case guardUp:
			if data2[startY-1][startX] == '#' {
				guardDirection = turnGuardNinetyDegreesRight(guardDirection)
				continue
			}
		case guardDown:
			if data2[startY+1][startX] == '#' {
				guardDirection = turnGuardNinetyDegreesRight(guardDirection)
				continue
			}
		case guardLeft:
			if data2[startY][startX-1] == '#' {
				guardDirection = turnGuardNinetyDegreesRight(guardDirection)
				continue
			}
		case guardRight:
			if data2[startY][startX+1] == '#' {
				guardDirection = turnGuardNinetyDegreesRight(guardDirection)
				continue
			}
		}
		// clear current position
		data2[startY][startX] = 'X'

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
		data2[startY][startX] = guardDirection

	}
	fmt.Printf("Part 1: %d\n", coundDistinct(data2)+1)

	part2()
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

// debug func
func printPattern(d [][]rune) {
	fmt.Println("==========")
	for _, l := range d {
		fmt.Println(string(l))
	}
	fmt.Println("==========")
}

func inRange(maxV, maxH, targetX, targetY int) bool {
	return targetY < maxV && targetX < maxH && targetX > 0 && targetY > 0
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

func part2() {

	data, err := helpers.ReadStrings2D(INPUT)
	if err != nil {
		panic(err)
	}
	guardDirection := guardUp
	guardDefaultX, guardDefaulY := findGuardPosition(data)

	// boundaries, if we step higher than any of theses we are out of bounds
	maxLengthVertical := len(data) - 1
	maxLengthHorizontal := len(data[0]) - 1

	maxMoves := 24

	loopPoints := [][]int{}

	for i := 0; i < maxLengthVertical+1; i++ {
		for j := 0; j < maxLengthHorizontal+1; j++ {

			// printPattern(data)
			// for i := 0; i < 1; i++ {
			// 	for j := 0; j < 1; j++ {

			// init starting position for guard
			startX, startY := guardDefaultX, guardDefaulY

			// if we already at sign we skip this point
			if data[i][j] == '#' {
				continue
			}

			// place new sign
			if data[i][j] == '.' {
				data[i][j] = '#'
			}

			if i == guardDefaulY && j == guardDefaultX {
				continue
			}

			// count number of times we hit a sign
			up, down, left, right := 0, 0, 0, 0
			// fmt.Printf("============i:%d=================j:%d=====================================\n", i, j)

			// fmt.Printf("(i:%d, j:%d)\n%d %d %d %d\n%d %d\n", i, j, up, down, left, right, startY, startX)
			// walk loop
			for inRange(maxLengthVertical, maxLengthHorizontal, startX, startY) {
				// if i == 9 && j == 7 {
				// 	printPattern(data)
				// }
				// fmt.Printf("(%d, %d)\n", startX, startY)
				switch guardDirection {
				case guardUp:
					if data[startY-1][startX] == '#' {
						up++
						guardDirection = turnGuardNinetyDegreesRight(guardDirection)
						continue
					}
				case guardDown:
					if data[startY+1][startX] == '#' {
						down++
						guardDirection = turnGuardNinetyDegreesRight(guardDirection)
						continue
					}
				case guardLeft:
					if data[startY][startX-1] == '#' {
						left++
						guardDirection = turnGuardNinetyDegreesRight(guardDirection)
						continue
					}
				case guardRight:
					if data[startY][startX+1] == '#' {
						right++
						guardDirection = turnGuardNinetyDegreesRight(guardDirection)
						continue
					}
				}

				// clear current position
				data[startY][startX] = '.'

				//
				if up >= maxMoves && down >= maxMoves && left >= maxMoves && right >= maxMoves {
					loopPoints = append(loopPoints, []int{i, j})
					fmt.Printf("Appended to loop points! (%d, %d)\n", i, j)
					break
				}

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

			}

			// fmt.Printf("(i:%d, j:%d)\n%d %d %d %d\n%d %d\n", i, j, up, down, left, right, startY, startX)

			// reset point we're at
			data[startY][startX] = '.'

			// reset placed sign
			data[i][j] = '.'

			// reset guard
			data[guardDefaulY][guardDefaultX] = '^'
			guardDirection = guardUp

			// printPattern(data)
			// fmt.Println(up, down, left, right)

		}
	}

	fmt.Printf("Part 2: %d\n", len(loopPoints))
}

// test data loop points
// (6,3) (7, 6) (7, 7) (8, 1) (8, 3) (9, 7)
