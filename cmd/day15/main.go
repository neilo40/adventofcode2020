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
	//part1()
	part2()
}

func part1() {
	// input is 1,12,0,20,8,16 (16 is handled below)
	var input = []int64{0, 1, 12, 0, 20, 8}
	game := make(map[int64]*Number)
	for i, n := range input {
		game[n] = &Number{n, int64(i + 1), int64(i + 1)}
	}

	currentNumber := int64(16)
	var newCurrentNumber int64

	for i := 7; i < 2021; i++ {
		_, ok := game[currentNumber]
		if !ok {
			game[currentNumber] = &Number{currentNumber, int64(i), int64(i)}
			newCurrentNumber = 0
		} else {
			game[currentNumber].LastSpokenPrev = game[currentNumber].LastSpoken
			game[currentNumber].LastSpoken = int64(i)
			newCurrentNumber = game[currentNumber].LastSpoken - game[currentNumber].LastSpokenPrev
		}
		currentNumber = newCurrentNumber
	}

	fmt.Printf("2020th number spoken is %d\n", currentNumber)
}

func part2() {
	// input is 1,12,0,20,8,16 (16 is handled below)
	var input = []int64{0, 1, 12, 0, 20, 8}
	game := make(map[int64]*Number)
	for i, n := range input {
		game[n] = &Number{n, int64(i + 1), int64(i + 1)}
	}

	currentNumber := int64(16)
	var newCurrentNumber int64

	for i := 7; i < 30000001; i++ {
		_, ok := game[currentNumber]
		if !ok {
			game[currentNumber] = &Number{currentNumber, int64(i), int64(i)}
			newCurrentNumber = 0
		} else {
			game[currentNumber].LastSpokenPrev = game[currentNumber].LastSpoken
			game[currentNumber].LastSpoken = int64(i)
			newCurrentNumber = game[currentNumber].LastSpoken - game[currentNumber].LastSpokenPrev
		}
		currentNumber = newCurrentNumber
	}

	fmt.Printf("30000000th number spoken is %d\n", currentNumber)
}
