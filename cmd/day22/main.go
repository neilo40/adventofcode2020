package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/neilo40/adventofcode2020/internal/common"
)

func main() {
	part1()
	//part2()
}

func part1() {
	lines := common.ReadFileString("day22.input")
	deck1 := make([]int64, 0, len(lines))
	deck2 := make([]int64, 0, len(lines))
	deck := &deck1
	for _, l := range lines {
		if strings.HasPrefix(l, "Player") {
			if l == "Player 2:" {
				deck = &deck2
			}
			continue
		}
		v, _ := strconv.ParseInt(l, 10, 64)
		*deck = append(*deck, v)
	}

	var winningDeck *[]int64
	var newDeck1 []int64
	var newDeck2 []int64
	for {
		if len(deck1) == 0 {
			winningDeck = &deck2
			break
		}
		if len(deck2) == 0 {
			winningDeck = &deck1
			break
		}
		if deck1[0] > deck2[0] {
			newDeck1 = append(deck1[1:], deck1[0], deck2[0])
			newDeck2 = deck2[1:]
		} else {
			newDeck1 = deck1[1:]
			newDeck2 = append(deck2[1:], deck2[0], deck1[0])
		}
		deck1 = newDeck1
		deck2 = newDeck2
	}

	score := 0
	for i, v := range *winningDeck {
		score += (len(*winningDeck) - i) * int(v)
	}

	fmt.Printf("Winning score: %d\n", score)
}

func part2() {
	//lines := common.ReadFileString("dayX.input")
	fmt.Printf("Result: \n")
}
