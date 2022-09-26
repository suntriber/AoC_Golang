package myutils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadIntegers(path string) []int {
	// read the whole content of file and pass it to file variable, in case of error pass it to err variable
	file, err := os.Open(path)
	Check(err)
	if err != nil {
		fmt.Printf("Could not open the file due to this %s error \n", err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var listIntegers []int

	for fileScanner.Scan() {
		i, _ := strconv.Atoi(fileScanner.Text())
		listIntegers = append(listIntegers, i)
	}
	if err = file.Close(); err != nil {
		fmt.Printf("Could not close the file due to this %s error \n", err)
	}

	return listIntegers
}
