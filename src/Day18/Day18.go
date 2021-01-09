package main

import (
	"fmt"
	"jonthomas/AdventOfCode2020/files"
	"strconv"
	"strings"
)

type expression []byte

func main() {

	file, err := files.ReadFile("./Day18Input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	parsedInput := parseInput(file)

	printParsedInput(parsedInput)

	answer := 0
	thisLineAnswer := 0
	for lineNr, line := range parsedInput {
		thisLineAnswer, _ = findNextValue(line)
		answer += thisLineAnswer
		fmt.Printf("\nLine %d: This line = %d. Total answer = %d\n", lineNr, thisLineAnswer, answer)
	}

	fmt.Printf("END. Answer = %d\n", answer)
}

// Returns (answer, new index into expression)
func findNextValue(val expression) (int, int) {

	currentAnswer := 0
	var currentOperation byte = ' '

	for i := 0; i < len(val); i++ {
		next := val[i]
		fmt.Printf(" %s", string(next))

		number, _ := strconv.Atoi(string(next))
		if number != 0 {
			currentAnswer, currentOperation = calculate(currentOperation, number, currentAnswer)
			continue
		}
		switch next {
		case '+':
			currentOperation = '+'
		case '*':
			currentOperation = '*'
		case '(':
			subValue, j := findNextValue(val[i+1:])
			i += j
			currentAnswer, currentOperation = calculate(currentOperation, subValue, currentAnswer)
		case ')':
			return currentAnswer, i + 1
		default:
			fmt.Printf("\nfindNextValue: Operation %c is unknown (%s)", next, val)
		}
	}
	return currentAnswer, -1
}

func calculate(currentOperation byte, number int, currentAnswer int) (int, byte) {
	if currentOperation == '+' {
		currentAnswer += number
		currentOperation = ' '
	} else if currentOperation == '*' {
		if currentAnswer == 0 {
			currentAnswer = 1
		}
		currentAnswer *= number
		currentOperation = ' '
	} else if currentOperation == ' ' {
		currentAnswer = number
	}
	return currentAnswer, currentOperation
}

func parseInput(file []string) []expression {
	var expressions []expression
	for _, line := range file {
		line = strings.ReplaceAll(line, " ", "")
		expressions = append(expressions, []byte(line))
	}
	return expressions
}

func printParsedInput(expressions []expression) {
	for idx, expression := range expressions {
		fmt.Printf("%d: %s\n", idx+1, expression)
	}
}
