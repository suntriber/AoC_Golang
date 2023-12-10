package main

import (
	"AoC_Golang/myutils"
	"fmt"
	"strconv"
)

func main() {
	input := myutils.ReadStrings("input.txt")
	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	time := myutils.NumbersFromString(input[0])
	distance := myutils.NumbersFromString(input[1])

	fmt.Printf("%T %v\n", time, time)
	fmt.Printf("%T %v\n", distance, distance)

	sum := 1
	for i := 0; i < len(time); i++ {
		wins := winRace(time[i], distance[i])
		sum *= len(wins)
	}

	return sum
}

func winRace(t, d int) []int {
	w := []int{}
	speed := 0
	for i := 1; i <= t; i++ {
		speed++
		if speed*(t-i) > d {
			w = append(w, i)
		}
	}
	return w
}

func part2(input []string) int {
	time := myutils.NumbersFromString(input[0])
	distance := myutils.NumbersFromString(input[1])

	t := ""
	d := ""

	for i := 0; i < len(time); i++ {
		t += fmt.Sprintf("%d", time[i])
		d += fmt.Sprintf("%d", distance[i])
	}

	tInt, err := strconv.Atoi(t)
	if err != nil {
		panic(err)
	}
	dInt, err := strconv.Atoi(d)
	if err != nil {
		panic(err)
	}
	return len(winRace(tInt, dInt))
}
