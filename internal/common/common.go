package common

import (
	"bufio"
	"os"
	"strconv"
)

// ReadFileInt read the given file and returns a slice of integers
func ReadFileInt(filename string) []int64 {
	lines := make([]int64, 0, 1000)
	f, _ := os.Open(filename)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		n, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		lines = append(lines, n)
	}
	return lines
}

// ReadFileString read the given file and returns a slice of integers
func ReadFileString(filename string) []string {
	lines := make([]string, 0, 1000)
	f, _ := os.Open(filename)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
