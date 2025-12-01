package main

import (
	"AoC_Golang/AoC_2024/helpers"
	"fmt"
)

func main() {
	fmt.Println("day 08..")
	data, err := helpers.ReadStrings2D("test.txt")
	if err != nil {
		panic(err)
	}
	for _, l := range data {
		fmt.Printf("%s\n", string(l))
	}
}
