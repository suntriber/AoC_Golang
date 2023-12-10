package main

import (
	"AoC_Golang/myutils"
	"fmt"
)

func main() {
	input := myutils.ReadStrings("input.txt")
	// fmt.Println(myutils.NumbersFromString(input[0]))
	// fmt.Println(Diff([]int{0, 3, 6, 9, 12, 15}))
	// fmt.Println(extrapolate([]int{0, 3, 6, 9, 12, 15}))
	// fmt.Println(extrapolate([]int{1, 3, 6, 10, 15, 21}))
	fmt.Printf("Part 1: %d\n", part1(input))
}

func part1(input []string) int64 {
	var sum int64
	for _, line := range input {
		sum += int64(extrapolate(myutils.NumbersFromString(line)))
	}
	return sum
}

func extrapolate(nums []int) int {
	n := [][]int{nums}
	i := 1
	for {
		n = append(n, Diff(n[i-1]))
		if Sum(n[i]) == 0 {
			break
		}
		i++
	}

	n[len(n)-1] = append(n[len(n)-1], 0)
	for i := len(n) - 2; i >= 0; i-- {
		n[i] = append(n[i], n[i][len(n[i])-1]+n[i+1][len(n[i+1])-1])
	}

	return n[0][len(n[0])-1]
}

// Diff takes a slice of ints and returns a new slice of the diff between the elements of the incoming slice
func Diff(n []int) []int {
	nums := []int{}
	for i := 1; i < len(n); i++ {
		nums = append(nums, n[i]-n[i-1])
	}
	return nums
}

// Sum takes a slice of ints and returns the sum of all its elements
func Sum(n []int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum
}
