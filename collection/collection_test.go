package collection

import (
	"fmt"
	"testing"
)

func initCollection() []interface{} {
	var c1 = []string{"apple", "banana", "apple", "cat", "dog"}
	var c2 = make([]interface{}, len(c1))
	for i, c := range c1 {
		c2[i] = c
	}
	return c2
}
func TestIndex(t *testing.T) {
	col := initCollection()
	tests := []struct {
		col      []interface{}
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
	col := initCollection()
	tests := []struct {
		col      []interface{}
		p        func(interface{}) bool
		expected bool
	}{
		{col, func(s interface{}) bool { return s == "kitten" }, false},
		{col, func(s interface{}) bool { return s == "cat" }, true},
		{col, func(s interface{}) bool { s2, _ := s.(string); return len(s2) >= 3 }, true},
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

func TestForall(t *testing.T) {
	col := initCollection()
	tests := []struct {
		col      []interface{}
		p        func(interface{}) bool
		expected bool
	}{
		{col, func(s interface{}) bool { return s == "missing" }, false},
		{col, func(s interface{}) bool { return s == "cat" }, false},
		{col, func(s interface{}) bool { s2, _ := s.(string); return len(s2) >= 3 }, true},
	}

	for i, test := range tests {
		testname := fmt.Sprintf("Test %d", i)
		t.Run(testname, func(t *testing.T) {
			ans := Forall(test.col, test.p)
			if ans != test.expected {
				t.Errorf("Wanted %v, Got %v", test.expected, ans)
			}
		})
	}
}
