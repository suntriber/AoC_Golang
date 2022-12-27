package main

import (
	utils "AoC_Golang/myutils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type coords struct {
	x int
	y int
}

func main() {
	fmt.Println("Welcome to day 9")
	data := utils.ReadStrings("input.txt")

	getDirAndNum := func(s string) (string, int) {
		l := strings.Split(s, " ")
		d := l[0]
		i, err := strconv.Atoi(l[1])
		if err != nil {
			panic(err)
		}
		return d, i
	}

	shouldTailMove := func(head, tail coords) bool {
		if math.Abs(float64(head.x)-float64(tail.x)) > 1 || math.Abs(float64(head.y)-float64(tail.y)) > 1 {
			return true
		}
		return false
	}

	makeMove := func(dir string, head, tail coords) (coords, coords) {
		switch dir {
		case "R":
			head.x++
			if shouldTailMove(head, tail) {
				tail.y = head.y
				tail.x = head.x - 1
			}
		case "L":
			head.x--
			if shouldTailMove(head, tail) {
				tail.y = head.y
				tail.x = head.x + 1
			}
		case "U":
			head.y--
			if shouldTailMove(head, tail) {
				tail.y = head.y + 1
				tail.x = head.x
			}
		case "D":
			head.y++
			if shouldTailMove(head, tail) {
				tail.y = head.y - 1
				tail.x = head.x
			}
		default:
			panic("AHH")
		}
		return head, tail
	}

	visitedHead := make(map[coords]int)
	visitedTail := make(map[coords]int)

	headCoords := &coords{
		x: 0,
		y: 0,
	}
	tailCoords := &coords{
		x: 0,
		y: 0,
	}

	headStart := &coords{
		x: 0,
		y: 0,
	}
	tailStart := &coords{
		x: 0,
		y: 0,
	}

	visitedHead[*headStart] = 1
	visitedTail[*tailStart] = 1

	for _, line := range data {
		d, i := getDirAndNum(line)
		for j := 0; j < i; j++ {
			*headCoords, *tailCoords = makeMove(d, *headCoords, *tailCoords)
			visitedHead[*headCoords]++
			visitedTail[*tailCoords]++
		}
	}
	fmt.Printf("part 1 : %d\n", len(visitedTail))
}
