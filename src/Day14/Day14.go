package main

import (
	"fmt"
	"jonthomas/AdventOfCode2020/files"
	"math"
	"strconv"
	"strings"
)

type memwrite struct {
	address int
	number  int
	mask    string
}

func main() {

	memory := make(map[int]int)

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

	for memIndex, writeOp := range memwrite {
		addresses, err := calculateMaskedValue(writeOp.address, writeOp.mask)
		if err != nil {
			fmt.Println("Memory caluclation error: ", err)
			return
		}
		for _, address := range addresses {
			memory[address] = writeOp.number
		}
		fmt.Printf("%d: Added %d addresses.\n", memIndex, len(addresses))
	}

	var theAnswer int = 0
	for _, memValue := range memory {
		theAnswer += memValue
	}

	fmt.Printf("\nFound solution: %d\n", theAnswer)
}

func calculateMaskedValue(address int, mask string) ([]int, error) {

	maskLen := 36
	var newBinaryAddress strings.Builder
	newBinaryAddress.Grow(maskLen)

	addressBinary := strconv.FormatInt(int64(address), 2)
	addressBinaryPadded := fmt.Sprintf("%036v", addressBinary) // %036 means: Pads with 0's to 36 characters

	for i := 0; i < maskLen; i++ {
		switch mask[i] {
		case '0':
			digit := addressBinaryPadded[i : i+1]
			newBinaryAddress.WriteString(digit)
		case '1':
			newBinaryAddress.Write([]byte{'1'})
		case 'X':
			newBinaryAddress.Write([]byte{'X'})
		default:
			return nil, fmt.Errorf("Unknown mask value: %s" + string(mask[i]))
		}
	}

	newBinaryAddressString := newBinaryAddress.String()
	addresses, err := generateAddresses(newBinaryAddressString)
	if err != nil {
		return nil, err
	}

	return addresses, nil
}

func generateAddresses(newBinaryNumber string) ([]int, error) {
	var addresses []int

	numberOfXs := strings.Count(newBinaryNumber, "X")
	numberOfaddresses := int(math.Pow(2, float64(numberOfXs)))
	for i := 0; i < numberOfaddresses; i++ {

		// 1: Convert i to binary number, padded to length (numberOfXs)
		binaryI := strconv.FormatInt(int64(i), 2)
		formatMask := "%" + fmt.Sprintf("0%dv", numberOfXs)
		binaryI = fmt.Sprintf(formatMask, binaryI)

		// 2: Replace Xs in copy with each digit in binary representation of i
		addressBinary := replaceXs(newBinaryNumber, binaryI)

		// 3: Converty binary string to decimal
		address, err := strconv.ParseInt(addressBinary, 2, 64)
		if err != nil {
			return nil, err
		}

		addresses = append(addresses, int(address))
	}

	return addresses, nil
}

func replaceXs(binaryNumberWithXs string, willOverwriteXs string) string {
	newAddress := binaryNumberWithXs
	for _, char := range willOverwriteXs {
		newAddress = strings.Replace(newAddress, "X", string(char), 1)
	}
	return newAddress
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
