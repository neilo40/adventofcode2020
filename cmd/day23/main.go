package main

import (
	"fmt"
	"strconv"
)

func main() {
	part1()
	//part2()
}

func part1() {
	state := []int{4, 5, 9, 6, 7, 2, 8, 1, 3}
	//state := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	iterationCount := 100

	for i := 0; i < iterationCount; i++ {
		state = iterate(state)
	}

	resultString := ""
	state2 := append(state, state...)
	for i, v := range state2 {
		if v == 1 {
			for j := i + 1; j < i+9; j++ {
				resultString += strconv.FormatInt(int64(state2[j]), 10)
			}
			break
		}

	}
	fmt.Printf("Result: %s\n", resultString)
}

func iterate(state []int) []int {
	var newState = make([]int, 0, 9)
	current := state[0]
	pickup := state[1:4]
	destination := 0
	preDest := make([]int, 0, 4)
	postDest := make([]int, 0, 4)
	d := current - 1
	for {
		if d == 0 {
			d = 9
		}
		for i, v := range state[4:] {
			if v == d {
				destination = v
				if i > 0 {
					preDest = state[4 : 4+i]
				}
				if i < 5 {
					postDest = state[5+i:]
				}
				break
			}
		}
		if destination > 0 {
			break
		}
		d--
	}

	newState = append(newState, preDest...)
	newState = append(newState, destination)
	newState = append(newState, pickup...)
	newState = append(newState, postDest...)
	newState = append(newState, current)
	return newState
}

func part2() {
	//lines := common.ReadFileString("dayX.input")
	fmt.Printf("Result: \n")
}
