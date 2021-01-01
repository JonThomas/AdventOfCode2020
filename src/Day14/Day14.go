package main

import (
	"fmt"
	"jonthomas/AdventOfCode2020/files"
	"strconv"
	"strings"
)

type memwrite struct {
	address int
	number  int
	mask    string
}

func main() {

	memory := make(map[int]int64)

	file, err := files.ReadFile("./Day14Input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	memwrite, parseErr := parseInput(file)
	if parseErr != nil {
		fmt.Println("parsing error", err)
		return
	}

	printParsedInput(memwrite)

	largestAddress := findLargestAddress(memwrite)

	for _, writeOp := range memwrite {
		memory[writeOp.address], err = calculateMaskedValue(writeOp.number, writeOp.mask)
		if err != nil {
			fmt.Println("Memory caluclation error: ", err)
			return
		}
	}

	var theAnswer int64 = 0
	for i := 0; i <= largestAddress; i++ {
		theAnswer += memory[i]
	}

	fmt.Printf("Found solution: %d\n", theAnswer)
}

func calculateMaskedValue(number int, mask string) (int64, error) {
	maskLen := len(mask)

	var newBinaryNumber strings.Builder
	newBinaryNumber.Grow(maskLen)

	numberBinary := strconv.FormatInt(int64(number), 2)
	numberBinaryPadded := fmt.Sprintf("%036v", numberBinary) // %036 means: Pads with 0's to 36 characters

	for i := 0; i < maskLen; i++ {
		switch mask[i] {
		case '0':
			newBinaryNumber.Write([]byte{'0'})
		case '1':
			newBinaryNumber.Write([]byte{'1'})
		case 'X':
			digit := numberBinaryPadded[i : i+1]
			newBinaryNumber.WriteString(digit)
		default:
			return 0, fmt.Errorf("Unknown mask value: %s" + string(mask[i]))
		}
	}

	binString := newBinaryNumber.String()
	dec, err := strconv.ParseInt(binString, 2, 64)
	if err != nil {
		return 0, err
	}
	return dec, nil
}

func findLargestAddress(writeOps []memwrite) int {
	largestAddress := 0
	for _, op := range writeOps {
		if op.address > largestAddress {
			largestAddress = op.address
		}
	}
	return largestAddress
}

func printParsedInput(program []memwrite) {
	for writeIdx, writeOp := range program {
		fmt.Printf("%d: %v\n", writeIdx, writeOp)
	}
}

func parseInput(fileLines []string) ([]memwrite, error) {
	var program []memwrite
	var mask string
	var err error

	for lineIndex, line := range fileLines {
		var operation memwrite
		if line[0:4] == "mask" {
			mask = line[7:]
			continue
		}
		if line[0:3] == "mem" {
			rightSquareIndex := strings.Index(line, "]")
			if rightSquareIndex == -1 {
				return nil, fmt.Errorf("No right square index at line %d", lineIndex)
			}
			operation.address, err = strconv.Atoi(line[4:rightSquareIndex])
			if err != nil {
				return nil, fmt.Errorf("Unable to extract address at line %d %v", lineIndex, err.Error())
			}
			operation.number, err = strconv.Atoi(line[rightSquareIndex+4:])
			if err != nil {
				return nil, fmt.Errorf("Unable to extract number at line %d, %v", lineIndex, err.Error())
			}
			operation.mask = mask
			program = append(program, operation)
		}
	}

	return program, nil
}
