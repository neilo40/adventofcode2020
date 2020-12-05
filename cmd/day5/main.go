package main

import (
	"fmt"

	"github.com/neilo40/adventofcode2020/internal/common"
)

func main() {
	lines := common.ReadFileString("day5.input")
	filledSeats := make([]bool, ((127 * 8) + 7))
	for _, l := range lines {
		seat := getSeatId(l)
		filledSeats[seat] = true
	}
	for i := 1; i < len(filledSeats)-1; i++ {
		if !filledSeats[i] && filledSeats[i-1] && filledSeats[i+1] {
			fmt.Printf("My seat is %d\n", i)
		}
	}
}

func getSeatId(locator string) int {
	row := findRowColumn(0, 127, locator[0:7])
	column := findRowColumn(0, 7, locator[7:])
	return (row * 8) + column
}

func findRowColumn(min int, max int, locator string) int {
	//fmt.Printf("Min: %d, Max: %d, Locator: %s\n", min, max, locator)
	if locator == "L" || locator == "F" {
		return min
	} else if locator == "R" || locator == "B" {
		return max
	} else {
		if locator[0] == 'L' || locator[0] == 'F' {
			newMax := min + (((max - 1) - min) / 2)
			return findRowColumn(min, newMax, locator[1:])
		} else {
			newMin := min + (((max + 1) - min) / 2)
			return findRowColumn(newMin, max, locator[1:])
		}
	}
}
