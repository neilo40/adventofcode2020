package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	part1()
	//part2()
}

func readFileInt(filename string) []int {
	lines := make([]int, 0, 1000)
	f, _ := os.Open(filename)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		n, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		lines = append(lines, int(n))
	}
	return lines
}

func part1() {
	lines := readFileInt("day10.input")
	sort.Ints(lines)
	var diffs = map[int]int{1: 0, 2: 0, 3: 1}
	joltage := 0
	for _, j := range lines {
		d := j - joltage
		diffs[d]++
		joltage = j
	}

	fmt.Printf("%v", diffs)
	fmt.Printf("Result: %d\n", diffs[1]*diffs[3])
}

func part2() {
	//lines := common.ReadFileString("dayX.input")
	fmt.Printf("Result: \n")
}
