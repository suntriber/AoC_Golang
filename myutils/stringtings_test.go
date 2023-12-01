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
