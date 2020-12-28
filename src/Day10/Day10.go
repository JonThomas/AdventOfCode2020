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

	// Algorithm:
	// Number on each side of threeInterval can't be swapped out.
	// All other numbers are marked as "Can be swapped out"
	// 1C = 2		<= 1 Consecutive C's means two possible combinations
	// 2C = 4		<= 2 Consecutive C's mean four possible combinations
	// 3C = 7		<= 3 Consecutive C's mean seven possible combinations
	// 4C = 11		<= 2 Consecutive C's mean eleven possible combinations
	// nC =
	// Total combinations is found by multiplying all possible combiations

	thisSequenceOfOnes := 0
	var sequenceList []int

	for i := 1; i < len(numbers); i++ {

		switch numbers[i] - numbers[i-1] {
		case 1:
			thisSequenceOfOnes++
		case 3:
			if thisSequenceOfOnes > 1 {
				sequenceList = append(sequenceList, thisSequenceOfOnes-1)
			}
			thisSequenceOfOnes = 0
		default:
			fmt.Printf("Index=%d, number %d=%d; %d=%d\n", i, i-1, numbers[i-1], i, numbers[i])
			return
		}
	}

	if thisSequenceOfOnes > 1 {
		sequenceList = append(sequenceList, thisSequenceOfOnes-1)
	}

	fmt.Println("Sequence: ", sequenceList)

	sum := 1
	for _, sequenceNr := range sequenceList {
		switch sequenceNr {
		case 1:
			sum *= 2
		case 2:
			sum *= 4
		case 3:
			sum *= 7
		}
	}

	fmt.Println("Sum: ", sum)
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
