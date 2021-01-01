package main

import (
	"fmt"
	"strconv"
	"strings"
)

type numberUsage struct {
	lastUsedRound      int
	secondLastUsedDiff int
	firstTime          bool
}

// key: The spoken number
// value: Information on when the number was used
type lastUsedIndex map[int]numberUsage

func main() {

	endAt := 30000000

	lastUsed, roundNr, lastNr, parseErr := parseInput("1,17,0,10,18,11,6")
	//lastUsed, roundNr, lastNr, parseErr := parseInput("1,2,3")

	if parseErr != nil {
		fmt.Println("parsing error", parseErr)
		return
	}

	var numInfo numberUsage
	exists := false

	for i := roundNr; i <= endAt; i++ {
		usageInfo, ok := lastUsed[lastNr]
		if ok && usageInfo.firstTime {
			// Insert 0, but first check if 0 has been used before
			numInfo, exists = lastUsed[0]
			lastUsed[0] = numberUsage{lastUsedRound: i, firstTime: !exists, secondLastUsedDiff: i - numInfo.lastUsedRound}
			//fmt.Printf("%d: Considering %d. Inserting 0\n", i, lastNr)
			lastNr = 0
		} else {
			numberToInsert := usageInfo.secondLastUsedDiff
			numInfo, exists = lastUsed[numberToInsert]
			lastUsed[numberToInsert] = numberUsage{lastUsedRound: i, firstTime: !exists, secondLastUsedDiff: i - numInfo.lastUsedRound}
			//fmt.Printf("%d: Considering %d. Inserting %d (%d was last used in round %d).\n", i, lastNr, numberToInsert, lastNr, usageInfo.lastUsedRound)
			lastNr = numberToInsert
		}
		if i%10000 == 0 {
			fmt.Println("At round: ", i)
		}
	}

	fmt.Printf("\nFound solution: %d \n", lastNr)
}

func parseInput(numbers string) (lastUsedIndex, int, int, error) {
	lastUsed := make(lastUsedIndex)
	strArray := strings.Split(numbers, ",")
	var roundNr int
	lastNr := 0
	for roundNr = 1; roundNr <= len(strArray); roundNr++ {
		num, err := strconv.Atoi(strArray[roundNr-1])
		if err != nil {
			return nil, 0, 0, err
		}

		lastUsed[num] = numberUsage{lastUsedRound: roundNr, firstTime: true, secondLastUsedDiff: 0}
		lastNr = num
		fmt.Printf("%d: Considering %d\n", roundNr, num)
	}
	return lastUsed, roundNr, lastNr, nil
}
