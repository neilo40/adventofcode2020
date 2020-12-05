package main

import (
	"testing"
)

func TestFindColumn(t *testing.T) {
	res := findRowColumn(0, 8, "RLR")
	if res != 5 {
		t.Errorf("Expected column 5, got %d\n", res)
	}
}

func TestFindRow(t *testing.T) {
	res := findRowColumn(0, 127, "FBFBBFF")
	if res != 44 {
		t.Errorf("Expected row 44, got %d\n", res)
	}
}

func TestGetSeatId(t *testing.T) {
	r1 := getSeatId("FBFBBFFRLR")
	if r1 != 357 {
		t.Errorf("Expected 357, got %d\n", r1)
	}
}

func TestGetSeatIds(t *testing.T) {
	r1 := getSeatId("BFFFBBFRRR")
	if r1 != 567 {
		t.Errorf("Expected 567, got %d\n", r1)
	}
	r2 := getSeatId("FFFBBBFRRR")
	if r2 != 119 {
		t.Errorf("Expected 119, got %d\n", r2)
	}
	r3 := getSeatId("BBFFBBFRLL")
	if r3 != 820 {
		t.Errorf("Expected 820, got %d\n", r3)
	}
}
