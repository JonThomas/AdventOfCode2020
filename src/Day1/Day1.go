package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("Hello World")
	// data, err := ioutil.ReadFile("Day1input.txt")
	lines, err := scanNumbers("./Day1input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	numElements := len(lines)

	for i := 0; i < numElements; i++ {
		for j := 1; j < numElements; j++ {
			for k := 2; k < numElements; k++ {
				sum := lines[i] + lines[j] + lines[k]
				fmt.Println("Sammenligner", lines[i], ", ", lines[j], " og ", lines[k], ": ", sum)
				if sum == 2020 {
					fmt.Println(lines[i] * lines[j] * lines[k])
					return
				}
			}
		}
	}

}

func scanNumbers(path string) ([]int, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var lines []int

	for scanner.Scan() {
		var t = scanner.Text()
		var num, _ = strconv.Atoi(t)
		//fmt.Println(num)
		lines = append(lines, num)
	}

	return lines, nil
}
