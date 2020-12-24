package main

import (
	"fmt"

	"github.com/neilo40/adventofcode2020/internal/common"
)

const (
	gridSize = 201
)

// Hex is a hex
type Hex struct {
	Black      bool
	NextBlack  bool
	Neighbours map[string]*Hex // keys are e, w, se, sw, ne, nw
}

func main() {
	// create a floor with Hexes connected to each other
	floor := make([][]*Hex, gridSize)
	// Create Hexes
	for i := 0; i < gridSize; i++ {
		floor[i] = make([]*Hex, gridSize)
		for j := 0; j < gridSize; j++ {
			floor[i][j] = &Hex{false, false, make(map[string]*Hex)}
		}
	}
	//Connect them
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			// w - every row, not first column
			if j > 0 {
				floor[i][j].Neighbours["w"] = floor[i][j-1]
			}
			// e - every row, not last column
			if j < gridSize-1 {
				floor[i][j].Neighbours["e"] = floor[i][j+1]
			}
			// ne - not first row, not last column on odd rows
			// nw - not first row, not first column on even rows
			if i > 0 {
				if i%2 == 1 {
					// odd rows
					if j < gridSize-1 {
						floor[i][j].Neighbours["ne"] = floor[i-1][j+1]
					}
					floor[i][j].Neighbours["nw"] = floor[i-1][j]
				} else {
					// even rows
					floor[i][j].Neighbours["ne"] = floor[i-1][j]
					if j > 0 {
						floor[i][j].Neighbours["nw"] = floor[i-1][j-1]
					}
				}
			}
			// se - not last row, not last column on odd rows
			// sw - not last row, not first column on even rows
			if i < gridSize-1 {
				if i%2 == 1 {
					// odd rows
					if j < gridSize-1 {
						floor[i][j].Neighbours["se"] = floor[i+1][j+1]
					}
					floor[i][j].Neighbours["sw"] = floor[i+1][j]
				} else {
					// even rows
					floor[i][j].Neighbours["se"] = floor[i+1][j]
					if j > 0 {
						floor[i][j].Neighbours["sw"] = floor[i+1][j-1]
					}
				}
			}
		}
	}

	lines := common.ReadFileString("day24.input")
	for _, l := range lines {
		i := 0
		currentHex := floor[100][100]
		for {
			if l[i] == 'e' || l[i] == 'w' {
				currentHex = currentHex.Neighbours[string(l[i])]
				i++
			} else {
				currentHex = currentHex.Neighbours[l[i:i+2]]
				i += 2
			}
			if i >= len(l) {
				break
			}
		}
		currentHex.Black = !currentHex.Black
	}

	part1(floor)
	part2(floor)
}

func part1(floor [][]*Hex) {
	fmt.Printf("Part 1 - Black tiles: %d\n", countBlack(floor))
}

func countBlack(floor [][]*Hex) int {
	blackCount := 0
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if floor[i][j].Black {
				blackCount++
			}
		}
	}
	return blackCount
}

func part2(floor [][]*Hex) {
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			floor[i][j].NextBlack = floor[i][j].Black
		}
	}

	for d := 0; d < 100; d++ {
		for i := 0; i < gridSize; i++ {
			for j := 0; j < gridSize; j++ {
				blackCount := 0
				for _, n := range floor[i][j].Neighbours {
					if n.Black {
						blackCount++
					}
				}
				if floor[i][j].Black {
					if blackCount == 0 || blackCount > 2 {
						floor[i][j].NextBlack = false
					}
				} else {
					if blackCount == 2 {
						floor[i][j].NextBlack = true
					}
				}
			}
		}

		for i := 0; i < gridSize; i++ {
			for j := 0; j < gridSize; j++ {
				floor[i][j].Black = floor[i][j].NextBlack
			}
		}
	}

	fmt.Printf("Part 2 - Black tiles: %d\n", countBlack(floor))
}
