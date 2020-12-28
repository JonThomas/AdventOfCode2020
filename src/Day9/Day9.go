package main

import (
	"fmt"
	"jonthomas/AdventOfCode2020/files"
	"strconv"
)

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

	sumFrom9a := 27911108
	//sumFrom9a := 127

	for i := 0; i < len(numbers); i++ {

		if findNumberSequence(i, numbers, sumFrom9a) {
			break
		}
	}
	fmt.Println("END")
}

func findNumberSequence(i int, numbers []int, goal int) bool {
	var numbersToCheck []int
	for j := i; j < len(numbers); j++ {
		fmt.Printf("%d - ", numbers[j])
		numbersToCheck = append(numbersToCheck, numbers[j])

		sum := addNumbers(numbersToCheck)
		if sum == goal {
			smallest := findSmallest(numbersToCheck)
			largest := findLargest(numbersToCheck)
			fmt.Printf("Found correct sum of %d. Smallest=%d, Largest=%d\nSum=%d\n", goal, smallest, largest, smallest+largest)
			return true
		}
		if sum > goal {
			fmt.Println("Too large!")
			return false
		}
	}
	fmt.Println("Exhausted sequence")
	return false
}

func addNumbers(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func findSmallest(numbers []int) int {
	smallest := numbers[0]
	for _, num := range numbers {
		if num < smallest {
			smallest = num
		}
	}
	return smallest
}

func findLargest(numbers []int) int {
	largest := numbers[0]
	for _, num := range numbers {
		if num > largest {
			largest = num
		}
	}
	return largest
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
