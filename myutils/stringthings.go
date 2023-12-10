package myutils

import (
	"strconv"
	"unicode"
)

func ReverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// NumbersFromString iterates over a string and returns a slice with all numbers present
func NumbersFromString(s string) []int {
	nums := []int{}
	for i := 0; i < len(s); i++ {
		tmp := ""
		if unicode.IsDigit(rune(s[i])) {
			for {
				if i == len(s) || !unicode.IsDigit(rune(s[i])) {
					n, err := strconv.Atoi(tmp)
					if err != nil {
						panic(err)
					}
					nums = append(nums, n)
					break
				}
				tmp += string(s[i])
				i++
			}
		}
	}
	return nums
}
