package main

import (
	"AoC_Golang/myutils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to day {i}")
	input := myutils.ReadStrings("input.txt")
	// input := myutils.ReadStrings("test.txt")
	fmt.Printf("Part 1: %d\n", part1(input))
}

func part1(input []string) int {
	sum := 1
	cycle := 0
	values := map[int]int{20: 0, 60: 0, 100: 0, 140: 0, 180: 0}
	// values := map[int]int{20: 0, 60: 0, 100: 0, 140: 0, 180: 0, 220: 0}
	for i := 0; i < len(input); i++ {
		if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
			tmp := 0
			if input[i-1] != "noop" {
				tmp, _ = strconv.Atoi(strings.Split(input[i-1], " ")[1])
			}
			values[cycle] = sum - tmp
		}
		if input[i] == "noop" {
			cycle++
			continue
		}
		cycle++
		val, _ := strconv.Atoi(strings.Split(input[i], " ")[1])
		if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
			values[cycle] = sum
		}
		cycle++
		sum += val
	}
	total := 0
	for k, v := range values {
		total += k * v
	}
	fmt.Println(values)
	return total
}
