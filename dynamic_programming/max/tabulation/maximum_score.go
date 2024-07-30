package data_structures

func maximumScore(grid [][]int) int64 {
	dp := make([][][2]int, len(grid)) // i, lastHeight, inclColScore
	for i := range dp {
		dp[i] = make([][2]int, len(grid)+1)
	}

	for i := 0; i < len(dp)-1; i++ { // i is the column number
		for lastHeight := range dp[0] { // height of last column (i), height is the amount of colored cells
			lastColScore := 0                       // score from column i added by coloring column i+1
			nextColScore := 0                       // score from column i+1 added by coloring column i
			for row := 0; row < lastHeight; row++ { // score from col i+1 calculated according to height of i
				nextColScore += grid[row][i+1]
			}
			for height := range dp[0] { // height of next column (i + 1)
				if height > 0 && height <= lastHeight { // next column doesn't contibute score from the colored cells
					nextColScore -= grid[height-1][i+1]
				}
				if height > lastHeight { // add score from the neighbour of the newly colored cell
					lastColScore += grid[height-1][i]
				}

				exclColScore := 0 // indecies of the 2-cell array
				inclColScore := 1
				dp[i+1][height][exclColScore] = max( // either previously conidered value or a new higher value
					dp[i+1][height][exclColScore],
					dp[i][lastHeight][exclColScore]+lastColScore, // include col score when it was excluded
					dp[i][lastHeight][inclColScore],
				)
				dp[i+1][height][inclColScore] = max(
					dp[i+1][height][inclColScore],
					dp[i][lastHeight][inclColScore]+nextColScore,
					dp[i][lastHeight][exclColScore]+nextColScore+lastColScore,
				)
			}
		}
	}

	res := 0
	i := len(dp) - 1
	for lastHeight := range dp[i] { // find max score
		res = max(res, dp[i][lastHeight][0], dp[i][lastHeight][1])
	}

	return int64(res)
}
