package myutils_test

import (
	"AoC_Golang/myutils"
	"testing"
)

func TestReverseString(t *testing.T) {
	testCases := []struct {
		name         string
		strToReverse string
		want         string
	}{
		{name: "caps", strToReverse: "ABCDEFGH", want: "HGFEDCBA"},
		{name: "lower", strToReverse: "abcdefgh", want: "hgfedcba"},
		{name: "mixed", strToReverse: "ABcdEFgh", want: "hgFEdcBA"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := myutils.ReverseString(tc.strToReverse)
			if got != tc.want {
				t.Errorf("%s error, got %s want %s", tc.name, got, tc.want)
			}
		})
	}
}

func TestNumbersFromString(t *testing.T) {

	equal := func(a, b []int) bool {
		if len(a) != len(b) {
			return false
		}
		for i, v := range a {
			if v != b[i] {
				return false
			}
		}
		return true
	}

	testCases := []struct {
		name       string
		testString string
		want       []int
	}{
		{name: "no edge numbers", testString: "..12...34...45...56..", want: []int{12, 34, 45, 56}},
		{name: "left edge numbers", testString: "123...34...45...56..", want: []int{123, 34, 45, 56}},
		{name: "right edge numbers", testString: "...123...34...45...56..678", want: []int{123, 34, 45, 56, 678}},
		{name: "no numbers", testString: "...ABCDEFGH....abcdefgh....", want: []int{}},
		{name: "big number", testString: "123456789", want: []int{123456789}},
		{name: "single numbers", testString: "1.2.3.4.5.6.7.8.9.0", want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := myutils.NumbersFromString(tc.testString)
			if !equal(got, tc.want) {
				t.Errorf("%s error, got %v want %v", tc.name, got, tc.want)
			}
		})
	}
}
