package main

import (
	"AoC_Golang/myutils"
	"fmt"
	"strings"
)

func main() {
	input := myutils.ReadStrings("test.txt")
	// fmt.Println(input)
	fmt.Println(len(input))
	fmt.Println(hasFiveOfAKind("AAAAA"))
	fmt.Println(hasFiveOfAKind("AAAAB"))
	i, b := hasFourOfAKind("AAAAA")
	fmt.Println(i, b)
	ii, bb := hasFourOfAKind("AAAAB")
	fmt.Println(ii, bb)
	fmt.Println(hasFullHouse("ABCDE"))
	fmt.Println(hasFullHouse("AABBC"))
	fmt.Println(hasFullHouse("ABBBC"))
	fmt.Println(hasFullHouse("AAABB"))
	fmt.Println(hasFullHouse("ABABA"))
}

func part1(input []string) int {
	return 0
}

// winningHand compares two hands and returns the winner
func winningHand(a, b string) string {
	return ""
}

func hasFiveOfAKind(s string) bool {
	return strings.Count(s, string(s[0])) == len(s)
}

// hasFourOfAKind checks if a hand has four of a kind, returns the index of a value withing the foursome and a boolen
func hasFourOfAKind(s string) (int, bool) {
	for i, c := range s {
		if strings.Count(s, string(c)) == 4 {
			return i, true
		}
	}
	return -1, false
}

func hasThreeOfAKind(s string) (int, bool) {
	for i, c := range s {
		if strings.Count(s, string(c)) == 3 {
			return i, true
		}
	}
	return -1, false
}

func hasFullHouse(s string) bool {
	hasThree := false
	hasTwo := false
	for _, c := range s {
		if strings.Count(s, string(c)) == 3 {
			hasThree = true
		}
		if strings.Count(s, string(c)) == 2 {
			hasTwo = true
		}
	}
	return hasThree == true && hasTwo == true
}
