package main

import (
	"AoC_Golang/AoC_2024/helpers"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("day 5..")

	data, err := helpers.ReadStrings("input.txt")
	if err != nil {
		panic(err)
	}

	var flipIndex int
	for i := 0; i < len(data); i++ {
		if data[i] == "" {
			flipIndex = i
			break
		}
	}
	pageOrderingRules, updates := data[:flipIndex], data[flipIndex+1:]

	validUpdatesIndex := []int{}
	invalidUpdatesIndex := []int{}

	for i := 0; i < len(updates); i++ {
		pageOrder := strings.Split(updates[i], ",")
		if !doesPageBreakRules(pageOrderingRules, pageOrder) {
			validUpdatesIndex = append(validUpdatesIndex, i)
		} else {
			invalidUpdatesIndex = append(invalidUpdatesIndex, i) // part 2
		}
	}

	sum := 0
	for _, n := range validUpdatesIndex {
		p := updates[n]
		parts := strings.Split(p, ",")
		intVal, err := strconv.Atoi(parts[(len(parts) / 2)])
		if err != nil {
			panic(err)
		}
		sum += intVal
	}
	fmt.Printf("Part 1: %d\n", sum)

	sum2 := 0
	correctedPages := [][]string{}
	for _, n := range invalidUpdatesIndex {
		p := updates[n]
		parts := strings.Split(p, ",")
		idx1, idx2, ok := doesPageBreakRulesPart2(pageOrderingRules, parts)
		for ok {
			parts[idx1], parts[idx2] = parts[idx2], parts[idx1]
			idx1, idx2, ok = doesPageBreakRulesPart2(pageOrderingRules, parts)
		}
		correctedPages = append(correctedPages, parts)

	}
	for _, p := range correctedPages {
		intVal, err := strconv.Atoi(p[(len(p) / 2)])
		if err != nil {
			panic(err)
		}
		sum2 += intVal
	}
	fmt.Printf("Part 2: %d\n", sum2)
}

func doesPageBreakRulesPart2(rules []string, pages []string) (int, int, bool) {
	for _, r := range rules {
		parts := strings.Split(r, "|")
		firstRule, secondRule := parts[0], parts[1]

		idx1, idx2, err := isBothRulesInPages(firstRule, secondRule, pages)
		if err != nil {
			continue
		}
		if idx2 < idx1 {
			return idx1, idx2, true
		}
	}
	return -1, -1, false
}

func doesPageBreakRules(rules []string, pages []string) bool {
	for _, r := range rules {
		parts := strings.Split(r, "|")
		firstRule, secondRule := parts[0], parts[1]

		idx1, idx2, err := isBothRulesInPages(firstRule, secondRule, pages)
		if err != nil {
			continue
		}
		if idx2 < idx1 {
			return true
		}
	}
	return false
}

func isBothRulesInPages(r1, r2 string, pgs []string) (int, int, error) {
	var idx1 int
	var idx2 int
	idx1Found := false
	idx2Found := false

	for i, p := range pgs {
		if p == r1 {
			idx1 = i
			idx1Found = true
		}
		if p == r2 {
			idx2 = i
			idx2Found = true
		}
	}
	if idx1Found && idx2Found {
		return idx1, idx2, nil
	}
	return 0, 0, fmt.Errorf("both rules not in page")
}
