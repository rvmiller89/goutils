package collection

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func initEmpty() []interface{} {
	return make([]interface{}, 0)
}

func initCollection() []interface{} {
	var c1 = []string{"apple", "banana", "apple", "cat", "dog"}
	var c2 = make([]interface{}, len(c1))
	for i, c := range c1 {
		c2[i] = c
	}
	return c2
}
func TestIndexOf(t *testing.T) {
	empty := initEmpty()
	col := initCollection()
	tests := []struct {
		col      []interface{}
		target   string
		expected int
	}{
		{empty, "banana", -1},
		{col, "banana", 1},
		{col, "apple", 0},
		{col, "missing", -1},
		{col, "", -1},
		{col, "applephone", -1},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%v,%v,%d", test.col, test.target, test.expected)
		t.Run(testname, func(t *testing.T) {
			ans := IndexOf(test.col, test.target)
			if ans != test.expected {
				t.Errorf("Wanted %d, Got %d", test.expected, ans)
			}
		})
	}
}

func TestExists(t *testing.T) {
	empty := initEmpty()
	col := initCollection()
	tests := []struct {
		col      []interface{}
		p        func(interface{}) bool
		expected bool
	}{
		{empty, func(s interface{}) bool { return s == "kitten" }, false},
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
	empty := initEmpty()
	col := initCollection()
	tests := []struct {
		col      []interface{}
		p        func(interface{}) bool
		expected bool
	}{
		{empty, func(s interface{}) bool { return s == "missing" }, true},
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

func TestFilter(t *testing.T) {
	empty := initEmpty()
	col := initCollection()
	tests := []struct {
		col      []interface{}
		p        func(interface{}) bool
		expected []interface{}
	}{
		{col, func(s interface{}) bool { return s == "cat" }, append(empty, "cat")},
		{empty, func(s interface{}) bool { return s == "cat" }, empty},
		{col, func(s interface{}) bool { s2, _ := s.(string); return strings.Contains(s2, "a") },
			append(empty, "apple", "banana", "apple", "cat")},
		{col, func(s interface{}) bool { s2, _ := s.(string); return len(s2) == 0 }, empty},
	}

	for i, test := range tests {
		testname := fmt.Sprintf("Test %d", i)
		t.Run(testname, func(t *testing.T) {
			ans := Filter(test.col, test.p)
			if !reflect.DeepEqual(ans, test.expected) {
				t.Errorf("Wanted %v, Got %v", test.expected, ans)
			}
		})
	}
}

func TestMap(t *testing.T) {
	empty := initEmpty()
	col := initCollection()
	tests := []struct{
		col []interface{}
		f func(interface{}) interface{}
		expected []interface{}
	}{
		{col, func(s interface{}) interface{} { s2, _ := s.(string); return strings.ToUpper(s2) }, 
		append(empty, "APPLE", "BANANA", "APPLE", "CAT", "DOG")},
		{empty, func(s interface{}) interface{} { s2, _ := s.(string); return strings.ToUpper(s2) }, empty},
		{col, func(s interface{}) interface{} { s2, _ := s.(string); return len(s2) }, append(empty, 5, 6, 5, 3, 3)},
	}

	for i, test := range tests {
		testname := fmt.Sprintf("Test %d", i)
		t.Run(testname, func(t *testing.T) {
			ans := Map(test.col, test.f)
			if !reflect.DeepEqual(ans, test.expected) {
				t.Errorf("Wanted %v, Got %v", test.expected, ans)
			}
		})
	}
}