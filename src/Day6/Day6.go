package main

import (
	"bufio"
	"fmt"
	"os"
)

type group struct {
	// Et array av byte-arrays
	PersonAnswers [][]byte
}

func main() {

	groups, err := scanAnswers("./Day6input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	sum := 0

	for groupIndex, group := range groups {
		result := group.PersonAnswers[0]

		for personIndex, personAnswers := range group.PersonAnswers {

			if personIndex == 0 {
				continue
			}

			result = intersect(personAnswers, result)
		}

		sum += len(result)
		fmt.Printf("Group %d: %d (%d) %s %s\n", groupIndex, len(result), sum, group, result)
	}
}

func intersect(a, b []byte) (c []byte) {
	m := make(map[byte]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

func scanAnswers(path string) ([]group, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	groups := []group{}
	groups = append(groups, group{})

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	groupIndex := 0

	for scanner.Scan() {
		var line = scanner.Text()

		if line == "" {
			// new passport
			groups = append(groups, group{})
			groupIndex++
			continue
		}

		groups[groupIndex].PersonAnswers = append(groups[groupIndex].PersonAnswers, []byte(line))

	}

	return groups, nil
}
