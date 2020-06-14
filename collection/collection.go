package collection

// IndexOf returns the first index of t in s, or -1 of it does not exist
func IndexOf(s []interface{}, t interface{}) int {
	for i, v := range s {
		if v == t {
			return i
		}
	}
	return -1
}

// Exists returns true if at least one element satisfies the given predicate function
func Exists(s []interface{}, p func(interface{}) bool) bool {
	for _, v := range s {
		if p(v) {
			return true
		}
	}
	return false
}

// Forall returns true if all elements satisfy the given predicate function
func Forall(s []interface{}, p func(interface{}) bool) bool {
	for _, v := range s {
		if !p(v) {
			return false
		}
	}
	return true
}

// Filter returns a slice of elements that satisfy the given predicate function
func Filter(s []interface{}, p func(interface{}) bool) []interface{} {
	match := make([]interface{}, 0)
	for _, v := range s {
		if p(v) {
			match = append(match, v)
		}
	}
	return match
}

// Map returns the slice of elements after applying the function f to each element
func Map(s []interface{}, f func(interface{}) interface{}) []interface{} {
	for i, v := range s {
		s[i] = f(v)
	}
	return s
}
