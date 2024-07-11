package data_structures

// approach
// * look for rotations
// * more efficient than rotations
// * take the string and get s + s - then remove first and last character, if it ever repeats then it is a pattern
func repeatedSubstringPattern(s string) bool {
	t := s + s
	// skip first and last matches
	for i := 1; i+len(s) < len(t); i++ {
		if t[i:i+len(s)] == s {
			return true
		}
	}
	return false
}
