package main

import (
	"fmt"
	"jonthomas/AdventOfCode2020/files"
	"strconv"
	"strings"
)

// key = busNr
// value = minutesUntilArrival
type busArrivaltime map[int]int

func main() {

	fmt.Println("Printing answer here, since it takes approx 4 hours to calulate: 760171380521445")

	file, err := files.ReadFile("./Day13Input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	_, buses, parseErr := parseInput(file)
	if parseErr != nil {
		fmt.Println("parsing error", err)
		return
	}

	printParsedInput(buses)

	largestBusNumber, largestBusPosition := findLargestBusNr(buses)

	timestamp := 0
	printSometimes := 0
	// All buses departing at offsets matching their position in the list can only happen every (largestBusNumber) timestamp
	// This is really only an optimization, to make the program run (largestBusNumber) times faster
	for timestamp = largestBusNumber; true; timestamp += largestBusNumber {
		if checkAllBuses(timestamp, largestBusPosition, buses) {
			break // Found the solution
		}
		if printSometimes%10000000 == 0 {
			fmt.Println(printSometimes)
		}
		printSometimes++
	}

	fmt.Printf("Found solution at timestamp %d", timestamp-largestBusPosition)
}

func checkAllBuses(largestBusTimestamp int, largestBusPosition int, buses []int) bool {
	// For each bus, check if its offset matches largestBusNumber
	for busOffset := 0; busOffset < len(buses); busOffset++ {
		thisBus := buses[busOffset]
		if thisBus == -1 {
			continue
		}

		if (largestBusTimestamp-largestBusPosition+busOffset)%thisBus != 0 {
			return false
		}
	}
	return true
}

func findLargestBusNr(buses []int) (int, int) {
	largestNr := 0
	largestNrIndex := 0
	for busIndex, bus := range buses {
		if bus > largestNr {
			largestNr = bus
			largestNrIndex = busIndex
		}
	}
	return largestNr, largestNrIndex
}

func calulateArrivalTimes(departureTime int, buses []int) busArrivaltime {
	minutesUntilArrival := make(busArrivaltime)
	for _, busNr := range buses {
		leftMinutesAgo := departureTime % busNr
		timeUntilDeparture := busNr - leftMinutesAgo
		minutesUntilArrival[busNr] = timeUntilDeparture
	}
	return minutesUntilArrival
}

func printParsedInput(buses []int) {
	fmt.Printf("Buses: ")
	for _, busNr := range buses {
		fmt.Printf("%d, ", busNr)
	}
	fmt.Println()
}

func parseInput(fileLines []string) (int, []int, error) {
	var buses []int

	ts, nrErr := strconv.Atoi(fileLines[0])
	if nrErr != nil {
		return 0, nil, nrErr
	}

	for _, number := range strings.Split(fileLines[1], ",") {
		if number == "x" {
			buses = append(buses, -1)
			continue
		}

		busNr, nrErr := strconv.Atoi(number)
		if nrErr != nil {
			return 0, nil, nrErr
		}

		buses = append(buses, busNr)
	}
	return ts, buses, nil
}
