package data_structures

import "math"

// these are interconnected subproblems combined for a solution -> dynamic programming
// can either do bottom up tabulation or top down recursion with memoization -> bottom up tabulation

// 3 things need to define
// 1. base cases
// 2. recurrence relation -> more options than standard dp as only constraint is have minimum 1 per day
// 3. state variables

// complexity
// time = O(n^2 d) because there are n x d possible states and it takes O(n) to calculate the result for each state
// space = O(n d) for all the n d states

// improvements
// could reduce size of the array stored to just the previous days calculations

// dynamic programming tip:
// min_diff[d][i] = min(daily_max_job_diff + min_diff[d - 1][j]) for all j > i
// rather than just some for loops
func minDifficulty(jobDifficulty []int, d int) int {
	if len(jobDifficulty) < d {
		return -1
	}
	// setup data structures
	// index of remaining jobs & num of days done (d - 1)
	memoizedCost := make([][]int, len(jobDifficulty))
	for i := range memoizedCost {
		memoizedCost[i] = make([]int, d)
		for j := 0; j < d; j++ {
			memoizedCost[i][j] = -1
		}
	}
	// setup data structures
	memoizedCostPreviousDay := make([]int, len(jobDifficulty))

	// base cases - what we do on first day
	maxCost := 0
	for j := 0; j < len(jobDifficulty)-(d-1); j++ {
		if jobDifficulty[j] > maxCost {
			maxCost = jobDifficulty[j]
		}
		memoizedCostPreviousDay[j] = maxCost
	}

	// tabulated cases - what we do on the dith day
	for di := 1; di < d; di++ {
		memoizedCostCurrentDay := make([]int, len(jobDifficulty))
		for j := di; j <= len(jobDifficulty)-(d-di); j++ {
			// recurrence relation
			memoizedCostCurrentDay[j] = recurrence(
				j,
				di,
				jobDifficulty,
				memoizedCostPreviousDay,
			)
		}
		memoizedCostPreviousDay = memoizedCostCurrentDay
	}
	return memoizedCostPreviousDay[len(jobDifficulty)-1]
}

func recurrence(j int, di int, jobDifficulty []int, memoizedCostPreviousDay []int) int {
	lowestCost := math.MaxInt
	maxCost := 0
	for i := j; i >= di; i-- {
		if jobDifficulty[i] > maxCost {
			maxCost = jobDifficulty[i]
		}
		currentCost := maxCost + memoizedCostPreviousDay[i-1]
		if currentCost < lowestCost {
			lowestCost = currentCost
		}
	}
	return lowestCost
}
