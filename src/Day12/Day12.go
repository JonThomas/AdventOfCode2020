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

	// 0 = north, 90 = east, ...
	direction := 90
	xPos := 0
	yPos := 0

	instructionNr := 1

	for _, instruction := range instructions {
		switch instruction.action {
		case "N":
			yPos -= instruction.value
		case "E":
			xPos += instruction.value
		case "S":
			yPos += instruction.value
		case "W":
			xPos -= instruction.value
		case "L":
			direction -= instruction.value
			direction = modulus(direction)
		case "R":
			direction += instruction.value
			direction = modulus(direction)
		case "F":
			switch direction {
			case 0:
				yPos -= instruction.value
			case 90:
				xPos += instruction.value
			case 180:
				yPos += instruction.value
			case 270:
				xPos -= instruction.value
			default:
				fmt.Println("Direction ", direction, " is invalid")
				return
			}
		default:
			fmt.Println("Action ", instruction.action, " is invalid")
			return
		}
		fmt.Printf("%d - %v: Now at %d,%d facing %d\n", instructionNr, instruction, xPos, yPos, direction)
		instructionNr++
	}

	fmt.Printf("Manhattan number = abs(%d) + abs(%d) = %d\n", xPos, yPos, abs(xPos)+abs(yPos))
	fmt.Println("END")
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
