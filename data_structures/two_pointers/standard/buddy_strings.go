package data_structures

// single swap? yes
// same index? no
// assume same length?

// 1. naive approach - for every c in s try swapping with every c in goal for O(n^2)
// 2. anagram approach? can identify where the same characters a requirement, but doesn't help with swapping indices
// 3. loop through together - find first instance not the same, second instance not the same
// see if characters are different. - doesn't work for repretition case - handle separately?
func buddyStrings(s string, goal string) bool {
	// base case
	if len(s) != len(goal) {
		return false
	}
	// 2 scenarios
	outOfPlace := []int{}
	counts := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] { // e.g. a vs b
			outOfPlace = append(outOfPlace, i)
		} else {
			counts[s[i]] += 1
		}
	}
	// 1. find only two indices out of place
	if len(outOfPlace) == 2 {
		i, j := outOfPlace[0], outOfPlace[1]
		if s[i] == goal[j] && s[j] == goal[i] {
			return true
		}
	}

	// 2. find repetition
	if len(outOfPlace) == 0 {
		for _, count := range counts {
			if count >= 2 {
				return true
			}
		}
	}
	return false
}
