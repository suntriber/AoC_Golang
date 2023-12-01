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
	data := utils.ReadStrings("test2.txt")

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

	makeMove := func(dir string, head, tail, tail2, tail3, tail4, tail5, tail6, tail7, tail8, tail9 coords) (coords, coords, coords, coords, coords, coords, coords, coords, coords, coords) {
		switch dir {
		case "R":
			head.x++
			if shouldTailMove(head, tail) {
				// tail.y = head.y
				tail.x = head.x - 1
			}
			if shouldTailMove(tail, tail2) {
				// tail2.y = tail.y
				tail2.x = tail.x - 1
			}
			if shouldTailMove(tail2, tail3) {
				// tail3.y = tail2.y
				tail3.x = tail2.x - 1
			}
			if shouldTailMove(tail3, tail4) {
				// tail4.y = tail3.y
				tail4.x = tail3.x - 1
			}
			if shouldTailMove(tail4, tail5) {
				// tail5.y = tail4.y
				tail5.x = tail4.x - 1
			}
			if shouldTailMove(tail5, tail6) {
				// tail6.y = tail5.y
				tail6.x = tail5.x - 1
			}
			if shouldTailMove(tail6, tail7) {
				// tail7.y = tail6.y
				tail7.x = tail6.x - 1
			}
			if shouldTailMove(tail7, tail8) {
				// tail8.y = tail7.y
				tail8.x = tail7.x - 1
			}
			if shouldTailMove(tail8, tail9) {
				// tail9.y = tail8.y
				tail9.x = tail8.x - 1
			}
		case "L":
			head.x--
			if shouldTailMove(head, tail) {
				// tail.y = head.y
				tail.x = head.x + 1
			}
			if shouldTailMove(tail, tail2) {
				// tail2.y = tail.y
				tail2.x = tail.x + 1
			}
			if shouldTailMove(tail2, tail3) {
				// tail3.y = tail2.y
				tail3.x = tail2.x + 1
			}
			if shouldTailMove(tail3, tail4) {
				// tail4.y = tail3.y
				tail4.x = tail3.x + 1
			}
			if shouldTailMove(tail4, tail5) {
				// tail5.y = tail4.y
				tail5.x = tail4.x + 1
			}
			if shouldTailMove(tail5, tail6) {
				// tail6.y = tail5.y
				tail6.x = tail5.x + 1
			}
			if shouldTailMove(tail6, tail7) {
				// tail7.y = tail6.y
				tail7.x = tail6.x + 1
			}
			if shouldTailMove(tail7, tail8) {
				// tail8.y = tail7.y
				tail8.x = tail7.x + 1
			}
			if shouldTailMove(tail8, tail9) {
				// tail9.y = tail8.y
				tail9.x = tail8.x + 1
			}
		case "U":
			head.y--
			if shouldTailMove(head, tail) {
				tail.y = head.y + 1
				// tail.x = head.x
			}
			if shouldTailMove(tail, tail2) {
				tail2.y = tail.y + 1
				// tail2.x = tail.x
			}
			if shouldTailMove(tail2, tail3) {
				tail3.y = tail2.y + 1
				// tail3.x = tail2.x
			}
			if shouldTailMove(tail3, tail4) {
				tail4.y = tail3.y + 1
				// tail4.x = tail3.x
			}
			if shouldTailMove(tail4, tail5) {
				tail5.y = tail4.y + 1
				// tail5.x = tail4.x
			}
			if shouldTailMove(tail5, tail6) {
				tail6.y = tail5.y + 1
				// tail6.x = tail5.x
			}
			if shouldTailMove(tail6, tail7) {
				tail7.y = tail6.y + 1
				// tail7.x = tail6.x
			}
			if shouldTailMove(tail7, tail8) {
				tail8.y = tail7.y + 1
				// tail8.x = tail7.x
			}
			if shouldTailMove(tail8, tail9) {
				tail9.y = tail8.y + 1
				// tail9.x = tail8.x
			}
		case "D":
			head.y++
			if shouldTailMove(head, tail) {
				tail.y = head.y - 1
				// tail.x = head.x
			}
			if shouldTailMove(tail, tail2) {
				tail2.y = tail.y - 1
				// tail2.x = tail.x
			}
			if shouldTailMove(tail2, tail3) {
				tail3.y = tail2.y - 1
				// tail3.x = tail2.x
			}
			if shouldTailMove(tail3, tail4) {
				tail4.y = tail3.y - 1
				// tail4.x = tail3.x
			}
			if shouldTailMove(tail4, tail5) {
				tail5.y = tail4.y - 1
				// tail5.x = tail4.x
			}
			if shouldTailMove(tail5, tail6) {
				tail6.y = tail5.y - 1
				// tail6.x = tail5.x
			}
			if shouldTailMove(tail6, tail7) {
				tail7.y = tail6.y - 1
				// tail7.x = tail6.x
			}
			if shouldTailMove(tail7, tail8) {
				tail8.y = tail7.y - 1
				// tail8.x = tail7.x
			}
			if shouldTailMove(tail8, tail9) {
				tail9.y = tail8.y - 1
				// tail9.x = tail8.x
			}
		default:
			panic("AHH")
		}
		return head, tail, tail2, tail3, tail4, tail5, tail6, tail7, tail8, tail9
	}

	visitedHead := make(map[coords]int)
	visitedTail := make(map[coords]int)
	visitedTail2 := make(map[coords]int)
	visitedTail3 := make(map[coords]int)
	visitedTail4 := make(map[coords]int)
	visitedTail5 := make(map[coords]int)
	visitedTail6 := make(map[coords]int)
	visitedTail7 := make(map[coords]int)
	visitedTail8 := make(map[coords]int)
	visitedTail9 := make(map[coords]int)

	headCoords := &coords{
		x: 0,
		y: 0,
	}
	tailCoords := &coords{
		x: 0,
		y: 0,
	}
	tailCoords2 := &coords{
		x: 0,
		y: 0,
	}
	tailCoords3 := &coords{
		x: 0,
		y: 0,
	}
	tailCoords4 := &coords{
		x: 0,
		y: 0,
	}
	tailCoords5 := &coords{
		x: 0,
		y: 0,
	}
	tailCoords6 := &coords{
		x: 0,
		y: 0,
	}
	tailCoords7 := &coords{
		x: 0,
		y: 0,
	}
	tailCoords8 := &coords{
		x: 0,
		y: 0,
	}
	tailCoords9 := &coords{
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
	visitedTail2[*tailStart] = 1
	visitedTail3[*tailStart] = 1
	visitedTail4[*tailStart] = 1
	visitedTail5[*tailStart] = 1
	visitedTail6[*tailStart] = 1
	visitedTail7[*tailStart] = 1
	visitedTail8[*tailStart] = 1
	visitedTail9[*tailStart] = 1

	for _, line := range data {

		fmt.Println("line : ", line)

		d, i := getDirAndNum(line)
		for j := 0; j < i; j++ {
			*headCoords, *tailCoords, *tailCoords2, *tailCoords3, *tailCoords4, *tailCoords5, *tailCoords6, *tailCoords7, *tailCoords8, *tailCoords9 = makeMove(d, *headCoords, *tailCoords, *tailCoords2, *tailCoords3, *tailCoords4, *tailCoords5, *tailCoords6, *tailCoords7, *tailCoords8, *tailCoords9)
			visitedHead[*headCoords]++
			visitedTail[*tailCoords]++
			visitedTail2[*tailCoords2]++
			visitedTail3[*tailCoords3]++
			visitedTail4[*tailCoords4]++
			visitedTail5[*tailCoords5]++
			visitedTail6[*tailCoords6]++
			visitedTail7[*tailCoords7]++
			visitedTail8[*tailCoords8]++
			visitedTail9[*tailCoords9]++
			fmt.Printf("[%d %d], [%d %d], [%d %d], [%d %d], [%d %d], [%d %d], [%d %d], [%d %d], [%d %d], [%d %d]\n", headCoords.x, headCoords.y, tailCoords.x, tailCoords.y, tailCoords2.x, tailCoords2.y, tailCoords3.x, tailCoords3.y, tailCoords4.x, tailCoords4.y, tailCoords5.x, tailCoords5.y, tailCoords6.x, tailCoords6.y, tailCoords7.x, tailCoords7.y, tailCoords8.x, tailCoords8.y, tailCoords9.x, tailCoords9.y)
		}

	}
	fmt.Printf("part 1 : %d\n", len(visitedTail))
	fmt.Printf("part 2 : %d\n", len(visitedTail9))

	// headCoords2 := &coords{
	// 	x: 0,
	// 	y: 0,
	// }
	// tailCoords2 := &coords{
	// 	x: 0,
	// 	y: 0,
	// }
	// visitedHead2 := make(map[coords]int)
	// visitedTail2 := make(map[coords]int)

	// visitedHead2[*headStart] = 1
	// visitedTail2[*tailStart] = 1

	// data2 := utils.ReadStrings("test2.txt")
	// for _, line := range data2 {
	// 	d, i := getDirAndNum(line)
	// 	for j := 0; j < i; j++ {
	// 		*headCoords2, *tailCoords2 = makeMove(d, *headCoords2, *tailCoords2)
	// 	}

	// }
}
