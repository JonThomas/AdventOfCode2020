package main

import (
	"fmt"
	"jonthomas/AdventOfCode2020/files"
	"sort"
	"strconv"
	"strings"
)

type ticketInfo struct {
	field           string
	start1          int
	end1            int
	start2          int
	end2            int
	possibleColumns []int
	column          int
}

type ticketInfoList []ticketInfo

// Functions for sorting tickets
func (t ticketInfoList) Len() int {
	return len(t)
}

func (t ticketInfoList) Less(i, j int) bool {
	return len(t[i].possibleColumns) < len(t[j].possibleColumns)
}

func (t ticketInfoList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

type ticket []int

func main() {

	file, err := files.ReadFile("./Day16Input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	ticketNotes, _ := parseTicketNotes(file)

	smallesNoteNumber := findSmallest(ticketNotes)
	largestNoteNumber := findLargest(ticketNotes)

	tickets, err := parseTickets(file, smallesNoteNumber, largestNoteNumber)
	if err != nil {
		fmt.Println("Parse ticket error", err)
		return
	}

	printParsedInput(ticketNotes, tickets)

	// Loop through each ticket note, and find matching colums
	for noteIndex := 0; noteIndex < len(ticketNotes); noteIndex++ {
		possibleColumns := findPossibleColumns(ticketNotes[noteIndex], tickets)
		ticketNotes[noteIndex].possibleColumns = possibleColumns
		//	fmt.Printf("Possible columns for %s: %v\n", tNote.field, possibleColumns)
	}

	// Sorting notes by length of possible columns, using sorting fuction above
	sort.Sort(ticketInfoList(ticketNotes))

	var usedColumns []int
	answer := 1

	for idx, tNote := range ticketNotes {
		for _, possibleColumn := range tNote.possibleColumns {
			if !isUsed(possibleColumn, usedColumns) {
				ticketNotes[idx].column = possibleColumn
				usedColumns = append(usedColumns, possibleColumn)
			}
		}

		fmt.Printf("%s: Column = %d. %d possible columns: %v\n", tNote.field, ticketNotes[idx].column, len(tNote.possibleColumns), tNote.possibleColumns)

		// If this is a "Departure *" columns, multiply the number in my ticket
		if len(tNote.field) > 10 && tNote.field[0:10] == "departure " {
			answer *= tickets[0][ticketNotes[idx].column]
		}
	}

	fmt.Printf("\nFound solution: Answer = %d \n", answer)
}

func isUsed(possibleColumn int, usedColumns []int) bool {
	for _, thisCol := range usedColumns {
		if thisCol == possibleColumn {
			return true
		}
	}
	return false
}

func findPossibleColumns(departureLocation ticketInfo, tickets []ticket) []int {
	var columnMatches []int
	for column := 0; column < len(tickets[0]); column++ {
		if allNumbersInColumnFit(column, departureLocation, tickets) {
			columnMatches = append(columnMatches, column)
		}
	}

	return columnMatches
}

func allNumbersInColumnFit(column int, departureLocation ticketInfo, tickets []ticket) bool {
	for i := 0; i < len(tickets); i++ {
		numberToCheck := tickets[i][column]
		if numberToCheck < departureLocation.start1 ||
			(numberToCheck > departureLocation.end1 && numberToCheck < departureLocation.start2) ||
			numberToCheck > departureLocation.end2 {
			return false
		}
	}
	return true
}

func isTicketValid(t ticket, smallesNoteNumber int, largestNoteNumber int) bool {
	for _, num := range t {
		if num < smallesNoteNumber || num > largestNoteNumber {
			return false
		}
	}
	return true
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

func printParsedInput(ticketNotes ticketInfoList, tickets []ticket) {
	for _, note := range ticketNotes {
		fmt.Printf("%s: %d-%d | %d-%d\n", note.field, note.start1, note.end1, note.start2, note.end2)
	}
	fmt.Println()
	for ticketIdx, ticket := range tickets {
		fmt.Printf("%d: ", ticketIdx)
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

func parseTickets(fileLines []string, smallesNoteNumber int, largestNoteNumber int) ([]ticket, error) {

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

		if isTicketValid(thisTicket, smallesNoteNumber, largestNoteNumber) {
			allTickets = append(allTickets, thisTicket)
		}
	}

	return allTickets, nil
}
