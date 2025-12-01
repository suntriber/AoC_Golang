package main

import (
	"fmt"
)

func main() {
	fmt.Println("day 10..")
	// data, err := helpers.ReadStrings("test.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// convert to 2d int array
	// dataInts := make([][]int, len(data))
	// for i := 0; i < len(data); i++ {
	// 	tmp := make([]int, len(data[i]))
	// 	for j := 0; j < len(data[i]); j++ {
	// 		intVal, err := strconv.Atoi(string(data[i][j]))
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		tmp[j] = intVal
	// 	}
	// 	dataInts[i] = tmp
	// }
	// fmt.Println(dataInts)
	dataInts := [][]int{
		{0, 1, 2, 3, 4},
		{9, 8, 7, 6, 5},
		{0, 1, 2, 3, 4},
		{9, 8, 7, 6, 5},
		{9, 9, 8, 7, 6},
	}

	MAX_I := len(dataInts)
	MAX_J := len(dataInts[0])

	// get all starting points
	startPoints := make([][]int, 0)
	for i := 0; i < MAX_I; i++ {
		for j := 0; j < MAX_J; j++ {
			tmp := make([]int, 0)
			if dataInts[i][j] == 0 {
				tmp = append(tmp, i, j)
				startPoints = append(startPoints, tmp)
			}
		}
	}
	// fmt.Println(startPoints)

	// winningPoints := make([][]int, 0)

	// for _, pos := range startPoints {
	// 	tmp := make([][]int, 0)
	// 	x, y := pos[1], pos[0]
	// 	tmp = append(tmp, pos)
	// 	currentValue := dataInts[y][x] // first value is 0
	// 	keepGoing := true
	// 	for keepGoing {
	// 		// check all directions
	// 		fmt.Printf("currentValue: [%d, %d] %d\n", y, x, currentValue)

	// 		// up
	// 		if y != 0 && dataInts[y-1][x] == currentValue+1 {
	// 			y--
	// 			currentValue = dataInts[y][x]
	// 			tmp = append(tmp, []int{y, x})
	// 			continue
	// 		}

	// 		// up left
	// 		if y != 0 && x != 0 && dataInts[y-1][x-1] == currentValue+1 {
	// 			y--
	// 			x--
	// 			currentValue = dataInts[y][x]
	// 			tmp = append(tmp, []int{y, x})
	// 			continue
	// 		}

	// 		// up right
	// 		if y != 0 && x != (MAX_J-1) && dataInts[y-1][x+1] == currentValue+1 {
	// 			y--
	// 			x++
	// 			currentValue = dataInts[y][x]
	// 			tmp = append(tmp, []int{y, x})
	// 			continue
	// 		}

	// 		// left
	// 		if x != 0 && dataInts[y][x-1] == currentValue+1 {
	// 			x--
	// 			currentValue = dataInts[y][x]
	// 			tmp = append(tmp, []int{y, x})
	// 			continue
	// 		}

	// 		// right
	// 		if x != (MAX_J-1) && dataInts[y][x+1] == currentValue+1 {
	// 			x++
	// 			currentValue = dataInts[y][x]
	// 			tmp = append(tmp, []int{y, x})
	// 			continue
	// 		}

	// 		// down
	// 		if y != (MAX_I-1) && dataInts[y+1][x] == currentValue+1 {
	// 			y++
	// 			currentValue = dataInts[y][x]
	// 			tmp = append(tmp, []int{y, x})
	// 			continue
	// 		}

	// 		// down left
	// 		if y != (MAX_I-1) && x != 0 && dataInts[y+1][x-1] == currentValue+1 {
	// 			y++
	// 			x--
	// 			currentValue = dataInts[y][x]
	// 			tmp = append(tmp, []int{y, x})
	// 			continue
	// 		}

	// 		// down right
	// 		if y != (MAX_I-1) && x != (MAX_J-1) && dataInts[y+1][x+1] == currentValue+1 {
	// 			y++
	// 			x++
	// 			currentValue = dataInts[y][x]
	// 			tmp = append(tmp, []int{y, x})
	// 			continue
	// 		}

	// 		if currentValue == 9 { // goal
	// 			// winningPoints = append(winningPoints, tmp)
	// 			fmt.Println(tmp)
	// 		}
	// 		keepGoing = false

	// 	}
	// }
	// for _, l := range winningPoints {
	// 	fmt.Println(l)
	// }

	// res := findFullPaths([][]int{
	// 	{0, 1, 2, 3, 4},
	// 	{9, 8, 7, 6, 5},
	// 	{0, 1, 2, 3, 4},
	// 	{9, 8, 7, 6, 5},
	// })
	// fmt.Println(res)

	totalPaths := findPaths(dataInts)
	fmt.Printf("Total number of paths from 0 to 9: %d\n", totalPaths)
}

func findPaths(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	totalPaths := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 {
				totalPaths += dfs(grid, i, j, 0)
			}
		}
		fmt.Println("------------------------------")
	}

	return totalPaths
}

func dfs(grid [][]int, row, col, current int) int {
	fmt.Printf("[%d %d]\n", row, col)
	if current > 9 {
		return 0
	}
	if current == 9 {
		fmt.Println()
		return 1
	}

	original := grid[row][col]
	grid[row][col] = -1 // Mark as visited

	paths := 0
	directions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	for _, dir := range directions {
		newRow, newCol := row+dir[0], col+dir[1]
		if newRow >= 0 && newRow < len(grid) && newCol >= 0 && newCol < len(grid[0]) {
			if grid[newRow][newCol] == current+1 {
				paths += dfs(grid, newRow, newCol, current+1)
			}
		}
	}

	grid[row][col] = original // Restore original value (backtrack)
	return paths
}
