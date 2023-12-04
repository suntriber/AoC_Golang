package main

import (
	utils "AoC_Golang/myutils"
	"fmt"
	"strconv"
	"strings"
)

const (
	MAX_RED   = 12
	MAX_GREEN = 13
	MAX_BLUE  = 14
)

func main() {
	input := utils.ReadStrings("input.txt")
	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))

}

func part1(input []string) int {
	sum := 0
	for _, line := range input {
		if validateSets(getSets(line)) {
			sum += getGameIDFromLine(line)
		}
	}
	return sum
}

func getGameIDFromLine(s string) int {
	parts := strings.Split(s, " ")
	id := strings.Replace(parts[1], ":", "", 1)
	n, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	return n
}

func getSets(s string) []string {
	setLines := strings.Split(s, ":")[1]
	sets := strings.Split(setLines, ";")
	for i := 0; i < len(sets); i++ {
		sets[i] = sets[i][1:]
	}
	return sets
}

func getAmountOfCubes(s string) int {
	parts := strings.Split(s, " ")
	n, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	return n
}

func validateSets(sets []string) bool {
	for _, set := range sets {
		cubes := strings.Split(set, ",")
		for i := 1; i < len(cubes); i++ {
			cubes[i] = cubes[i][1:]
		}
		for _, c := range cubes {
			if strings.Contains(c, "red") {
				if getAmountOfCubes(c) > MAX_RED {
					return false
				}
			} else if strings.Contains(c, "green") {
				if getAmountOfCubes(c) > MAX_GREEN {
					return false
				}
			} else if strings.Contains(c, "blue") {
				if getAmountOfCubes(c) > MAX_BLUE {
					return false
				}
			}
		}
	}
	return true
}

func part2(input []string) int {
	sum := 0
	for _, line := range input {
		sum += findPowerOfAGame(getSets(line))
	}
	return sum
}

func findPowerOfAGame(sets []string) int {
	red, green, blue := []int{}, []int{}, []int{}
	for _, set := range sets {
		cubes := strings.Split(set, ",")
		for i := 1; i < len(cubes); i++ {
			cubes[i] = cubes[i][1:]
		}
		for _, c := range cubes {
			if strings.Contains(c, "red") {
				red = append(red, getAmountOfCubes(c))
			} else if strings.Contains(c, "green") {
				green = append(green, getAmountOfCubes(c))
			} else if strings.Contains(c, "blue") {
				blue = append(blue, getAmountOfCubes(c))
			}
		}
	}
	return maxValIntSlice(red) * maxValIntSlice(green) * maxValIntSlice(blue)
}

func maxValIntSlice(v []int) (m int) {
	if len(v) > 0 {
		m = v[0]
	}
	for i := 1; i < len(v); i++ {
		if v[i] > m {
			m = v[i]
		}
	}
	return
}
