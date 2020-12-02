package main

import (
	"fmt"

	"github.com/neilo40/adventofcode2020/internal/common"
)

func main() {
	lines := common.ReadFileInt("day1.input")
	for i := range lines {
		for j := range lines[i+1:] {
			for k := range lines[i+2:] {
				sum := lines[i] + lines[j] + lines[k]
				if sum == 2020 {
					fmt.Printf("Result is %d\n", lines[i]*lines[j]*lines[k])
					return
				}
			}
		}
	}
}
