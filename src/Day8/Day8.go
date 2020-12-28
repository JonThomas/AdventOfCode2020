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

// program is a key-value data structure, with pointers to instructions.
// Pointers make it possible to make changes to instructions
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

	end := len(program)
	var originalOperation string
	var originalIndex int

	for i := end - 1; i > 0; i-- {

		// Change one instruction, so that the program run to its end
		switch program[i].Operation {
		case "nop":
			if originalOperation != "" {
				program[originalIndex].Operation = originalOperation
			}
			originalOperation = "nop"
			originalIndex = i

			program[i].Operation = "jmp"
		case "acc":
			continue
		case "jmp":
			if originalOperation != "" {
				program[originalIndex].Operation = originalOperation
			}
			originalOperation = "jmp"
			originalIndex = i

			program[i].Operation = "nop"
		default:
			fmt.Println("Unknown operator error: ", program[i].Operation)
			return
		}

		finishedSuccessfully := runProgram(program)
		if finishedSuccessfully {
			break
		}

		// Initialize program for new run
		for _, instructionToUpdate := range program {
			instructionToUpdate.Visited = false
		}
	}

	fmt.Println("END")
}

func runProgram(program program) bool {
	end := len(program)
	currentInstruction := program[0]
	currentIndex := 0
	accumulator := 0
	for {
		if currentIndex == end {
			fmt.Printf("Done at instruction %d (of %d). ", currentIndex, end)
			fmt.Println("Accumulator = ", accumulator)
			return true
		}

		if currentIndex > end {
			fmt.Printf("Done at instruction %d (of %d). ", currentIndex, end)
			fmt.Println("Accumulator = ", accumulator)
			return false
		}

		fmt.Printf("Current instruction: %d - %v\n", currentIndex, program[currentIndex])
		if program[currentIndex].Visited {
			fmt.Printf("Done at instruction %d (of %d). ", currentIndex, end)
			fmt.Println("Instruction has been hit twice. Accumulator = ", accumulator)
			return false
		}

		program[currentIndex].Visited = true

		switch currentInstruction.Operation {
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
			fmt.Println("Unknown operator error: ", currentInstruction.Operation)
			return false
		}
	}
}

func parseInput(fileLines []string) (program, error) {

	var program = make(program)

	for fileIndex, line := range fileLines {
		instruction := instruction{}
		instruction.Operation = line[0:3]
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
