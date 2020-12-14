package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/neilo40/adventofcode2020/internal/common"
)

func main() {
	part1()
	//part2()
}

func part1() {
	lines := common.ReadFileString("day14.input")
	memory := make(map[int64]int64)
	var setMask int64 = 0
	var resetMask int64 = 0
	memRegex := regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)
	for _, l := range lines {
		if strings.HasPrefix(l, "mask") {
			setMask, resetMask = createMasks(l)
		} else {
			m := memRegex.FindStringSubmatch(l)
			address, _ := strconv.ParseInt(m[1], 10, 64)
			value, _ := strconv.ParseInt(m[2], 10, 64)
			memory[address] = maskedWrite(value, setMask, resetMask)
		}
	}

	var sum int64 = 0
	for _, v := range memory {
		sum += v
	}
	fmt.Printf("Result: %d\n", sum)
}

func createMasks(line string) (setMask int64, resetMask int64) {
	mask := strings.Fields(line)[2]
	for i, c := range mask {
		bitPos := 35 - i
		switch c {
		case '1':
			setMask += int64(math.Pow(2, float64(bitPos)))
			resetMask += int64(math.Pow(2, float64(bitPos)))
		case 'X':
			resetMask += int64(math.Pow(2, float64(bitPos)))
		default:
		}
	}
	return
}

func maskedWrite(value int64, setMask int64, resetMask int64) (maskedVal int64) {
	maskedVal = value | setMask
	maskedVal = maskedVal & resetMask
	return
}

func part2() {
	//lines := common.ReadFileString("dayX.input")
	fmt.Printf("Result: \n")
}
