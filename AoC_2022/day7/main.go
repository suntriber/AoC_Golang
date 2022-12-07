package main

import (
	utils "AoC_Golang/myutils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to day 7")
	data := utils.ReadStrings("input.txt")
	dirs := func() map[string]int {
		d := make(map[string]int)
		loc := []string{"/"}
		for _, line := range data {
			commands := strings.Split(line, " ")
			if commands[0] == "$" {
				if commands[1] == "cd" {
					if commands[2] == ".." {
						loc = loc[:len(loc)-1]
					} else if commands[2] == "/" {
						loc = []string{"/"}
					} else {
						loc = append(loc, commands[2])
					}
				}
				continue
			}
			if commands[0] != "dir" {
				file_size, _ := strconv.Atoi(commands[0])
				for i := 0; i <= len(loc); i++ {
					d[strings.Join(loc[:len(loc)-i], "-")] += file_size
				}
			}
		}
		return d
	}()

	tot := func() int {
		t := 0
		for _, dir_size := range dirs {
			if dir_size <= 100000 {
				t += dir_size
			}
		}
		return t
	}()

	getMax := func(d map[string]int) []int {
		have := 70000000 - d["/"]
		want := 30000000 - have
		canRemove := []int{}
		for _, s := range d {
			if s > want {
				canRemove = append(canRemove, s)
			}
		}
		sort.Ints(canRemove)
		return canRemove
	}

	removeableDirs := getMax(dirs)
	max := removeableDirs[0]

	fmt.Printf("Part 1 : %d\n", tot)
	fmt.Printf("Part 2 : %d\n", max)
}
