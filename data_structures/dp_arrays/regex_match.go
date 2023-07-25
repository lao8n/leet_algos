package data_structures

// dynamic programming = don't know greedily how many characters to match * for, however
// do not have to try each matching independently can memoize computation in a grid

func isMatch(s string, p string) bool {
	lenS := len(s)
	lenP := len(p)
	// initialize grid with extra row for setup so we shift everything along
	dp := make([][]bool, lenS+1)
	for rowS := range dp {
		dp[rowS] = make([]bool, lenP+1)
	}

	// base case
	dp[0][0] = true
	// find all potential first elements - set in init row
	for pI := 2; pI <= lenP; pI += 2 {
		if p[pI-1] == '*' {
			dp[0][pI] = true
		} else {
			break
		}
	}

	for sI := 1; sI <= lenS; sI++ {
		for pI := 1; pI <= lenP; pI++ {
			// match (-1 only because we are storing in dp with +1)
			currentMatch := s[sI-1] == p[pI-1] || p[pI-1] == '.'
			currentStar := p[pI-1] == '*'

			if currentMatch {
				previousMatch := dp[sI-1][pI-1]
				dp[sI][pI] = previousMatch
			} else if currentStar { // e.g. ab vs c*
				ignoreStar := dp[sI][pI-2] // ignore c*
				matchOnceStar := dp[sI][pI-1]
				previousMatch := dp[sI-1][pI]                     // a match *?
				starMatch := s[sI-1] == p[pI-2] || p[pI-2] == '.' // b == c?
				dp[sI][pI] = ignoreStar || matchOnceStar || (previousMatch && starMatch)
			}
		}
	}
	return dp[lenS][lenP]
}
