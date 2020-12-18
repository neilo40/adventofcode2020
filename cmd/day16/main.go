package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/neilo40/adventofcode2020/internal/common"
)

//Range holds the inclusive upper and lower bounds
type Range struct {
	LowerBound int64
	UpperBound int64
}

//Rule represents a ticket field rule
type Rule struct {
	Name   string
	Bounds []Range
}

//InitFromString initializes a Rule from a text string
func (r *Rule) InitFromString(s string) {
	re := regexp.MustCompile(`(.*): (.*) or (.*)`)
	matches := re.FindStringSubmatch(s)
	r.Name = matches[1]
	r.Bounds = make([]Range, 2)
	for i := 0; i < 2; i++ {
		bounds := strings.Split(matches[i+2], "-")
		lower, _ := strconv.ParseInt(bounds[0], 10, 64)
		upper, _ := strconv.ParseInt(bounds[1], 10, 64)
		r.Bounds[i] = Range{lower, upper}
	}
}

//IsValidFor returns true if the given value is valid for this rule
func (r *Rule) IsValidFor(v int64) bool {
	lowerBoundInvalid := v < r.Bounds[0].LowerBound || v > r.Bounds[0].UpperBound
	upperBoundInvalid := v < r.Bounds[1].LowerBound || v > r.Bounds[1].UpperBound
	if lowerBoundInvalid && upperBoundInvalid {
		return false
	}
	return true
}

func main() {
	lines := common.ReadFileString("day16.input")
	rules := make([]Rule, 0)
	ticketValues := make([][]int64, 0, 300)
	myTicket := make([]int64, 0, 25)

	parsing := 0
	for _, l := range lines {
		if strings.HasPrefix(l, "your ticket") {
			parsing = 1
			continue
		} else if strings.HasPrefix(l, "nearby tickets") {
			parsing = 2
			continue
		} else if l == "" {
			continue
		}

		if parsing == 0 {
			var rule Rule
			rule.InitFromString(l)
			rules = append(rules, rule)
		} else if parsing == 1 {
			vals := strings.Split(l, ",")
			for _, v := range vals {
				vint, _ := strconv.ParseInt(v, 10, 64)
				myTicket = append(myTicket, vint)
			}
		} else {
			vals := strings.Split(l, ",")
			thisTicket := make([]int64, len(vals))
			for i, v := range vals {
				vint, _ := strconv.ParseInt(v, 10, 64)
				thisTicket[i] = vint
			}
			ticketValues = append(ticketValues, thisTicket)
		}
	}

	validValues := make(map[int64]bool)
	for _, rule := range rules {
		for _, r := range rule.Bounds {
			for i := r.LowerBound; i <= r.UpperBound; i++ {
				validValues[i] = true
			}
		}
	}

	var invalidSum int64 = 0
	validTickets := make([][]int64, 0, 300)
	for _, t := range ticketValues {
		invalid := false
		for _, v := range t {
			_, ok := validValues[v]
			if !ok {
				invalidSum += v
				invalid = true
			}
		}
		if !invalid {
			validTickets = append(validTickets, t)
		}
	}

	fmt.Printf("Ticket Scanning Error Rate: %d\n", invalidSum)

	validTickets = append(validTickets, myTicket)
	orderedRules := make([][]Rule, len(rules))
	for i := 0; i < len(validTickets[0]); i++ {
		for _, rule := range rules {
			ruleIsValid := true
			for _, ticket := range validTickets {
				if !rule.IsValidFor(ticket[i]) {
					ruleIsValid = false
					break
				}
			}
			if ruleIsValid {
				orderedRules[i] = append(orderedRules[i], rule)
			}
		}
	}

	for {
		rulesRemoved := 0
		for i, rules := range orderedRules {
			if len(rules) == 1 {
				for j, r := range orderedRules {
					if i == j {
						continue
					}
					for k, rule := range r {
						if rule.Name == rules[0].Name {
							orderedRules[j][k] = orderedRules[j][len(orderedRules[j])-1]
							orderedRules[j] = orderedRules[j][:len(orderedRules[j])-1]
							rulesRemoved++
						}
					}
				}
			}
		}
		if rulesRemoved == 0 {
			break
		}
	}

	var departureProduct int64 = 1
	for i, rule := range orderedRules {
		if strings.HasPrefix(rule[0].Name, "departure") {
			departureProduct *= myTicket[i]
		}
	}

	fmt.Printf("Departure Product: %d\n", departureProduct)
}
