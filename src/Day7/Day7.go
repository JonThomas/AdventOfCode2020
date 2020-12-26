package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type luggageRule struct {
	Inner string
	Outer string
}

func main() {

	rules, err := scanMap("./Day7input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	luggageRules := parseRules(rules)

	result := 0
	partialResult := 0
	var alreadyFound []string

	for index, rule := range luggageRules {
		fmt.Printf("%d: %s (inner) is contained in %s (outer) \n", index+1, rule.Inner, rule.Outer)

		if rule.Inner == "shiny gold bag" {
			partialResult, alreadyFound = recurse(rule.Outer, luggageRules, alreadyFound, strconv.Itoa(index+1), result)
			result = partialResult
		}
	}

	fmt.Println("Antall Outer: ", result)
}

func recurse(outer string, allRules []luggageRule, alreadyFound []string, indent string, result int) (int, []string) {
	if isAlreadyFound(outer, alreadyFound) {
		return result, alreadyFound
	}
	fmt.Println("Looking for ", outer)
	for index, rule := range allRules {
		fmt.Printf("%s: %s (inner) is contained in %s (outer) \n", indent+"."+strconv.Itoa(index+1), rule.Inner, rule.Outer)
		if rule.Inner == outer {
			partialResult := 0
			partialResult, alreadyFound = recurse(rule.Outer, allRules, alreadyFound, indent+"."+strconv.Itoa(index+1), result)
			result = partialResult
		}
	}
	fmt.Printf("%s has not any parents. Adding 1 to result\n", outer)
	alreadyFound = append(alreadyFound, outer)
	return result + 1, alreadyFound
}

func isAlreadyFound(ruleToCheck string, alreadyFound []string) bool {
	for _, found := range alreadyFound {
		if found == ruleToCheck {
			return true
		}
	}
	return false
}

func parseRules(rules []string) []luggageRule {
	var ruleList []luggageRule

	for _, rule := range rules {

		rule = strings.TrimRight(rule, ".")
		rule = strings.ReplaceAll(rule, "bags", "bag")

		outerAndInner := strings.Split(rule, " contain ")

		if outerAndInner[1] != "no other bag" {
			parsedRule := luggageRule{}
			inner := strings.Split(outerAndInner[1], ", ")
			for i := 0; i < len(inner); i++ {
				inner[i] = inner[i][2:]
				parsedRule.Outer = outerAndInner[0]
				parsedRule.Inner = inner[i]
				ruleList = append(ruleList, parsedRule)
			}
		}
	}
	return ruleList
}

func scanMap(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var rules []string

	for scanner.Scan() {
		var line = scanner.Text()
		rules = append(rules, line)
	}

	return rules, nil
}
