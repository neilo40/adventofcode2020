package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/neilo40/adventofcode2020/internal/common"
)

type instruction struct {
	Opcode string
	Amount int64
}

func main() {
	//part1()
	part2()
}

func part1() {
	lines := common.ReadFileString("day8.input")
	program := make([]instruction, 0, len(lines))
	for _, l := range lines {
		fields := strings.Fields(l)
		amount, _ := strconv.ParseInt(fields[1], 10, 64)
		program = append(program, instruction{fields[0], amount})
	}

	executed := make([]bool, len(program))
	var acc int64 = 0
	var pc int64 = 0
	for {
		if executed[pc] {
			break // we already executed this location, end program
		}
		executed[pc] = true

		if program[pc].Opcode == "acc" {
			acc += program[pc].Amount
			pc++
		} else if program[pc].Opcode == "jmp" {
			pc += program[pc].Amount
		} else {
			pc++
		}
	}

	fmt.Printf("Acc: %d\n", acc)
}

func part2() {
	lines := common.ReadFileString("day8.input")
	program := make([]instruction, 0, len(lines))
	nopCount := 0
	jmpCount := 0
	for _, l := range lines {
		fields := strings.Fields(l)
		amount, _ := strconv.ParseInt(fields[1], 10, 64)
		if fields[0] == "nop" {
			nopCount++
		} else if fields[0] == "jmp" {
			jmpCount++
		}
		program = append(program, instruction{fields[0], amount})
	}

	changeNop := 0
	changeJmp := 0
	for {
		acc, success := execute(changeNop, changeJmp, program)
		if success {
			fmt.Printf("Acc: %d\n", acc)
			break
		}
		if changeNop == nopCount-1 {
			changeJmp++
		} else {
			changeNop++
		}
	}

}

func execute(changeNop int, changeJmp int, program []instruction) (acc int64, success bool) {
	executed := make([]bool, len(program))
	acc = 0
	var pc int64 = 0
	curNop := 0
	curJmp := 0
	for {
		if pc == int64(len(program)) {
			return acc, true
		}

		if executed[pc] {
			return acc, false
		}
		executed[pc] = true

		if program[pc].Opcode == "acc" {
			acc += program[pc].Amount
			pc++
		} else if program[pc].Opcode == "jmp" {
			if curJmp == changeJmp {
				pc++ // treat this as a nop
			} else {
				pc += program[pc].Amount
			}
			curJmp++
		} else {
			if curNop == changeNop {
				pc += program[pc].Amount // treat this as a jmp
			} else {
				pc++
			}
			curNop++
		}
	}
}
