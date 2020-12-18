package main

import (
	"fmt"

	"github.com/neilo40/adventofcode2020/internal/common"
)

const (
	zMax       = 13
	xMax       = 20
	yMax       = 20
	iterations = 6
)

//Space is a space
type Space struct {
	Type [2]rune // can be ., or #
}

func main() {
	//part1()
	part2()
}

func part1() {
	lines := common.ReadFileString("day17.input")
	// 13x20x20 space should be enough given initial conditions and 6 iteraton limit
	pocket := [13][20][20]*Space{}
	for i, l := range lines {
		for j, s := range l {
			// This should put the known locatons in the center of the pocket
			pocket[iterations][i+iterations][j+iterations] = &Space{[2]rune{s, '.'}}
		}
	}

	typeIndex := 0
	nextTypeIndex := 1
	for i := 0; i < iterations; i++ {
		for z := 0; z < zMax; z++ {
			for x := 0; x < xMax; x++ {
				for y := 0; y < yMax; y++ {
					iterate([3]int{z, x, y}, &pocket, typeIndex, nextTypeIndex)
				}
			}
		}

		if typeIndex == 0 {
			typeIndex = 1
			nextTypeIndex = 0
		} else {
			typeIndex = 0
			nextTypeIndex = 1
		}
	}

	activeCubes := 0
	for z := 0; z < zMax; z++ {
		for x := 0; x < xMax; x++ {
			for y := 0; y < yMax; y++ {
				if pocket[z][x][y].Type[typeIndex] == '#' {
					activeCubes++
				}
			}
		}
	}

	fmt.Printf("Active cubes: %d\n", activeCubes)
}

func iterate(coord [3]int, pocket *[13][20][20]*Space, typeIndex int, nextTypeIndex int) {
	activeNeighbours := 0
	for z := coord[0] - 1; z <= coord[0]+1; z++ {
		for x := coord[1] - 1; x <= coord[1]+1; x++ {
			for y := coord[2] - 1; y <= coord[2]+1; y++ {
				if z < 0 || z >= zMax || x < 0 || x >= xMax || y < 0 || y >= yMax {
					continue // edge of pocket cases
				}
				if pocket[z][x][y] == nil {
					pocket[z][x][y] = &Space{[2]rune{'.', '.'}}
					continue
				}
				if z == coord[0] && x == coord[1] && y == coord[2] {
					continue // it's us!
				}
				if pocket[z][x][y].Type[typeIndex] == '#' {
					activeNeighbours++
				}
			}
		}
	}
	if pocket[coord[0]][coord[1]][coord[2]].Type[typeIndex] == '.' {
		if activeNeighbours == 3 {
			pocket[coord[0]][coord[1]][coord[2]].Type[nextTypeIndex] = '#'
		} else {
			pocket[coord[0]][coord[1]][coord[2]].Type[nextTypeIndex] = '.'
		}
	} else {
		if activeNeighbours == 2 || activeNeighbours == 3 {
			pocket[coord[0]][coord[1]][coord[2]].Type[nextTypeIndex] = '#'
		} else {
			pocket[coord[0]][coord[1]][coord[2]].Type[nextTypeIndex] = '.'
		}
	}
}

func part2() {
	lines := common.ReadFileString("day17.input")
	// 13x13x20x20 space should be enough given initial conditions and 6 iteraton limit
	pocket := [13][13][20][20]*Space{}
	for i, l := range lines {
		for j, s := range l {
			// This should put the known locatons in the center of the pocket
			pocket[iterations][iterations][i+iterations][j+iterations] = &Space{[2]rune{s, '.'}}
		}
	}

	typeIndex := 0
	nextTypeIndex := 1
	for i := 0; i < iterations; i++ {
		for w := 0; w < zMax; w++ {
			for z := 0; z < zMax; z++ {
				for x := 0; x < xMax; x++ {
					for y := 0; y < yMax; y++ {
						iterate4D([4]int{w, z, x, y}, &pocket, typeIndex, nextTypeIndex)
					}
				}
			}
		}

		if typeIndex == 0 {
			typeIndex = 1
			nextTypeIndex = 0
		} else {
			typeIndex = 0
			nextTypeIndex = 1
		}
	}

	activeCubes := 0
	for w := 0; w < zMax; w++ {
		for z := 0; z < zMax; z++ {
			for x := 0; x < xMax; x++ {
				for y := 0; y < yMax; y++ {
					if pocket[w][z][x][y].Type[typeIndex] == '#' {
						activeCubes++
					}
				}
			}
		}
	}

	fmt.Printf("Active cubes: %d\n", activeCubes)
}

func iterate4D(coord [4]int, pocket *[13][13][20][20]*Space, typeIndex int, nextTypeIndex int) {
	activeNeighbours := 0
	for w := coord[0] - 1; w <= coord[0]+1; w++ {
		for z := coord[1] - 1; z <= coord[1]+1; z++ {
			for x := coord[2] - 1; x <= coord[2]+1; x++ {
				for y := coord[3] - 1; y <= coord[3]+1; y++ {
					if w < 0 || w >= zMax || z < 0 || z >= zMax || x < 0 || x >= xMax || y < 0 || y >= yMax {
						continue // edge of pocket cases
					}
					if pocket[w][z][x][y] == nil {
						pocket[w][z][x][y] = &Space{[2]rune{'.', '.'}}
						continue
					}
					if w == coord[0] && z == coord[1] && x == coord[2] && y == coord[3] {
						continue // it's us!
					}
					if pocket[w][z][x][y].Type[typeIndex] == '#' {
						activeNeighbours++
					}
				}
			}
		}
	}

	if pocket[coord[0]][coord[1]][coord[2]][coord[3]].Type[typeIndex] == '.' {
		if activeNeighbours == 3 {
			pocket[coord[0]][coord[1]][coord[2]][coord[3]].Type[nextTypeIndex] = '#'
		} else {
			pocket[coord[0]][coord[1]][coord[2]][coord[3]].Type[nextTypeIndex] = '.'
		}
	} else {
		if activeNeighbours == 2 || activeNeighbours == 3 {
			pocket[coord[0]][coord[1]][coord[2]][coord[3]].Type[nextTypeIndex] = '#'
		} else {
			pocket[coord[0]][coord[1]][coord[2]][coord[3]].Type[nextTypeIndex] = '.'
		}
	}
}
