package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("day 2..")

	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	dataStrList := strings.Split(string(data), "\n")
	part1 := 0
	for _, l := range dataStrList {
		list := strings.Split(l, " ")
		if isSafe(list) {
			part1 += 1
		}
	}

	fmt.Printf("Part 1: %d\n", part1)

	part2 := 0
	for _, l := range dataStrList {
		list := strings.Split(l, " ")
		if isSafe(list) {
			part2 += 1
		} else {
			for i := 0; i < len(list); i++ {
				tmpList := make([]string, 0, len(list)-1)
				for j := 0; j < len(list); j++ {
					if j != i {
						tmpList = append(tmpList, list[j])
					}
				}

				if isSafe(tmpList) {
					part2 += 1
					break
				}
			}
		}
	}

	fmt.Printf("Part 2: %d\n", part2)

}

func isSafe(l []string) bool {
	intList := []int{}

	for _, v := range l {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		intList = append(intList, i)
	}

	// check increasing
	v := intList[0]
	for i := 1; i < len(intList); i++ {
		if v < intList[i] {
			if (intList[i] - v) > 3 {
				return false
			}
			v = intList[i]
		} else if v == intList[i] {
			return false
		}
	}

	// check decreasing
	vv := intList[0]
	for i := 1; i < len(intList); i++ {
		if vv > intList[i] {
			if (vv - intList[i]) > 3 {
				return false
			}
			vv = intList[i]
		} else if vv == intList[i] {
			return false
		}
	}

	// check if change flips
	vvv := intList[0]
	increasing := false
	decreasing := false
	for i := 1; i < len(intList); i++ {
		if vvv == intList[i] {
			return false
		} else if vvv < intList[i] {
			decreasing = true
			if increasing {
				return false
			}
		} else if vvv > intList[i] {
			increasing = true
			if decreasing {
				return false
			}
		}
		vvv = intList[i]
	}
	return true
}
