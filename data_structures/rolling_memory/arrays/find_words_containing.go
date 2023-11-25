package data_structures

// clarifying questions
// * doesn't matter repetitions it is just the letters
// approaches
// * no choice but to brute force - can bail out of for loop early if find letter
func findWordsContaining(words []string, x byte) []int {
	solution := make([]int, 0)
	for i, word := range words {
		for _, l := range word {
			if byte(l) == x {
				solution = append(solution, i)
				break // don't need to add a second time
			}
		}
	}
	return solution
}
