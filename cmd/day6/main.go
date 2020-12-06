package main

import (
	"fmt"

	"github.com/neilo40/adventofcode2020/internal/common"
)

func main() {
	part2()
}

func part1() {
	lines := common.ReadFileString("day6.input")
	sum := 0
	answers := make(map[string]bool)
	for _, l := range lines {
		if l == "" {
			sum += len(answers)
			answers = make(map[string]bool)
		} else {
			for _, q := range l {
				answers[string(q)] = true
			}
		}
	}

	// last set don't have a newline after them
	sum += len(answers)

	fmt.Printf("Sum of answers = %d\n", sum)
}

func part2() {
	lines := common.ReadFileString("day6.input")
	sum := 0
	answers := make(map[string]int)
	personCount := 0
	for _, l := range lines {
		if l == "" {
			sum += getAnswerCount(answers, personCount)
			answers = make(map[string]int)
			personCount = 0
		} else {
			personCount += 1
			for _, q := range l {
				_, ok := answers[string(q)]
				if ok {
					answers[string(q)] += 1
				} else {
					answers[string(q)] = 1
				}
			}
		}
	}

	// last set don't have a newline after them
	sum += getAnswerCount(answers, personCount)

	fmt.Printf("Sum of answers = %d\n", sum)
}

func getAnswerCount(answers map[string]int, personCount int) int {
	c := 0
	for _, v := range answers {
		if v == personCount {
			c++
		}
	}
	return c
}
