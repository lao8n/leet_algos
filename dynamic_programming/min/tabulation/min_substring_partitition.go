package data_structures

// clarifying questions
// * not trying to find all possible ways - just trying to find the minimum

// approaches
// * minimum with interconnected problems feels like dynamic programming - can either do top down recursion or bottom up tabulation

// specifics
// * want a table for start to finish - what is minimum number
// * what is the recurrent formula? if it is a new character then easy
// * if given a substring s -> calculate counts for each number -> either it is balanced or need to partition somewhere?
// * can we prune at all? - we know worst case is we partition each section into character?
func minimumSubstringsInPartition(s string) int {
	// set up memoization table
	n := len(s)
	mem := make([]int, n+1)

	for i := 0; i < n; i++ { // i is end of string inclusive
		counts := make(map[byte]int)
		mem[i+1] = mem[i] + 1
		for j := i; j >= 0; j-- { // up to j inclusive
			counts[s[j]]++
			if isBalanced(counts) { // up to i inclusive is i + 1 exclusive
				if mem[j]+1 < mem[i+1] {
					mem[i+1] = mem[j] + 1
				}
			}
		}
	}
	return mem[n]
}

func isBalanced(counts map[byte]int) bool {
	if len(counts) == 0 {
		return false
	}
	expectedCount := 0
	for _, count := range counts {
		if expectedCount == 0 {
			expectedCount = count
		}
		if count != expectedCount {
			return false
		}
	}
	return true
}
