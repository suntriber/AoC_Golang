package main

import (
	utils "AoC_Golang/myutils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to day 5")
	data := utils.ReadStrings("input.txt")[10:]

	stacks := make(map[int]string)
	stacks[1] = "HTZD"
	stacks[2] = "QRWTGCS"
	stacks[3] = "PBFQNRCH"
	stacks[4] = "LCNFHC"
	stacks[5] = "GLFQS"
	stacks[6] = "VPWZBRCS"
	stacks[7] = "ZFJ"
	stacks[8] = "DLVZRHQ"
	stacks[9] = "BHGNFZLD"

	part1(stacks, data)
	part2(stacks, data)
}

func getCommand(line string) (int, int, int) {
	splitLine := strings.Split(line, " ")
	last, _ := strconv.Atoi(string(splitLine[5]))
	first, _ := strconv.Atoi(string(splitLine[1]))
	middle, _ := strconv.Atoi(string(splitLine[3]))
	return first, middle, last
}

func part1(s map[int]string, data []string) {
	stacks := make(map[int]string)
	for k, v := range s {
		stacks[k] = v
	}
	for _, line := range data {
		numberBoxes, fromStack, toStack := getCommand(line)
		for i := 0; i < numberBoxes; i++ {
			stacks[toStack] += string(stacks[fromStack][len(stacks[fromStack])-1])
			stacks[fromStack] = stacks[fromStack][:len(stacks[fromStack])-1]
		}
	}
	fmt.Print("Part 1 : ")
	for i := 1; i <= 9; i++ {
		fmt.Printf("%c", stacks[i][len(stacks[i])-1])
	}
	fmt.Println()
}

func part2(s map[int]string, data []string) {
	stacks := make(map[int]string)
	for k, v := range s {
		stacks[k] = v
	}
	for _, line := range data {
		numberBoxes, fromStack, toStack := getCommand(line)
		l := len(stacks[fromStack])
		stacks[toStack] += stacks[fromStack][l-numberBoxes : l]
		stacks[fromStack] = stacks[fromStack][:l-numberBoxes]
	}
	fmt.Print("Part 2 : ")
	for i := 1; i <= 9; i++ {
		fmt.Printf("%c", stacks[i][len(stacks[i])-1])
	}
	fmt.Println()
}
