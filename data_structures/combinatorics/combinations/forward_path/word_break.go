package data_structures

// approaches
// * dynamic programming - could either use tabulation to find segments and then create slices - but might be easier with recursion to manage all paths

// specifics
// * dp with recursion - accumulate reduced remaining string forward every time reach end of string then return as a solution
// * memoize? but want all options?
func wordBreak(s string, wordDict []string) []string {
	// setup word dict
	wd := make(map[string]struct{}, len(wordDict))
	for _, word := range wordDict {
		wd[word] = struct{}{}
	}

	// recursion
	return recursive(s, "", wd)
}

func recursive(s string, path string, wordDict map[string]struct{}) []string {
	// base case
	if len(s) == 0 {
		return []string{path}
	}
	// recursive case
	solution := []string{}
	for i := 1; i <= len(s); i++ { // exclusive end
		if _, ok := wordDict[s[0:i]]; ok {
			newPath := s[0:i]
			if path != "" {
				newPath = path + " " + newPath
			}
			recursiveSolution := recursive(s[i:], newPath, wordDict)
			solution = append(solution, recursiveSolution...)
		}
	}
	return solution
}
