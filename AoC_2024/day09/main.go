package main

import (
	"AoC_Golang/AoC_2024/helpers"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("day 9..")
	data, err := helpers.ReadString("test.txt")
	if err != nil {
		panic(err)
	}

	s := ""
	v := 0
	for i, r := range data {

		// read rune value as int
		intVal := int(r - '0')

		if i%2 == 0 {
			for j := 0; j < intVal; j++ {
				s += strconv.Itoa(v)
			}
			v++
		} else {
			for j := 0; j < intVal; j++ {
				s += "."
			}
		}

	}
	fmt.Println(s+"  len(s) : ", len(s))
	ss := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		ss[i] = s[i]
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if ss[j] == '.' {
				ss[j] = s[i]
				ss[i] = '.'
				break
			}
		}
		fmt.Println(string(ss))
	}

	var sum int64
	for i := 0; i < len(ss); i++ {
		if ss[i] == '.' {
			continue
		}
		intVal, err := strconv.Atoi(string(ss[i]))
		if err != nil {
			panic(err)
		}
		sum += int64((intVal * i))
	}
	fmt.Printf("Part 1: %d\n", sum)
}
