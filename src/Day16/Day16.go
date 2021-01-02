package main

import (
	"fmt"
	"jonthomas/AdventOfCode2020/files"
	"strconv"
	"strings"
)

type ticketInfo struct {
	field  string
	start1 int
	end1   int
	start2 int
	end2   int
}

type ticket []int

func main() {

	file, err := files.ReadFile("./Day16Input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	ticketNotes, _ := parseTicketNotes(file)
	tickets, err := parseTickets(file)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	printParsedInput(ticketNotes, tickets)

	smallesNoteNumber := findSmallest(ticketNotes)
	largestNoteNumber := findLargest(ticketNotes)

	ticketErrorRate := 0

	// Check if any *nearby* tickets are valid
	// Skipping my ticket in row 0
	for i := 1; i < len(tickets); i++ {
		ticketErrorRate += findTicketErrorRate(tickets[i], smallesNoteNumber, largestNoteNumber)
	}

	fmt.Printf("\nFound solution: Number of invalid tickets = %d \n", ticketErrorRate)
}

func findTicketErrorRate(t ticket, smallesNoteNumber int, largestNoteNumber int) int {
	errorRate := 0
	for _, num := range t {
		if num < smallesNoteNumber || num > largestNoteNumber {
			errorRate += num
		}
	}
	return errorRate
}

func findSmallest(info []ticketInfo) int {
	smallest := 999999
	for _, i := range info {
		if i.start1 < smallest {
			smallest = i.start1
		}
	}
	return smallest
}

func findLargest(info []ticketInfo) int {
	largest := 0
	for _, i := range info {
		if i.end2 > largest {
			largest = i.end2
		}
	}
	return largest
}

func printParsedInput(ticketNotes []ticketInfo, tickets []ticket) {
	for _, note := range ticketNotes {
		fmt.Printf("%s: %d-%d | %d-%d\n", note.field, note.start1, note.end1, note.start2, note.end2)
	}
	fmt.Println()
	for _, ticket := range tickets {
		for _, nr := range ticket {
			fmt.Printf("%d, ", nr)
		}
		fmt.Println()
	}
}

func parseTicketNotes(fileLines []string) ([]ticketInfo, error) {

	var allTicketInfo []ticketInfo

	for _, line := range fileLines {
		if line == "your tickets:" || line == "" {
			break
		}

		var thisTicketInfo ticketInfo

		fieldAndRanges := strings.Split(line, ": ")
		thisTicketInfo.field = fieldAndRanges[0]

		ranges := strings.Split(fieldAndRanges[1], " or ")
		start, end, err := splitRange(ranges[0])
		if err != nil {
			return nil, err
		}
		thisTicketInfo.start1 = start
		thisTicketInfo.end1 = end
		start, end, err = splitRange(ranges[1])
		if err != nil {
			return nil, err
		}
		thisTicketInfo.start2 = start
		thisTicketInfo.end2 = end

		allTicketInfo = append(allTicketInfo, thisTicketInfo)
	}

	return allTicketInfo, nil
}

func splitRange(dashSeparated string) (int, int, error) {
	startEnd := strings.Split(dashSeparated, "-")
	start, numErr1 := strconv.Atoi(startEnd[0])
	if numErr1 != nil {
		return 0, 0, numErr1
	}
	end, numErr2 := strconv.Atoi(startEnd[1])
	if numErr2 != nil {
		return 0, 0, numErr2
	}
	return start, end, nil
}

func parseTickets(fileLines []string) ([]ticket, error) {

	var allTickets []ticket

	skipThisLine := true

	for _, line := range fileLines {

		if line == "your ticket:" {
			skipThisLine = false
			continue
		}

		if skipThisLine || line == "" || line == "nearby tickets:" {
			continue
		}

		var thisTicket ticket

		numbers := strings.Split(line, ",")
		for _, numStr := range numbers {
			num, numErr := strconv.Atoi(numStr)
			if numErr != nil {
				return nil, numErr
			}
			thisTicket = append(thisTicket, num)
		}

		allTickets = append(allTickets, thisTicket)
	}

	return allTickets, nil
}
