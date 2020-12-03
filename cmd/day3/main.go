package main

import (
	"fmt"

	"github.com/neilo40/adventofcode2020/internal/common"
)

type slope struct {
	Right int
	Down  int
}

func main() {
	lines := common.ReadFileString("day3.input")
	acc := 1
	slopes := []slope{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	for _, s := range slopes {
		horizPos := 0
		treeCount := 0
		for i := 0; i < len(lines); i += s.Down {
			// actual index into line is horizPos modulo len(line)
			index := horizPos % len(lines[i])
			// check if tree and increment treeCount as necessary
			if lines[i][index] == '#' {
				treeCount++
			}
			// add to horizPos
			horizPos += s.Right
		}
		acc = acc * treeCount
	}
	fmt.Printf("Accumulated result: %d\n", acc)
}
