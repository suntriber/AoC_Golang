package main

import (
	"AoC_Golang/AoC_2024/helpers"
	"fmt"
)

type Point struct {
	X int
	Y int
}

type Region struct {
	Points []Point
}

func main() {
	fmt.Println("day 12..")
	data, err := helpers.ReadStrings("test.txt")
	if err != nil {
		panic(err)
	}

	for _, l := range data {
		fmt.Println(l)
	}

	smallExample := []string{
		"AAAA",
		"BBCD",
		"BBCC",
		"EEEC",
	}

	res := findRegions(smallExample)
	for _, n := range res {
		fmt.Println(n)
	}
}

func findRegions(s []string) []Region {
	regions := []Region{}

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			
		}
	}

	return regions
}
