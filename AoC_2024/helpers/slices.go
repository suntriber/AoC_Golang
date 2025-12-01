package helpers

import (
	"strconv"
	"strings"
)

// StringsToInts takes a string and a separator. It splits the string at the separator and returns a slice with all elements converted to integers.
func StringsToInts(s string, separator string) ([]int, error) {
	tmp := []int{}
	parts := strings.Split(s, separator)
	for _, v := range parts {
		if v == "" { // edge case
			continue
		}
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		tmp = append(tmp, n)
	}
	return tmp, nil
}
