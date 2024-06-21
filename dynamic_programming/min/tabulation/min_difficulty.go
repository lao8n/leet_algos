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
	memoizedCost := make([][]int, len(jobDifficulty))
	for i := range memoizedCost {
		memoizedCost[i] = make([]int, d)
	}
	// [6] [-1]
	// [6] [11]
	// [6] [10]
	// [6] [9]
	// [6] [8]
	// [-1][7]
	// j = jobs done
	// di = days done
	// base cases - what we do on first day
	maxCost := 0
	for j := 0; j < len(jobDifficulty)-(d-1); j++ {
		if jobDifficulty[j] > maxCost {
			maxCost = jobDifficulty[j]
		}
		memoizedCost[j][0] = maxCost
	}

	// tabulated cases - what we do on the dith day
	for di := 1; di < d; di++ {
		for j := di; j <= len(jobDifficulty)-(d-di); j++ {
			// recurrence relation
			memoizedCost[j][di] = recurrence(j, di, jobDifficulty, memoizedCost)
		}
	}
	return memoizedCost[len(jobDifficulty)-1][d-1]
}

func recurrence(j int, di int, jobDifficulty []int, memoizedCost [][]int) int {
	lowestCost := math.MaxInt
	maxCost := 0
	for i := j; i >= di; i-- {
		if jobDifficulty[i] > maxCost {
			maxCost = jobDifficulty[i]
		}
		currentCost := maxCost + memoizedCost[i-1][di-1]
		if currentCost < lowestCost {
			lowestCost = currentCost
		}
	}
	return lowestCost
}
