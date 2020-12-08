package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/neilo40/adventofcode2020/internal/common"
)

type requiredContent struct {
	Colour   string
	Quantity int64
}

func main() {
	//part1()
	part2()
}

func part1() {
	mapping := getMapping()
	parents := getAllParents("shiny gold", mapping)
	fmt.Println(countUnique(parents))
}

// gets mapping of childColour -> []parentColour
func getMapping() (mapping map[string][]string) {
	lines := common.ReadFileString("day7.input")
	mapping = make(map[string][]string)
	for _, l := range lines {
		lineParts := strings.Fields(l)
		parentColour := fmt.Sprintf("%s %s", lineParts[0], lineParts[1])
		childParts := lineParts[4:]
		for i := 0; i < len(childParts); i += 4 {
			childColour := fmt.Sprintf("%s %s", childParts[i+1], childParts[i+2])
			_, ok := mapping[childColour]
			if ok {
				mapping[childColour] = append(mapping[childColour], parentColour)
			} else {
				mapping[childColour] = []string{parentColour}
			}
		}
	}
	return
}

func getAllParents(colour string, mapping map[string][]string) []string {
	_, ok := mapping[colour]
	if !ok {
		return []string{} // no parents
	}

	parents := mapping[colour]
	for _, pc := range mapping[colour] {
		parents = append(parents, getAllParents(pc, mapping)...)
	}
	return parents

}

func countUnique(colours []string) int {
	colourMap := make(map[string]bool)
	for _, c := range colours {
		colourMap[c] = true
	}
	return len(colourMap)
}

func part2() {
	mapping := getRequiredContentsMapping()
	//countBags includes the shiny gold bag in the total so we need to subtract 1 from the total
	fmt.Println(countBags(1, mapping, "shiny gold") - 1)
}

// gets mapping of parentColour -> []requiredContent
func getRequiredContentsMapping() (mapping map[string][]requiredContent) {
	lines := common.ReadFileString("day7.input")
	mapping = make(map[string][]requiredContent)
	for _, l := range lines {
		lineParts := strings.Fields(l)
		parentColour := fmt.Sprintf("%s %s", lineParts[0], lineParts[1])
		childParts := lineParts[4:]
		for i := 0; i < len(childParts); i += 4 {
			childColour := fmt.Sprintf("%s %s", childParts[i+1], childParts[i+2])
			qty, _ := strconv.ParseInt(childParts[i], 10, 64)
			required := requiredContent{childColour, qty}
			_, ok := mapping[parentColour]
			if ok {
				mapping[parentColour] = append(mapping[parentColour], required)
			} else {
				mapping[parentColour] = []requiredContent{required}
			}
		}
	}
	return
}

func countBags(qty int64, mapping map[string][]requiredContent, colour string) int64 {
	_, ok := mapping[colour]
	if !ok {
		return qty // no other contents, just return self
	}

	var thisTotal int64 = 0
	for _, r := range mapping[colour] {
		thisTotal += countBags(r.Quantity, mapping, r.Colour)
	}
	return (qty * thisTotal) + qty //sum of all other contents + self
}
