package main

import (
	"fmt"
)

func generatePermutations(slice []string) [][]string {
	if len(slice) == 1 {
		return [][]string{slice}
	}

	var result [][]string
	for i, str := range slice {
		// Create a new slice excluding the current element
		rest := make([]string, 0)
		rest = append(rest, slice[:i]...)
		rest = append(rest, slice[i+1:]...)

		// Generate permutations of the rest
		subPermutations := generatePermutations(rest)
		for _, subPerm := range subPermutations {
			perm := append([]string{str}, subPerm...)
			result = append(result, perm)
		}
	}
	return result
}

func main() {
	slice := []string{"a", "b", "c"}
	permutations := generatePermutations(slice)
	for _, perm := range permutations {
		fmt.Println(perm)
	}
}
