package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("day 1..")
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	dataStr := string(data)
	dataList := strings.Split(dataStr, "\n")

	listA, listB := []int{}, []int{}

	for _, line := range dataList {
		t := strings.Split(line, "   ")
		a, err := strconv.Atoi(t[0])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(t[1])
		if err != nil {
			panic(err)
		}
		listA = append(listA, a)
		listB = append(listB, b)
	}

	slices.Sort(listA)
	slices.Sort(listB)

	sum := 0

	for i := 0; i < len(listA); i++ {
		sum += getDifference(listA[i], listB[i])
	}
	fmt.Printf("Part 1: %d\n", sum)

	similarityScore := 0

	for _, n := range listA {
		appears := 0
		for _, m := range listB {
			if n == m {
				appears += 1
			}
		}
		similarityScore += n * appears
		appears = 0
	}

	fmt.Printf("Part 2: %d\n", similarityScore)
}

func getDifference(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
