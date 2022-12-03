package main

import (
	utils "AoC_Golang/myutils"
	"fmt"
)

func main() {
	fmt.Println("Welcome to day 3")
	data := utils.ReadStrings("input.txt")
	part1(data)
	part2(data)
}

// map that has map[char]value
func getValuesMap() map[rune]int {
	values := make(map[rune]int)

	k := 'a'
	v := 1

	for i := 0; i < 26; i++ {
		values[k] = v
		k++
		v++
	}

	k = 'A'
	v = 27

	for i := 0; i < 26; i++ {
		values[k] = v
		k++
		v++
	}
	return values
}

func part1(data []string) {
	values := getValuesMap()
	total := 0

	for _, line := range data {
		t1 := line[:len(line)/2]
		t2 := line[len(line)/2:]
		found := false
		for _, v := range t1 {
			for _, v2 := range t2 {
				if v == v2 {
					total += values[v]
					// fmt.Printf("%c\n", v)
					found = true
					break
				}
			}
			if found {
				break
			}
		}
	}

	fmt.Printf("Part 1 : %d\n", total)

}

func part2(data []string) {
	values := getValuesMap()
	total := 0

	for i := 0; i < len(data); i += 3 {
		t1 := data[i]
		t2 := data[i+1]
		t3 := data[i+2]

		found := false
		for _, c1 := range t1 {
			for _, c2 := range t2 {
				if c1 == c2 {
					for _, c3 := range t3 {
						if c2 == c3 {
							found = true
							total += values[c3]
							// fmt.Printf("%c\n", c3)
							break
						}
					}
					if found {
						break
					}
				}
			}
			if found {
				break
			}
		}
	}
	fmt.Printf("Part 2 : %d\n", total)
}
