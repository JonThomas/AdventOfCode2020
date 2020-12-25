package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	rule, pwd, err := scanRules("./Day2input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	numElements := len(rule)
	if numElements != len(pwd) {
		log.Fatal("Forskjellig antall elementer i rule: ", numElements, " og pwd: ", len(pwd))
		return
	}

	validPwds := 0

	for i := 0; i < numElements; i++ {
		min, max, letter := parseRule(rule[i])
		fmt.Print("min: ", min, " max: ", max, " letter: ", letter, " pwd: ", pwd[i])
		// occurrencesInPwd := strings.Count(pwd[i], letter)
		// if min <= occurrencesInPwd && max >= occurrencesInPwd {
		// 	validPwds++
		// }
		thisPwd := pwd[i]
		letterAtPos1 := string(thisPwd[min])
		letterAtPos2 := string(thisPwd[max])
		fmt.Print(" En: ", letterAtPos1, " To: ", letterAtPos2)
		if (letterAtPos1 == letter && letterAtPos2 != letter) || (letterAtPos1 != letter && letterAtPos2 == letter) {
			validPwds++
			fmt.Println("\t\tOK!!")
		} else {
			fmt.Println("\t\tFEIL!!")
		}
	}

	fmt.Println("Antall gyldig passord: ", validPwds)
}

// Rune = En bokstav - tilsvare C# char
func parseRule(rule string) (int, int, string) {
	rangeAndLetter := strings.Split(rule, " ")
	minMax := strings.Split(rangeAndLetter[0], "-")
	min, _ := strconv.Atoi(minMax[0])
	max, _ := strconv.Atoi(minMax[1])
	return min, max, rangeAndLetter[1]
}

func scanRules(path string) ([]string, []string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var rules []string
	var pwds []string

	for scanner.Scan() {
		var line = scanner.Text()
		var lineArr = strings.Split(line, ":")

		rules = append(rules, lineArr[0])
		pwds = append(pwds, lineArr[1])
	}

	return rules, pwds, nil
}
