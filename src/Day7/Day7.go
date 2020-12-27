package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type luggageRule struct {
	Outer        string
	Inner        string
	HowManyInner int
}

func main() {

	rules, err := scanRules("./Day7input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	luggageRules, parseErr := parseRules(rules)
	if parseErr != nil {
		fmt.Println("parsing error", err)
		return
	}

	result := 0

	for index, rule := range luggageRules {
		fmt.Printf("%d: %s contains %d %s\n", index+1, rule.Outer, rule.HowManyInner, rule.Inner)

		if rule.Outer == "shiny gold bag" {
			result += recurse(rule.Inner, luggageRules, strconv.Itoa(index+1)) * rule.HowManyInner
		}
	}

	fmt.Println("Antall: ", result)
}

func recurse(inner string, allRules []luggageRule, indent string) int {
	result := 0
	for index, rule := range allRules {
		fmt.Printf("%s: Looking for %s in %s -> %s \n", indent+"."+strconv.Itoa(index+1), inner, rule.Outer, rule.Inner)
		if rule.Outer == inner {
			if rule.Inner == "" {
				return 1
			}
			result += recurse(rule.Inner, allRules, indent+"."+strconv.Itoa(index+1)) * rule.HowManyInner
		}
	}
	return result + 1 // + 1 for å ta med seg selv
}

func parseRules(rules []string) ([]luggageRule, error) {
	var ruleList []luggageRule

	for _, rule := range rules {

		rule = strings.TrimRight(rule, ".")
		rule = strings.ReplaceAll(rule, "bags", "bag")

		outerAndInner := strings.Split(rule, " contain ")

		parsedRule := luggageRule{}
		inner := strings.Split(outerAndInner[1], ", ")
		for i := 0; i < len(inner); i++ {
			num := 0
			var err error
			innerBag := ""
			if inner[i] != "no other bag" {
				num, err = strconv.Atoi(inner[i][0:1])
				if err != nil {
					return nil, err
				}
				innerBag = inner[i][2:]
			}
			parsedRule.Outer = outerAndInner[0]
			parsedRule.Inner = innerBag
			parsedRule.HowManyInner = num

			ruleList = append(ruleList, parsedRule)
		}
	}
	return ruleList, nil
}

func scanRules(path string) ([]string, error) {
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
