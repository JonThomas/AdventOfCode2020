package main

import (
	"fmt"
	"jonthomas/AdventOfCode2020/files"
	"strconv"
)

type instruction struct {
	Operation string
	Count     int
	Visited   bool
}

func main() {

	file, err := files.ReadFile("./Day9Input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	numbers, parseErr := parseInput(file)
	if parseErr != nil {
		fmt.Println("parsing error", err)
		return
	}

	numbersToConsider := 25

	for i := numbersToConsider; i < len(numbers); i++ {
		fmt.Printf("Looking for %d: ", numbers[i])

		var numbersToCheck []int
		for j := i - numbersToConsider; j < i; j++ {
			fmt.Printf("%d - ", numbers[j])
			numbersToCheck = append(numbersToCheck, numbers[j])
		}
		fmt.Println()
		if !twoNumbersAddUp(numbers[i], numbersToCheck) {
			fmt.Printf("No two numbers add up to &d (%v)\n", numbers[i], numbersToCheck)
			break
		}
	}
	fmt.Printf("END\n")
}

func twoNumbersAddUp(sum int, toCheck []int) bool {
	for i := 0; i < len(toCheck)-1; i++ {
		for j := i + 1; j < len(toCheck); j++ {
			if toCheck[i]+toCheck[j] == sum {
				return true
			}
		}
	}
	return false
}

func parseInput(fileLines []string) ([]int, error) {

	var intLines []int

	for _, line := range fileLines {
		nr, nrErr := strconv.Atoi(line)
		if nrErr != nil {
			return nil, nrErr
		}
		intLines = append(intLines, nr)
	}
	return intLines, nil
}
