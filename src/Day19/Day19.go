package main

import (
	"fmt"
	"jonthomas/AdventOfCode2020/files"
	"strconv"
	"strings"
)

// Rules are parsed into these types
type piped struct {
	before string
	after  string
}
type rules map[int]piped

func main() {

	file, err := files.ReadFile("./Day19Input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	parsedRules, err := parseRules(file)
	printParsedRules(parsedRules)
	if err != nil {
		fmt.Println("Parsing error: ", err)
		return
	}

	parsedMessages := parseMessages(file)
	printParsedMessages(parsedMessages)

	abRules, err := findAbRules(parsedRules)

	answer := 0
	for _, msg := range parsedMessages {
		for _, rule := range abRules {
			if msg == rule {
				answer++
				break
			}
		}
	}

	fmt.Printf("END. Answer = %d\n", answer)
}

func findAbRules(parsedRules rules) ([]string, error) {

	// Rule 0.
	var flattenedNumbers [][]int
	before, after, err := fetchNumbers(parsedRules[0])
	if err != nil {
		return nil, fmt.Errorf("Fetch number[0] error: ", err)
	}
	flattenedNumbers = append(flattenedNumbers, before)
	if after != nil {
		flattenedNumbers = append(flattenedNumbers, after)
	}

	// Loop through each number until we're only left with "a" and "b" equivalents
	// Creates a copy of flattenedNumbers on loop, and replaces flattenedNumbers with the copy at the end
	// Algrithm:
	// 	Loop through each number in flattenedNumbers
	// 	Each time a rule with | is encounterd, copy current line of numbers, and replace rule with new rules.
	round := 2
	needMoreReplacements := true

	for needMoreReplacements {
		needMoreReplacements = false
		var flattenedNumbersCopy [][]int
		for y := 0; y < len(flattenedNumbers); y++ {
			yCopy := len(flattenedNumbersCopy)
			flattenedNumbersCopy = append(flattenedNumbersCopy, []int{})

			for x := 0; x < len(flattenedNumbers[y]); x++ {
				thisNumber := flattenedNumbers[y][x]

				// Case 1: rule contains a letter
				if parsedRules[thisNumber].before == "\"a\"" || parsedRules[thisNumber].before == "\"b\"" {
					flattenedNumbersCopy[yCopy] = append(flattenedNumbersCopy[yCopy], thisNumber)
					continue
				}
				needMoreReplacements = true

				before, after, err := fetchNumbers(parsedRules[thisNumber])
				if err != nil {
					return nil, fmt.Errorf("fetch rules for number[%d] error: %v", thisNumber, err)
				}

				// Case 2: rule contains one set of numbers (no 'OR')
				if after == nil {
					flattenedNumbersCopy[yCopy] = append(flattenedNumbersCopy[yCopy], before...)
					// Case 3: rule contains two set of numbers (an 'OR' is present)
				} else {
					// copy row. replace
					nextLine := createNextLine(flattenedNumbersCopy[yCopy], after, flattenedNumbers[y][x+1:])
					flattenedNumbersCopy[yCopy] = append(flattenedNumbersCopy[yCopy], before...)

					flattenedNumbersCopy = append(flattenedNumbersCopy, nextLine)
				}
			}
		}
		//printFlattenedNumbers(flattenedNumbersCopy, round)
		round++
		flattenedNumbers = flattenedNumbersCopy
	}

	ruleA, err := findRule("\"a\"", parsedRules)
	if err != nil {
		return nil, fmt.Errorf("find rule error: ", err)
	}
	ruleB, err := findRule("\"b\"", parsedRules)
	if err != nil {
		return nil, fmt.Errorf("find rule error: ", err)
	}

	fmt.Printf("Rule A: %d. Rule B: %d\n", ruleA, ruleB)

	abRules, err := convertFromNumbersToLetters(flattenedNumbers, ruleA, ruleB)
	if err != nil {
		return nil, fmt.Errorf("Error converting from numbers to letters: ", err)
	}

	return abRules, nil
}

func convertFromNumbersToLetters(flattenedNumbers [][]int, ruleA int, ruleB int) ([]string, error) {
	var abRules []string
	for i := 0; i < len(flattenedNumbers); i++ {
		thisRule := strings.Builder{}
		for _, num := range flattenedNumbers[i] {
			if num == ruleA {
				thisRule.WriteRune('a')
			} else if num == ruleB {
				thisRule.WriteRune('b')
			} else {
				return nil, fmt.Errorf("number %d is neither A nor B", num)
			}
		}
		thisRuleString := thisRule.String()
		//fmt.Println(thisRuleString)
		abRules = append(abRules, thisRuleString)
	}
	return abRules, nil
}

func findRule(rule string, allRules rules) (int, error) {
	for ruleNum, thisRule := range allRules {
		if thisRule.before == rule {
			return ruleNum, nil
		}
	}
	return 0, fmt.Errorf("Rule %s is not found", rule)
}

func createNextLine(start []int, middle []int, end []int) []int {
	var ret []int

	ret = append(ret, start...)
	ret = append(ret, middle...)
	ret = append(ret, end...)

	return ret
}

func printFlattenedNumbers(flattenedNumbersCopy [][]int, round int) {
	fmt.Println()
	fmt.Printf("Round %d\n", round)
	for y := 0; y < len(flattenedNumbersCopy); y++ {
		for x := 0; x < len(flattenedNumbersCopy[y]); x++ {
			fmt.Printf("%d ", flattenedNumbersCopy[y][x])
		}
		fmt.Println()
	}
}

func fetchNumbers(beforeAndAfter piped) ([]int, []int, error) {
	var beforeRet []int
	before := beforeAndAfter.before
	numberStrings := strings.Split(before, " ")
	for _, num := range numberStrings {
		thisNum, err := strconv.Atoi(num)
		if err != nil {
			return nil, nil, err
		}
		beforeRet = append(beforeRet, thisNum)
	}

	var afterRet []int
	after := beforeAndAfter.after
	if after == "" {
		return beforeRet, nil, nil
	}
	numberStrings = strings.Split(after, " ")
	for _, num := range numberStrings {
		thisNum, err := strconv.Atoi(num)
		if err != nil {
			return nil, nil, err
		}
		afterRet = append(afterRet, thisNum)
	}
	return beforeRet, afterRet, nil
}

func printParsedRules(parsedRules rules) {
	for index, r := range parsedRules {
		fmt.Printf("%d: %v\n", index, r)
	}
}

func printParsedMessages(parsedMessages []string) {
	for _, msg := range parsedMessages {
		fmt.Println(msg)
	}
}

func parseMessages(file []string) []string {
	var ret []string
	startReading := false
	for _, line := range file {
		if line == "" {
			startReading = true
			continue
		}
		if startReading {
			ret = append(ret, line)
		}
	}
	return ret
}

func parseRules(file []string) (rules, error) {
	var allRules = make(map[int]piped)
	for _, line := range file {
		if line == "" {
			return allRules, nil
		}
		numberAndReferences := strings.Split(line, ": ")
		number, err := strconv.Atoi(numberAndReferences[0])
		if err != nil {
			return nil, err
		}
		beforeAndAfter := strings.Split(numberAndReferences[1], " | ")
		if len(beforeAndAfter) == 1 {
			allRules[number] = piped{before: beforeAndAfter[0], after: ""}
		} else {
			allRules[number] = piped{before: beforeAndAfter[0], after: beforeAndAfter[1]}
		}
	}
	return allRules, nil
}
