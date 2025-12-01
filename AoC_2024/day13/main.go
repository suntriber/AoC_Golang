package main

import (
	"AoC_Golang/AoC_2024/helpers"
	"fmt"
	"regexp"
	"slices"
	"strconv"
)

/*
Button A = 3 tokens
Button B = 1 token

*/

func main() {
	fmt.Println("day 13..")

	data, err := helpers.ReadStrings("input.txt")
	if err != nil {
		panic(err)
	}

	total := 0

	for i := 0; i < len(data); i += 4 {
		xA, yA := getButtonXY(data[i])
		xB, yB := getButtonXY(data[i+1])
		prizeX, prizeY := getButtonXY(data[i+2])

		// Solve for X equation: 94a + 22b = targetX
		gcdX, xCoeffA, xCoeffB := extendedGCD(xA, xB)
		if prizeX%gcdX != 0 {
			// fmt.Printf("Prize [X=%d, Y=%d] No solution exists for X\n", prizeX, prizeY)
			continue
		}

		// Solve for Y equation: 34a + 67b = targetY
		gcdY, yCoeffA, yCoeffB := extendedGCD(yA, yB)
		if prizeY%gcdY != 0 {
			// fmt.Printf("Prize [X=%d, Y=%d] No solution exists for Y\n", prizeX, prizeY)
			continue
		}

		// Scale coefficients to match target values
		scaleX := prizeX / gcdX
		xCoeffA *= scaleX
		xCoeffB *= scaleX

		scaleY := prizeY / gcdY
		yCoeffA *= scaleY
		yCoeffB *= scaleY

		spentTokens := []int{}
		// Combine solutions to minimize button presses (a and b must be integers)
		for k := -1000; k <= 1000; k++ { // Adjust range as needed for practical solutions
			a := xCoeffA + k*(xB/gcdX)
			b := xCoeffB - k*(xA/gcdX)

			if a >= 0 && b >= 0 { // Ensure non-negative solutions
				if a+b <= 100 {
					spentTokens = append(spentTokens, a*3+b)

				}
			}
		}
		slices.Sort(spentTokens)
		if len(spentTokens) != 0 {
			total += spentTokens[0]

		}
		// fmt.Println(spentTokens)
	}
	fmt.Println(total)
}

func getButtonXY(line string) (int, int) {
	// Create a regular expression to match one or more digits
	re := regexp.MustCompile(`\d+`)

	// Find all matches in the text
	matches := re.FindAllString(line, -1)

	x, y := matches[0], matches[1]

	xInt, err := strconv.Atoi(x)
	if err != nil {
		panic(err)
	}
	yInt, err := strconv.Atoi(y)
	if err != nil {
		panic(err)
	}
	return xInt, yInt
}

// Function to calculate GCD and coefficients using Extended Euclidean Algorithm
func extendedGCD(a, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, x1, y1 := extendedGCD(b, a%b)
	x := y1
	y := x1 - (a/b)*y1
	return gcd, x, y
}
