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

	file, err := files.ReadFile("./Day13Input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	departureTime, buses, parseErr := parseInput(file)
	if parseErr != nil {
		fmt.Println("parsing error", err)
		return
	}

	printParsedInput(departureTime, buses)

	//minutesUntilArrival := make(busArrivaltime)
	minutesUntilArrival := calulateArrivalTimes(departureTime, buses)

	// Shortest arrivaltime
	shortestWait := 9999999999999
	busNr := 0
	for bus, minUntilArrival := range minutesUntilArrival {
		if minUntilArrival < shortestWait {
			shortestWait = minUntilArrival
			busNr = bus
		}
	}

	fmt.Printf("Answer: busId (%d) * shortestWait (%d) = %d", busNr, shortestWait, busNr*shortestWait)
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

func printParsedInput(timestamp int, buses []int) {
	fmt.Println("Timestamp: ", timestamp)
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
