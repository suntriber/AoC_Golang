package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	filePath := "input.txt"

	// read the whole content of file and pass it to file variable, in case of error pass it to err variable
	file, err := os.Open(filePath)
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

	current := listIntegers[0]
	total := 0

	for _, n := range listIntegers {
		if n > current {
			total++
		}
		current = n
	}

	fmt.Println(total) // part 1
}
