package data_structures

// approach
// * recurse through string from start to finish
// * when reach * have two recursions one is where move to next and the other where stay on star
func isMatch(s string, p string) bool {
	var lastStar bool
	var recursive func(s, p string) bool
	recursive = func(s, p string) bool {
		// base case
		if len(p) == 0 {
			return len(s) == 0
		}
		// recursive case
		matched := false
		if len(s) > 0 && (p[0] == '?' || p[0] == s[0]) {
			matched = recursive(s[1:], p[1:])
		} else if p[0] == '*' {
			matched = len(p) == 1 || recursive(s, p[1:])
			if !lastStar && !matched && len(s) > 0 {
				matched = recursive(s[1:], p)
				lastStar = true
			}
		}
		return matched
	}
	return recursive(s, p)
}
