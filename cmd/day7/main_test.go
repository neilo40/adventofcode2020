package main

import (
	"testing"
)

func TestCountBags(t *testing.T) {
	mapping := map[string][]requiredContent{
		"vibrant plum": {{"faded blue", 5}, {"dotted black", 6}},
		"dark olive":   {{"faded blue", 3}, {"dotted black", 4}},
		"shiny gold":   {{"dark olive", 1}, {"vibrant plum", 2}},
	}
	res := countBags(1, mapping, "shiny gold")
	if res != 32 {
		t.Errorf("Expected 32 bags, got %d\n", res)
	}
}

func TestCountBags2(t *testing.T) {
	mapping := map[string][]requiredContent{
		"dark red":    {{"dark orange", 2}},
		"dark orange": {{"dark yellow", 2}},
		"dark yellow": {{"dark green", 2}},
		"dark green":  {{"dark blue", 2}},
		"dark blue":   {{"dark violet", 2}},
		"shiny gold":  {{"dark red", 2}},
	}
	res := countBags(1, mapping, "shiny gold")
	if res != 126 {
		t.Errorf("Expected 126 bags, got %d\n", res)
	}
}
