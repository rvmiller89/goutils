package collection

// Index returns the first index of t in s, or -1 of it does not exist
func Index(s []interface{}, t interface{}) int {
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
