package main

import (
	"fmt"
	"jonthomas/AdventOfCode2020/files"
)

var numSeats int
var numRows int

func main() {

	// L = free
	// # = occupied
	// . = floor
	var seatRows [][]byte

	file, err := files.ReadFile("./Day11Input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	seatRows, parseErr := parseInput(file)
	if parseErr != nil {
		fmt.Println("parsing error", err)
		return
	}

	numRows = len(seatRows)
	numSeats = len(seatRows[0])

	printSeatLayout("Round 1", seatRows)

	seatChange := false

	numRounds := 2

	// Loop indefinately, until passengers have stabilized their seating
	for {
		seatChange, seatRows = performSeatChange(seatRows)
		if !seatChange {
			occupiedSeats := countOccupiedSeats(seatRows)
			fmt.Println("Stable seating situation. Occupied seats: ", occupiedSeats)
			break
		}
		printSeatLayout(fmt.Sprintf("Round %d", numRounds), seatRows)
		numRounds++
	}

	fmt.Println("END")
}

func countOccupiedSeats(seatRows [][]byte) int {
	numOccupied := 0
	for row := 0; row < numRows; row++ {
		for seat := 0; seat < numSeats; seat++ {
			if seatRows[row][seat] == '#' {
				numOccupied++
			}
		}
	}
	return numOccupied
}

func printSeatLayout(info string, seatRows [][]byte) {
	fmt.Println(info)
	// for _, row := range seatRows {
	// 	fmt.Println(string(row))
	// }
}

// Returns:
// Has any seats changes status
// The new seating arrangement
func performSeatChange(seatRows [][]byte) (bool, [][]byte) {
	var seatRowsCopy [][]byte
	seatChange := false
	for rowIndex := 0; rowIndex < numRows; rowIndex++ {
		rowCopy := make([]byte, numSeats)
		for seatIndex := 0; seatIndex < numSeats; seatIndex++ {
			seatStatus := seatRows[rowIndex][seatIndex]
			if seatStatus == '.' {
				rowCopy[seatIndex] = seatStatus
				continue // Floor
			}

			occupiedAjacentSeats := countOccupiedAjacentSeats(rowIndex, seatIndex, seatRows)
			if occupiedAjacentSeats == 0 {
				rowCopy[seatIndex] = '#' // occupied
				if rowCopy[seatIndex] != seatStatus {
					seatChange = true
				}
			} else if occupiedAjacentSeats >= 4 {
				rowCopy[seatIndex] = 'L' // free
				if rowCopy[seatIndex] != seatStatus {
					seatChange = true
				}
			} else {
				rowCopy[seatIndex] = seatStatus
			}
		}
		seatRowsCopy = append(seatRowsCopy, rowCopy)
	}

	return seatChange, seatRowsCopy
}

func countOccupiedAjacentSeats(rowIndex int, seatIndex int, seatRows [][]byte) int {
	occupiedSeats := 0

	occupiedSeats += checkIfSeatIsOccupied(rowIndex-1, seatIndex-1, seatRows)
	occupiedSeats += checkIfSeatIsOccupied(rowIndex-1, seatIndex, seatRows)
	occupiedSeats += checkIfSeatIsOccupied(rowIndex-1, seatIndex+1, seatRows)

	occupiedSeats += checkIfSeatIsOccupied(rowIndex, seatIndex-1, seatRows)
	occupiedSeats += checkIfSeatIsOccupied(rowIndex, seatIndex+1, seatRows)

	occupiedSeats += checkIfSeatIsOccupied(rowIndex+1, seatIndex-1, seatRows)
	occupiedSeats += checkIfSeatIsOccupied(rowIndex+1, seatIndex, seatRows)
	occupiedSeats += checkIfSeatIsOccupied(rowIndex+1, seatIndex+1, seatRows)

	return occupiedSeats
}

// returns 1 if seat is occupied, otherwise 0
func checkIfSeatIsOccupied(rowIndex int, seatIndex int, seatRows [][]byte) int {
	if rowIndex < 0 || rowIndex >= numRows {
		return 0
	}
	if seatIndex < 0 || seatIndex >= numSeats {
		return 0
	}
	if seatRows[rowIndex][seatIndex] == '#' {
		return 1
	}
	return 0
}

func parseInput(fileLines []string) ([][]byte, error) {

	var seterader [][]byte

	for _, seatRow := range fileLines {
		rad := make([]byte, len(seatRow))
		for seatIndex := 0; seatIndex < len(seatRow); seatIndex++ {
			rad[seatIndex] = seatRow[seatIndex]
		}
		seterader = append(seterader, rad)
	}
	return seterader, nil
}
