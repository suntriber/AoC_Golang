package main

import (
	"AoC_Golang/AoC_2024/helpers"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("day 7..")
	data, err := helpers.ReadStrings("test.txt")
	if err != nil {
		panic(err)
	}

	sum1 := 0

	for _, l := range data {
		parts := strings.Split(l, ":")
		targetValue, _ := strconv.Atoi(parts[0])
		nums, err := helpers.StringsToInts(parts[1], " ")
		if err != nil {
			panic(err)
		}
		tmp := make([]string, 0)
		generateExpressions(nums, 0, "+", &tmp)
		fmt.Println(targetValue, " : ", tmp)
		for _, ex := range tmp {
			val, err := calculateExpression(ex)
			if err != nil {
				panic(err)
			}
			fmt.Println("val: ", val)
			if val == targetValue {
				sum1 += val
				fmt.Printf("Expression: %s\n", ex)
			}
		}
	}
	fmt.Printf("Part1: %d\n", sum1)
}

func generateExpressions(nums []int, index int, currentExpr string, results *[]string) {
	// Base case: if we've processed all numbers
	if index == len(nums) {
		*results = append(*results, currentExpr)
		return
	}

	// If this is the first number, just add it to the expression
	if index == 0 {
		generateExpressions(nums, index+1, strconv.Itoa(nums[index]), results)
	} else {
		// Add '+' and recursively call for the next number
		generateExpressions(nums, index+1, currentExpr+"+"+strconv.Itoa(nums[index]), results)
		// Add '*' and recursively call for the next number
		generateExpressions(nums, index+1, currentExpr+"*"+strconv.Itoa(nums[index]), results)
	}
}

func calculateExpression(expr string) (int, error) {
	sum := 0

	op := ""
	for _, c := range expr {
		if string(c) == "+" {
			op = "+"
			break
		}
		if string(c) == "*" {
			op = "*"
		}
	}

	s := ""
	s += string(expr[0])
	for i := 1; i < len(expr); i++ {
		if string(expr[i]) == "+" {
			n, err := strconv.Atoi(s)
			if err != nil {
				return 0, err
			}
			if op == "*" {
				sum *= n
			} else if op == "+" {
				sum += n
			}
			op = "+"
			s = ""
		} else if string(expr[i]) == "*" {
			n, err := strconv.Atoi(s)
			if err != nil {
				return 0, err
			}
			if op == "*" {
				sum *= n
			} else if op == "+" {
				sum += n
			}
			op = "*"
			s = ""
		} else {
			s += string(expr[i])
		}
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	if op == "*" {
		sum *= n
	} else if op == "+" {
		sum += n
	}

	return sum, nil
}
