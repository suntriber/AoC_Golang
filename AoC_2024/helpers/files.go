package helpers

import (
	"os"
	"strconv"
	"strings"
)

// ReadString take a path to a file and returns the contents as a single string
func ReadString(f string) (string, error) {
	data, err := os.ReadFile(f)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ReadStrings take a path to a file and return the contents as an array of strings
func ReadStrings(f string) ([]string, error) {
	data, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}
	dataStr := string(data)

	return strings.Split(dataStr, "\n"), nil
}

// StringListToIntList take a list of strings and returns a list of ints
func StringListToIntList(s []string) ([]int, error) {
	intList := []int{}
	for _, v := range s {
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		intList = append(intList, n)
	}
	return intList, nil
}

// ReadInts takes a path to a file and returns the contents as a list of ints
func ReadInts(f string) ([]int, error) {
	listStr, err := ReadStrings(f)
	if err != nil {
		return nil, err
	}
	listInt, err := StringListToIntList(listStr)
	if err != nil {
		return nil, err
	}
	return listInt, nil
}
