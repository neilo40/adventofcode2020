package main

import (
	"fmt"
)

//Number represents the spoken numbers in the game
type Number struct {
	Digit          int64
	LastSpoken     int64
	LastSpokenPrev int64
}

func main() {
	part1()
	//part2()
}

func part1() {
	var input = []int64{0, 1, 12, 0, 20, 8, 16}
	game := make(map[int64]Number)
	for i, n := range input {
		game[n] = Number{n, int64(i)}
	}
	currentNumber := input[6]
	for i := 7; i < 2021; i++ {
		n, ok := game[currentNumber]
		if !ok {
			currentNumber = 0
		} else {

		}
	}
	fmt.Printf("2020th number spoken is %d\n", currentNumber)
}

func part2() {
	//lines := common.ReadFileString("dayX.input")
	fmt.Printf("Result: \n")
}
