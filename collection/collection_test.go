package collection

import (
	"fmt"
	"testing"
)

var col = []string{"apple", "banana", "apple", "cat", "dog"}

func TestIndex(t *testing.T) {
	tests := []struct {
		col      []string
		target   string
		expected int
	}{
		{col, "banana", 1},
		{col, "apple", 0},
		{col, "missing", -1},
		{col, "", -1},
		{col, "applephone", -1},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%v,%v,%d", test.col, test.target, test.expected)
		t.Run(testname, func(t *testing.T) {
			ans := Index(test.col, test.target)
			if ans != test.expected {
				t.Errorf("Wanted %d, Got %d", test.expected, ans)
			}
		})
	}
}

func TestExists(t *testing.T) {
	tests := []struct {
		col      []string
		p        func(string) bool
		expected bool
	}{
		{col, func(s string) bool { return len(s) < 3 }, false},
		{col, func(s string) bool { return len(s) >= 3 }, true},
		{col, func(s string) bool { return s == "cat" }, true},
	}

	for i, test := range tests {
		testname := fmt.Sprintf("Test %d", i)
		t.Run(testname, func(t *testing.T) {
			ans := Exists(test.col, test.p)
			if ans != test.expected {
				t.Errorf("Wanted %v, Got %v", test.expected, ans)
			}
		})
	}
}
