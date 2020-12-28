package main

import (
	"fmt"
	"jonthomas/AdventOfCode2020/files"
	"sort"
	"strconv"
)

func main() {

	file, err := files.ReadFile("./Day10Input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	numbers, parseErr := parseInput(file)
	if parseErr != nil {
		fmt.Println("parsing error", err)
		return
	}

	sort.Ints(numbers)

	oneIntervals := 0
	threeIntervals := 1 // Counting the Built-in adapter

	for i := 1; i < len(numbers); i++ {

		switch numbers[i] - numbers[i-1] {
		case 1:
			oneIntervals++
		case 3:
			threeIntervals++
		default:
			fmt.Printf("Index=%d, number %d=%d; %d=%d\n", i, i-1, numbers[i-1], i, numbers[i])
			return
		}
	}

	fmt.Printf("%d*%d=%d\n", oneIntervals, threeIntervals, oneIntervals*threeIntervals)
	fmt.Println("END")
}

func parseInput(fileLines []string) ([]int, error) {

	intLines := []int{0}

	for _, line := range fileLines {
		nr, nrErr := strconv.Atoi(line)
		if nrErr != nil {
			return nil, nrErr
		}
		intLines = append(intLines, nr)
	}
	return intLines, nil
}
