package main

import (
	"fmt"
	"math"

	"github.com/neilo40/adventofcode2020/internal/common"
)

func main() {
	target, lines := part1()
	part2(target, lines)
}

func part1() (int64, []int64) {
	lines := common.ReadFileInt("day9.input")
	for i := 25; i < len(lines); i++ {
		if !isValid(lines[i-25:i], lines[i]) {
			fmt.Printf("Result: %d\n", lines[i])
			return lines[i], lines
		}
	}
	return 0, nil
}

func isValid(numbers []int64, target int64) bool {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if numbers[i] == numbers[j] {
				continue
			}
			if numbers[i]+numbers[j] == target {
				return true
			}
		}
	}
	return false
}

func part2(target int64, lines []int64) {
	for i := 0; i < len(lines); i++ {
		var sum int64 = lines[i]
		for j := 0; j < len(lines); j++ {
			if j <= i {
				continue
			}
			sum += lines[j]
			if sum == target {
				fmt.Printf("Weakness: %d\n", min(lines[i:j+1])+max(lines[i:j+1]))
				return
			}
		}
	}

}

func min(numbers []int64) int64 {
	var min int64 = math.MaxInt64
	for i := 0; i < len(numbers); i++ {
		if numbers[i] < min {
			min = numbers[i]
		}
	}
	return min
}

func max(numbers []int64) int64 {
	var max int64 = 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		}
	}
	return max
}
