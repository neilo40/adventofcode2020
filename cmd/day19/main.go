package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/neilo40/adventofcode2020/internal/common"
)

// Rule is a rule
type Rule struct {
	Number      int64
	SubRules    [][]int64
	MatchString string
}

// GetMatchStrings recursively builds the match strings for a given rule
func (r *Rule) GetMatchStrings(rules map[int64]*Rule) []string {
	if r.MatchString != "" {
		return []string{r.MatchString}
	}
	matchStrings := []string{}
	for _, subRule := range r.SubRules {
		if len(subRule) > 1 {
			for _, leftString := range rules[subRule[0]].GetMatchStrings(rules) {
				for _, rightString := range rules[subRule[1]].GetMatchStrings(rules) {
					matchStrings = append(matchStrings, leftString+rightString)
				}
			}
		} else {
			matchStrings = append(matchStrings, rules[subRule[0]].GetMatchStrings(rules)...)
		}
	}
	return matchStrings
}

func main() {
	part1()
	//part2()
}

func part1() {
	lines := common.ReadFileString("day19.input")
	rules := make(map[int64]*Rule)
	messages := make([]string, 0, len(lines))
	parsingRules := true
	for _, l := range lines {
		if l == "" {
			parsingRules = false
			continue
		}

		if parsingRules {
			parts := strings.Split(l, ":")
			ruleNum, _ := strconv.ParseInt(parts[0], 10, 64)
			subRules := make([][]int64, 0, 2)
			matchString := ""
			if strings.Contains(parts[1], "a") {
				matchString = "a"
			} else if strings.Contains(parts[1], "b") {
				matchString = "b"
			} else {
				subRulesPairs := strings.Split(parts[1], "|")
				for _, pair := range subRulesPairs {
					subRule := make([]int64, 0, 2)
					for _, n := range strings.Fields(pair) {
						subRuleNum, _ := strconv.ParseInt(n, 10, 64)
						subRule = append(subRule, subRuleNum)
					}
					subRules = append(subRules, subRule)
				}
			}
			rules[ruleNum] = &Rule{
				Number:      ruleNum,
				SubRules:    subRules,
				MatchString: matchString,
			}
		} else {
			messages = append(messages, l)
		}
	}

	rule0MatchStrings := rules[0].GetMatchStrings(rules)
	matchCount := 0
	for _, message := range messages {
		for _, ruleMatchString := range rule0MatchStrings {
			if message == ruleMatchString {
				matchCount++
				break
			}
		}
	}
	fmt.Printf("Completely matching messages: %d\n", matchCount)
}

func part2() {
	//lines := common.ReadFileString("dayX.input")
	fmt.Printf("Result: \n")
}
