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
	lines := common.ReadFileString("day18.input")
	var sum int64 = 0
	for _, l := range lines {
		sum += calcExpression(l)
	}
	fmt.Printf("Sum: %d\n", sum)
}

func calcExpression(line string) int64 {
	line = strings.ReplaceAll(line, "(", "( ")
	line = strings.ReplaceAll(line, ")", " )")
	sum, _ := calculate(strings.Fields(line), 0, "+")
	return sum
}

func calculate(expression []string, sum int64, operation string) (int64, []string) {
	// operation, capture and move on to next operand
	if expression[0] == "+" || expression[0] == "*" {
		return calculate(expression[1:], sum, expression[0])
	}

	// next amount - either digit or parens block
	var nextAmount int64 = 0
	if expression[0] == "(" {
		nextAmount, expression = calculate(expression[1:], 0, "+")
	} else if expression[0] == ")" {
		if len(expression) == 1 {
			return sum, []string{}
		}
		return sum, expression
	} else {
		nextAmount, _ = strconv.ParseInt(expression[0], 10, 64)
	}

	// do operation
	if operation == "*" {
		sum *= nextAmount
	} else {
		sum += nextAmount
	}

	// continue parsing unless at end
	if len(expression) > 1 {
		return calculate(expression[1:], sum, operation)
	}
	// end
	return sum, []string{}
}

func part2() {
	//lines := common.ReadFileString("dayX.input")
	fmt.Printf("Result: \n")
}
