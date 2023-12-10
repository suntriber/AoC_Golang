package myutils

// Sum takes a slice of ints and returns the sum of all its elements
func Sum(n []int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum
}
