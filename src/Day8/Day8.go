package main

import (
	"fmt"
	"jonthomas/AdventOfCode2020/files"
)

type fileLines struct {
	Outer        string
	Inner        string
	HowManyInner int
}

func main() {

	file, err := files.ReadFile("./Day8input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	parseErr := parseInput(file)
	if parseErr != nil {
		fmt.Println("parsing error", err)
		return
	}
}

func parseInput(fileLines []string) error {

	for _, line := range fileLines {
		fmt.Println(line)
	}
	return nil
}
