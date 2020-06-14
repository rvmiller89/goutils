package collection

// Index returns the first index of t in s, or -1 of it does not exist
func Index(s []string, t string) int {
	for i, v := range s {
		if v == t {
			return i
		}
	}
	return -1
}

// Exists returns true if at least one element satisfies the given predicate function
func Exists(s []string, p func(string) bool) bool {
	for _, v := range s {
		if p(v) {
			return true
		}
	}
	return false
}
