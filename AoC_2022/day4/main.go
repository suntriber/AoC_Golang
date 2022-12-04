package main

import (
	utils "AoC_Golang/myutils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to day 4")
	data := utils.ReadStrings("input.txt")
	part1(data)
	part2(data)

}

func part1(data []string) {
	total := 0

	for _, line := range data {
		nums := strings.Split(line, ",")
		left := strings.Split(nums[0], "-")
		right := strings.Split(nums[1], "-")
		l1, _ := strconv.Atoi(left[0])
		l2, _ := strconv.Atoi(left[1])
		r1, _ := strconv.Atoi(right[0])
		r2, _ := strconv.Atoi(right[1])

		if ((l1 <= r1) && (l2 >= r2)) || ((r1 <= l1) && (r2 >= l2)) {
			total++
		}
	}
	fmt.Printf("Part 1 : %d\n", total)
}

func part2(data []string) {
	total := 0

	for _, line := range data {
		nums := strings.Split(line, ",")
		left := strings.Split(nums[0], "-")
		right := strings.Split(nums[1], "-")
		l1, _ := strconv.Atoi(left[0])
		l2, _ := strconv.Atoi(left[1])
		r1, _ := strconv.Atoi(right[0])
		r2, _ := strconv.Atoi(right[1])

		if (l1 == r1) || (l2 == r2) || (l1 == r2) || (l2 == r1) {
			total++
		} else if (l1 <= r1 && l2 >= r1) || (r1 <= l1 && r2 >= l1) {
			total++
		}

	}
	fmt.Printf("Part 2 : %d\n", total)
}
