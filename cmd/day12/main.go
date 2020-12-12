package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/neilo40/adventofcode2020/internal/common"
)

type coord struct {
	NorthSouth int64
	EastWest   int64
}

func main() {
	//part1()
	part2()
}

func part1() {
	lines := common.ReadFileString("day12.input")
	curPos := coord{0, 0}
	curDir := 'E'
	for _, l := range lines {
		amount, _ := strconv.ParseInt(l[1:], 10, 64)
		switch l[0] {
		case 'N', 'S', 'E', 'W':
			curPos = move(rune(l[0]), amount, curPos)
		case 'L', 'R':
			curDir = turn(rune(l[0]), amount, curDir)
		case 'F':
			curPos = move(curDir, amount, curPos)
		default:
			panic("Invalid case")
		}
	}

	dist := math.Abs(float64(curPos.NorthSouth)) + math.Abs(float64(curPos.EastWest))
	fmt.Printf("Final coord: %v.  Manhattan distance: %f\n", curPos, dist)
}

func move(dir rune, amount int64, curPos coord) (newPos coord) {
	newPos = curPos
	switch dir {
	case 'N':
		newPos.NorthSouth += amount
	case 'S':
		newPos.NorthSouth -= amount
	case 'E':
		newPos.EastWest += amount
	case 'W':
		newPos.EastWest -= amount
	}
	return
}

func turn(dir rune, amount int64, facing rune) (newDir rune) {
	var directions = []rune{'N', 'E', 'S', 'W'}
	var dirIndices = map[rune]int64{'N': 0, 'E': 1, 'S': 2, 'W': 3}
	steps := amount / 90
	if dir == 'L' {
		steps = steps * -1
	}
	curIndex := dirIndices[facing]
	newIndex := curIndex + steps
	if newIndex < 0 {
		newIndex += 4
	} else if newIndex > 3 {
		newIndex -= 4
	}
	newDir = directions[newIndex]
	return
}

func part2() {
	lines := common.ReadFileString("day12.input")
	curPos := coord{0, 0}
	waypoint := coord{1, 10}
	for _, l := range lines {
		amount, _ := strconv.ParseInt(l[1:], 10, 64)
		switch l[0] {
		case 'N', 'S', 'E', 'W':
			waypoint = move(rune(l[0]), amount, waypoint)
		case 'L', 'R':
			waypoint = rotateWaypoint(rune(l[0]), amount, curPos, waypoint)
		case 'F':
			curPos, waypoint = moveToWaypoint(curPos, waypoint, amount)
		default:
			panic("Invalid case")
		}
	}

	dist := math.Abs(float64(curPos.NorthSouth)) + math.Abs(float64(curPos.EastWest))
	fmt.Printf("Final coord: %v.  Manhattan distance: %f\n", curPos, dist)
}

func rotateWaypoint(dir rune, amount int64, curPos coord, waypoint coord) coord {
	waypointVector := coord{
		NorthSouth: waypoint.NorthSouth - curPos.NorthSouth,
		EastWest:   waypoint.EastWest - curPos.EastWest,
	}
	// we will only go right, so convert the amount if dir == 'L'
	if dir == 'L' {
		amount = 360 - amount
	}
	times := amount / 90 // we'll repeat in 90 degree rotations
	for i := 0; i < int(times); i++ {
		waypointVector = coord{
			EastWest:   waypointVector.NorthSouth,
			NorthSouth: waypointVector.EastWest * -1,
		}
	}
	newWaypoint := coord{
		EastWest:   curPos.EastWest + waypointVector.EastWest,
		NorthSouth: curPos.NorthSouth + waypointVector.NorthSouth,
	}
	return newWaypoint
}

func moveToWaypoint(curPos coord, waypoint coord, times int64) (coord, coord) {
	newCurPos := curPos
	newWaypoint := waypoint
	waypointVector := coord{
		NorthSouth: waypoint.NorthSouth - curPos.NorthSouth,
		EastWest:   waypoint.EastWest - curPos.EastWest,
	}
	for i := 0; i < int(times); i++ {
		newCurPos.NorthSouth += waypointVector.NorthSouth
		newCurPos.EastWest += waypointVector.EastWest
		newWaypoint.NorthSouth += waypointVector.NorthSouth
		newWaypoint.EastWest += waypointVector.EastWest
	}
	return newCurPos, newWaypoint
}
