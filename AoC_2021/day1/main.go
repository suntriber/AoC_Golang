package main

import (
	"fmt"

	utils "AoC_Golang/myutils"
)

func main() {
	filePath := "input.txt"

	listIntegers := utils.ReadIntegers(filePath)

	current := listIntegers[0]
	total := 0

	for _, n := range listIntegers {
		if n > current {
			total++
		}
		current = n
	}

	fmt.Printf("Part 1: %d\n", total) // part 1

	listIntegers = utils.ReadIntegers("input2.txt")
	current = listIntegers[0] + listIntegers[1] + listIntegers[2]
	total = 0

	for i := 0; i < len(listIntegers)-3; i++ {
		if current < listIntegers[i+1]+listIntegers[i+2]+listIntegers[i+3] {
			total += 1
		}
		current = listIntegers[i+1] + listIntegers[i+2] + listIntegers[i+3]
	}

	fmt.Printf("Part 2: %d\n", total) // part 2

}
