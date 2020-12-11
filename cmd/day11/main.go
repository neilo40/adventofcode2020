package main

import (
	"fmt"

	"github.com/neilo40/adventofcode2020/internal/common"
)

type space struct {
	Type     rune // can be L, ., or #
	NextType rune
	Adjacent map[string]*space
}

type iter func(*[][]space) int

func main() {
	lines := common.ReadFileString("day11.input")

	// build a 2d slice of all spaces (no links)
	layout := make([][]space, len(lines))
	for i, l := range lines {
		layout[i] = make([]space, len(l))
		for j, s := range l {
			layout[i][j].Type = s
			layout[i][j].Adjacent = make(map[string]*space)
		}
	}

	// Now insert the links
	for i, r := range layout {
		for j, s := range r {
			if i > 0 {
				s.Adjacent["N"] = &layout[i-1][j]
			}
			if i > 0 && j < len(r)-1 {
				s.Adjacent["NE"] = &layout[i-1][j+1]
			}
			if j < len(r)-1 {
				s.Adjacent["E"] = &layout[i][j+1]
			}
			if i < len(layout)-1 && j < len(r)-1 {
				s.Adjacent["SE"] = &layout[i+1][j+1]
			}
			if i < len(layout)-1 {
				s.Adjacent["S"] = &layout[i+1][j]
			}
			if i < len(layout)-1 && j > 0 {
				s.Adjacent["SW"] = &layout[i+1][j-1]
			}
			if j > 0 {
				s.Adjacent["W"] = &layout[i][j-1]
			}
			if i > 0 && j > 0 {
				s.Adjacent["NW"] = &layout[i-1][j-1]
			}
		}
	}

	//part1(layout)
	part2(layout)
}

func part1(layout [][]space) {
	iterateAndCount(iterate, layout)
}

func part2(layout [][]space) {
	iterateAndCount(iterateVisible, layout)
}

func iterateAndCount(iterFunc iter, layout [][]space) {
	// iterate until steady state
	for {
		changes := iterFunc(&layout)
		fmt.Printf("%d changes...\n", changes)
		if changes == 0 {
			break
		}
	}

	// count number of seats occupied
	occupied := 0
	for _, r := range layout {
		for _, s := range r {
			if s.Type == '#' {
				occupied++
			}
		}
	}

	fmt.Printf("Occupied seats: %d\n", occupied)
}

func iterate(layout *[][]space) (changes int) {
	for i, r := range *layout {
		for j, s := range r {
			if s.Type == 'L' {
				// if empty (L) and no adjacent seats occupied (#), become occupied (#)
				(*layout)[i][j].NextType = '#'
				for _, a := range s.Adjacent {
					if a.Type == '#' {
						(*layout)[i][j].NextType = 'L'
						break
					}
				}

			} else if s.Type == '#' {
				// if occupied (#) and 4+ adjacent seats occupied (#), become empty (L)
				occupied := 0
				for _, a := range s.Adjacent {
					if a.Type == '#' {
						occupied++
					}
				}
				if occupied > 3 {
					(*layout)[i][j].NextType = 'L'
				}
			}
		}
	}

	// count changes and apply them to the current state
	changes = 0
	for i, r := range *layout {
		for j, s := range r {
			if s.NextType != s.Type {
				changes++
				(*layout)[i][j].Type = s.NextType
			}
		}
	}

	return
}

func iterateVisible(layout *[][]space) (changes int) {
	for i, r := range *layout {
		for j, s := range r {
			if s.Type == 'L' {
				// if empty (L) and no visible seats occupied (#), become occupied (#)
				(*layout)[i][j].NextType = '#'
				for dir := range s.Adjacent {
					if getVisibleSeat(dir, s) == '#' {
						(*layout)[i][j].NextType = 'L'
						break
					}
				}

			} else if s.Type == '#' {
				// if occupied (#) and 5+ visible seats occupied (#), become empty (L)
				occupied := 0
				for dir := range s.Adjacent {
					if getVisibleSeat(dir, s) == '#' {
						occupied++
					}
				}
				if occupied > 4 {
					(*layout)[i][j].NextType = 'L'
				}
			}
		}
	}

	// count changes and apply them to the current state
	changes = 0
	for i, r := range *layout {
		for j, s := range r {
			if s.NextType != s.Type {
				changes++
				(*layout)[i][j].Type = s.NextType
			}
		}
	}

	return
}

func getVisibleSeat(direction string, startingSeat space) rune {
	currentSeat := startingSeat
	for {
		seat, ok := currentSeat.Adjacent[direction]
		if ok {
			if seat.Type != '#' && seat.Type != 'L' {
				currentSeat = *seat
			} else {
				return seat.Type
			}
		} else {
			return '.'
		}
	}
}
