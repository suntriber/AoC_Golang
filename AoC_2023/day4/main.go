package main

import (
	"AoC_Golang/myutils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	input := myutils.ReadStrings("input.txt")
	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part2(input []string) int {

	type cardAndFactor struct {
		line  string
		cards int
		wins  int
	}

	scratchTickets := []cardAndFactor{}

	for _, l := range input {
		c := cardAndFactor{}
		c.line = l
		c.wins = getHowManyWinningNumsFromLine(l)
		c.cards = 1
		scratchTickets = append(scratchTickets, c)
	}

	for i, c := range scratchTickets {
		for j := 1; j <= c.wins; j++ {
			scratchTickets[i+j].cards += c.cards
		}

	}
	sum := 0
	for _, c := range scratchTickets {
		sum += c.cards
	}

	return sum
}

func part1(input []string) int {
	sum := 0
	for _, line := range input {
		val := 0
		myNums := getMyNumbersFromLine(line)
		winningNums := getWinningNumbersFromLine(line)
		for _, myN := range myNums {
			for _, winN := range winningNums {
				if myN == winN {
					val++
				}
			}
		}
		sum += int(math.Pow(2, float64(val-1)))
	}
	return sum
}

func getWinningNumbersFromLine(s string) []int {
	first := strings.Split(s, ":")
	second := strings.Split(first[1], "|")
	nums := strings.Split(second[0], " ")
	intNums := []int{}
	for _, n := range nums {
		if n == "" {
			continue
		}
		nn, err := strconv.Atoi(strings.TrimSpace(n))
		if err != nil {
			panic(err)
		}
		intNums = append(intNums, nn)
	}
	return intNums
}

func getMyNumbersFromLine(s string) []int {
	first := strings.Split(s, ":")
	second := strings.Split(first[1], "|")
	nums := strings.Split(second[1], " ")
	intNums := []int{}
	for _, n := range nums {
		if n == "" {
			continue
		}
		nn, err := strconv.Atoi(strings.TrimSpace(n))
		if err != nil {
			panic(err)
		}
		intNums = append(intNums, nn)
	}
	return intNums
}

func getHowManyWinningNumsFromLine(s string) int {
	out := []int{}
	bucket := map[int]bool{}
	myNums := getMyNumbersFromLine(s)
	winningNums := getWinningNumbersFromLine(s)
	for _, myN := range myNums {
		for _, winN := range winningNums {
			if myN == winN && !bucket[myN] {
				out = append(out, myN)
				bucket[myN] = true
			}
		}
	}
	return len(out)
}
