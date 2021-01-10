package main

import (
	"fmt"
	"jonthomas/AdventOfCode2020/datastructures"
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
		thisLineAnswer, _, err = findNextValue(line)
		if err != nil {
			fmt.Println(err)
			return
		}
		answer += thisLineAnswer
		fmt.Printf("\nLine %d: This line = %d. Total answer = %d\n", lineNr, thisLineAnswer, answer)
	}

	fmt.Printf("END. Answer = %d\n", answer)
}

// Had good help from https://leetcode.com/problems/basic-calculator-ii/solution/ for this one
func findNextValue(val expression) (int, int, error) {

	stack := new(datastructures.Stack)
	currentNumber := 0
	answer := 1

	for i := 0; i < len(val); i++ {
		next := val[i]
		fmt.Printf(" %s", string(next))

		number, _ := strconv.Atoi(string(next))
		if number != 0 {
			currentNumber = number
			continue
		}

		switch next {
		case '*':
			addedNums, err := addAllNumbers(stack)
			if err != nil {
				return -1, -1, err
			}
			answer *= (currentNumber + addedNums)
		case '+':
			stack.Push(currentNumber)
		case '(':
			subValue, j, err := findNextValue(val[i+1:])
			if err != nil {
				return -1, -1, err
			}
			i += j
			currentNumber = subValue
		case ')':
			addedNums, err := addAllNumbers(stack)
			if err != nil {
				return -1, -1, err
			}
			answer *= (currentNumber + addedNums)
			return answer, i + 1, nil
		default:
			fmt.Printf("\nfindNextValue: Operation %c is unknown (%s)", next, val)
		}
	}

	addedNums, err := addAllNumbers(stack)
	if err != nil {
		return -1, -1, err
	}
	answer *= (currentNumber + addedNums)

	return answer, -1, nil
}

func addAllNumbers(stack *datastructures.Stack) (int, error) {
	subVal := 0
	for stack.Len() > 0 {
		v, err := stack.Pop()
		if err == nil {
			subVal += v
		} else {
			return 0, err
		}
	}
	return subVal, nil
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
