package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var gameHistory map[int]int

	gameHistory, parseErr := parseInput("1,17,0,10,18,11,6")
	//gameHistory, parseErr := parseInput("3,1,2")

	if parseErr != nil {
		fmt.Println("parsing error", parseErr)
		return
	}

	//printParsedInput(gameHistory)

	start := len(gameHistory)

	for i := start; i < 2020; i++ {
		numberToConsider := gameHistory[i-1]

		alreadyUsed, numberOfTurnsAgo := checkGameHistory(i, numberToConsider, gameHistory)
		if alreadyUsed {
			gameHistory[i] = numberOfTurnsAgo
			fmt.Printf("%d, ", numberOfTurnsAgo)
		} else {
			gameHistory[i] = 0
			fmt.Printf("0, ")
		}

	}

	fmt.Printf("\nFound solution: %d \n", gameHistory[2019])
}

func checkGameHistory(currentRound int, numberToConsider int, gameHistory map[int]int) (bool, int) {
	for i := currentRound - 2; i >= 0; i-- {
		if gameHistory[i] == numberToConsider {
			return true, currentRound - i - 1
		}
	}
	return false, 0
}

func printParsedInput(numbers map[int]int) {
	for _, i := range numbers {
		fmt.Printf("%d, ", i)
	}
}

func parseInput(numbers string) (map[int]int, error) {
	numMap := make(map[int]int)
	strArray := strings.Split(numbers, ",")
	for numIndex, numStr := range strArray {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, err
		}
		numMap[numIndex] = num
		fmt.Printf("%d, ", num)
	}
	return numMap, nil
}
