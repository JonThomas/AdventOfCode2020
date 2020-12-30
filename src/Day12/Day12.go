package main

import (
	"fmt"
	"jonthomas/AdventOfCode2020/files"
	"strconv"
)

type instruction struct {
	action string
	value  int
}

func main() {

	file, err := files.ReadFile("./Day12Input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	instructions, parseErr := parseInput(file)
	if parseErr != nil {
		fmt.Println("parsing error", err)
		return
	}

	//          North (-)
	//              ^
	//              |
	// West (-) <-------> East (+)
	//              |
	//          South (+)

	xPos := 0
	yPos := 0

	waypointxPos := 10
	waypointyPos := -1

	instructionNr := 1

	for _, instruction := range instructions {
		switch instruction.action {
		case "N":
			waypointyPos -= instruction.value
		case "E":
			waypointxPos += instruction.value
		case "S":
			waypointyPos += instruction.value
		case "W":
			waypointxPos -= instruction.value
		case "L":
			waypointxPos, waypointyPos = rotateLeft(instruction.value, waypointxPos, waypointyPos)
		case "R":
			waypointxPos, waypointyPos = rotateRight(instruction.value, waypointxPos, waypointyPos)
		case "F":
			xPos += waypointxPos * instruction.value
			yPos += waypointyPos * instruction.value
		default:
			fmt.Println("Action ", instruction.action, " is invalid")
			return
		}
		fmt.Printf("%d - %v: Waypoint at %d, %d. Ship at %d,%d\n", instructionNr, instruction, waypointxPos, waypointyPos, xPos, yPos)
		instructionNr++
	}

	fmt.Printf("Manhattan number = abs(%d) + abs(%d) = %d\n", xPos, yPos, abs(xPos)+abs(yPos))
	fmt.Println("END")
}

func rotateLeft(degrees int, waypointxPos int, waypointyPos int) (int, int) {
	if degrees == 0 {
		return waypointxPos, waypointyPos
	}

	tmpx := 0
	for i := 0; i < degrees/90; i++ {
		tmpx = waypointxPos
		waypointxPos = waypointyPos
		waypointyPos = -tmpx
	}
	return waypointxPos, waypointyPos
}

func rotateRight(degrees int, waypointxPos int, waypointyPos int) (int, int) {
	if degrees == 0 {
		return waypointxPos, waypointyPos
	}

	tmpx := 0
	for i := 0; i < degrees/90; i++ {
		tmpx = waypointxPos
		waypointxPos = -waypointyPos
		waypointyPos = tmpx
	}
	return waypointxPos, waypointyPos
}

// Abs returns the absolute value of x.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func modulus(value int) int {
	if value < 0 {
		return value + 360
	}
	if value >= 360 {
		return value - 360
	}
	return value
}

func parseInput(fileLines []string) ([]instruction, error) {

	var instructions []instruction

	for _, line := range fileLines {
		thisAction := line[0:1]
		thisValue, valErr := strconv.Atoi(line[1:])
		if valErr != nil {
			return nil, valErr
		}
		instruction := instruction{action: thisAction, value: thisValue}
		instructions = append(instructions, instruction)
	}
	return instructions, nil
}
