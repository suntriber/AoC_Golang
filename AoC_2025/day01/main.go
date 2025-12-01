package main

import (
	"AoC_Golang/AoC_2024/helpers"
	"fmt"
	"strconv"
)

func main() {
	data, err := helpers.ReadStrings("input.txt")
	if err != nil {
		panic(err)
	}

	dial := 50 // starting position
	zeroCounter := 0
	clicks := 0
	wasAtZero := false

	for _, line := range data {
		direction := line[0]
		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		if steps > 100 {
			clicks += steps / 100
			steps = steps % 100
		}

		switch direction {
		case 'L':
			dial -= steps
			if dial < 0 {
				if !wasAtZero {
					clicks++
				}
				dial += 100
			}
			dial = dial % 100
		case 'R':
			dial += steps
			if dial > 100 {
				if !wasAtZero {
					clicks++
				}
				dial -= 100
			}
			dial = dial % 100
		default:
			panic("invalid direction")
		}

		if dial == 0 {
			zeroCounter++
			wasAtZero = true
		} else {
			wasAtZero = false
		}
	}

	fmt.Printf("Part 1: %d\n", zeroCounter)        // part 1 answer
	fmt.Printf("Part 2: %d\n", clicks+zeroCounter) // part 2 answer

}
