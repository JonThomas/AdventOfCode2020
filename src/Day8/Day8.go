package main

import (
	"fmt"
	"jonthomas/AdventOfCode2020/files"
	"strconv"
)

type instruction struct {
	Operator string
	Count    int
	Visited  bool
}

// program is a key-value data structure, with pointers to instructions.
// Pointers makes it possible to make changes to instructions
type program map[int]*instruction

func main() {

	file, err := files.ReadFile("./Day8Input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	program, parseErr := parseInput(file)
	if parseErr != nil {
		fmt.Println("parsing error", err)
		return
	}

	currentInstruction := program[0]
	currentIndex := 0
	accumulator := 0

	for {
		fmt.Printf("Current instruction: %d - %v\n", currentIndex, program[currentIndex])
		if program[currentIndex].Visited {
			break
		}

		program[currentIndex].Visited = true

		switch currentInstruction.Operator {
		case "nop":
			currentIndex++
			currentInstruction = program[currentIndex]
			continue
		case "acc":
			currentIndex++
			accumulator += currentInstruction.Count
			currentInstruction = program[currentIndex]
			continue
		case "jmp":
			currentIndex += currentInstruction.Count
			currentInstruction = program[currentIndex]
			continue
		default:
			fmt.Println("Unknown operator error: ", currentInstruction.Operator)
			return
		}
	}

	fmt.Println("Done. Accumulator = ", accumulator)
}

func parseInput(fileLines []string) (program, error) {

	var program = make(program)

	for fileIndex, line := range fileLines {
		instruction := instruction{}
		instruction.Operator = line[0:3]
		nr, nrErr := strconv.Atoi(line[4:])
		if nrErr != nil {
			return nil, nrErr
		}
		instruction.Count = nr
		instruction.Visited = false
		program[fileIndex] = &instruction

		//fmt.Println(fileIndex, ": ", instruction, " - ", line)
	}
	return program, nil
}
