package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/neilo40/adventofcode2020/internal/common"
)

func main() {
	lines := common.ReadFileString("day2.input")
	re := regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)
	matchingPasswords := 0
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		var matchCount int64 = 0
		lowerIndex, _ := strconv.ParseInt(matches[1], 10, 64)
		upperIndex, _ := strconv.ParseInt(matches[2], 10, 64)
		if int64(len(matches[4])) >= upperIndex {
			// indexes in input are 1-based so subtract 1
			if string(matches[4][lowerIndex-1]) == matches[3] {
				matchCount++
			}
			if string(matches[4][upperIndex-1]) == matches[3] {
				matchCount++
			}
		}
		if matchCount == 1 {
			matchingPasswords++
		}
	}
	fmt.Printf("Found %d matching passwords\n", matchingPasswords)
}
