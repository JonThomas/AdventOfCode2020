package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	snowMap, err := scanMap("./Day3input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	numberOfTreesHit := 0
	xStep := 1
	yStep := 2

	pos := xStep
	wrapPos := len(snowMap[0])
	for line := yStep; line < len(snowMap); line += yStep {
		snowRow := snowMap[line]
		tree := snowRow[pos%wrapPos]
		fmt.Printf("Linje %d: %s Pos %d (%d) er %s\n", line+1, snowRow, pos+1, pos%wrapPos+1, string(tree))
		if tree == '#' {
			numberOfTreesHit++
		}
		pos += xStep
	}

	fmt.Println("Antall trær som ble truffet: ", numberOfTreesHit)
}

func scanMap(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var mapPart []string

	for scanner.Scan() {
		var line = scanner.Text()
		mapPart = append(mapPart, line)
	}

	return mapPart, nil
}
