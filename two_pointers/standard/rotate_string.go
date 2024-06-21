// clarifying questions
// * can have repeated characters? a appears twice?
// * can have different lengths?

// approaches
// * could just manually create each string and try but that would involve O(n^2) comparisons as you consider n starting positions and compare n length strings
// * instead could loop through s until get to starting letter - and then try incrementing from there...
package data_structures

func rotateString(s string, goal string) bool {
	if len(goal) != len(s) {
		return false
	}
	// try to find starting point
	for i, c := range s {
		// found candidate starting point
		if c == rune(goal[0]) {
			j := 0
			for ; j < len(s); j++ {
				m := (i + j) % len(s)
				if s[m] != goal[j] {
					break
				}
			}
			if j == len(s) {
				return true
			}
		}
	}
	return false
}
