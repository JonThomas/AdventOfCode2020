package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	boardingCards, err := scanMap("./Day5input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	seatIds := []int{}

	for _, boardingCard := range boardingCards {
		row, err := getRow(boardingCard)
		if err != nil {
			fmt.Println(err)
			return
		}

		seat, seatErr := getSeat(boardingCard)
		if seatErr != nil {
			fmt.Println(err)
			return
		}

		thisSeatId := row*8 + seat
		seatIds = append(seatIds, int(thisSeatId))
	}

	sort.Ints(seatIds)

	for i := 0; i < len(seatIds)-3; i++ {
		//		fmt.Println(seatIds[i], " ", seatIds[i+1], " ", seatIds[i+2], " Plus1: ", seatIds[i]+1)

		if seatIds[i]+1 != seatIds[i+1] {
			fmt.Println(seatIds[i], " ", seatIds[i+1], " ", seatIds[i+2])
		}
	}
}

func getRow(boardingCard string) (int64, error) {
	rowWeird := boardingCard[0:7]
	temp := strings.Replace(rowWeird, "F", "0", -1)
	rowString := strings.Replace(temp, "B", "1", -1)
	return strconv.ParseInt(rowString, 2, 64)
}

func getSeat(boardingCard string) (int64, error) {
	seatWeird := boardingCard[7:10]
	temp := strings.Replace(seatWeird, "R", "1", -1)
	seatString := strings.Replace(temp, "L", "0", -1)
	return strconv.ParseInt(seatString, 2, 64)
}

func scanMap(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var boardingPasses []string

	for scanner.Scan() {
		var line = scanner.Text()
		boardingPasses = append(boardingPasses, line)
	}

	return boardingPasses, nil
}
