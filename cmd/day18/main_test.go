package main

import (
	"testing"
)

func Test1(t *testing.T) {
	result := calcExpression("1 + 2 * 3 + 4 * 5 + 6")
	if result != 71 {
		t.Errorf("Expected 71, got %d\n", result)
	}
}

func Test2(t *testing.T) {
	result := calcExpression("2 * 3 + (4 * 5)")
	if result != 26 {
		t.Errorf("Expected 26, got %d\n", result)
	}
}

func Test3(t *testing.T) {
	result := calcExpression("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))")
	if result != 12240 {
		t.Errorf("Expected 12240, got %d\n", result)
	}
}

func Test4(t *testing.T) {
	result := calcExpression("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2")
	if result != 13632 {
		t.Errorf("Expected 13632, got %d\n", result)
	}
}
