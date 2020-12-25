package main

import (
	"bufio"
	"fmt"
	"os"
)

type Group struct {
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
		var groupCombinedAnswers []byte

		for _, personAnswers := range group.PersonAnswers {
			// Sort each string
			//sort.Slice(personAnswers, func(i int, j int) bool { return personAnswers[i] < personAnswers[j] })

			for _, answer := range personAnswers {
				if notAlreadyAdded(groupCombinedAnswers, answer) {
					groupCombinedAnswers = append(groupCombinedAnswers, answer)
				}
			}
		}
		sum += len(groupCombinedAnswers)
		fmt.Printf("Group %d: %d (%d) %s %s\n", groupIndex, len(groupCombinedAnswers), sum, group, groupCombinedAnswers)
	}
}

func notAlreadyAdded(existingAnswers []byte, newAnswer byte) bool {
	for _, answer := range existingAnswers {
		if answer == newAnswer {
			return false
		}
	}
	return true
}

func scanAnswers(path string) ([]Group, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	groups := []Group{}
	groups = append(groups, Group{})

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	groupIndex := 0

	for scanner.Scan() {
		var line = scanner.Text()

		if line == "" {
			// new passport
			groups = append(groups, Group{})
			groupIndex++
			continue
		}

		groups[groupIndex].PersonAnswers = append(groups[groupIndex].PersonAnswers, []byte(line))

	}

	return groups, nil
}
