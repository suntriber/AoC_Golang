package main

import (
	utils "AoC_Golang/myutils"
	"fmt"
	"sort"
)

func main() {
	input := utils.ReadIntegers("part1.txt")
	part1(input)
	part2(input)
}

func part1(input []int) {
	max := 0
	tmp := 0
	for _, v := range input {
		if v == 0 {
			if tmp > max {
				max = tmp
			}
			tmp = 0
			continue
		}
		tmp += v
	}
	fmt.Printf("Part 1: %d\n", max)
}

func part2(input []int) {
	var total []int
	tmp := 0
	for _, v := range input {
		if v == 0 {
			total = append(total, tmp)
			tmp = 0
		}
		tmp += v
	}
	sort.Ints(total)
	s := len(total) - 1
	fmt.Printf("Part 2: %d\n", total[s]+total[s-1]+total[s-2])
}
