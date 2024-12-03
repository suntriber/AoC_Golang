package main

import (
	"AoC_Golang/AoC_2024/helpers"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("day 3..")
	data, err := helpers.ReadString("input.txt")
	if err != nil {
		panic(err)
	}

	pattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	sum := 0
	submatches := pattern.FindAllStringSubmatch(data, -1)
	for _, match := range submatches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		sum += (num1 * num2)
	}
	fmt.Printf("Part 1: %d\n", sum)

	pattern2 := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d+),(\d+)\)`)
	matches2 := pattern2.FindAllStringSubmatch(data, -1)

	sum2 := 0
	processMul := true
	for _, match := range matches2 {
		if match[0] == "do()" {
			processMul = true
		} else if match[0] == "don't()" {
			processMul = false
		} else if processMul && match[1] != "" && match[2] != "" { // It's a mul(xxx,xxx)
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			sum2 += (num1 * num2)
		}
	}
	fmt.Printf("Part 2: %d\n", sum2)
}
