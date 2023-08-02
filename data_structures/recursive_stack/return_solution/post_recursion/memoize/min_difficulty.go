package data_structures

import "math"

// these are interconnected subproblems combined for a solution -> dynamic programming
// can either do bottom up tabulation or top down recursion with memoization -> try top down first

// 3 things need to define
// 1. base cases
// 2. recurrence relation -> more options than standard dp as only constraint is have minimum 1 per day
// (this is the key alarm bell that more complicated dp)
// 3. state variables

// complexity
// time = O(n^2 d) because there are n d possible states (of lowest cost) and we need n time to calculate each one
// space = O(n d) for all the n d states
func minDifficulty(jobDifficulty []int, d int) int {
	// setup data structures
	memoizedCost := make([][]int, len(jobDifficulty))
	for i := range memoizedCost {
		memoizedCost[i] = make([]int, d)
		for j := 0; j < d; j++ {
			memoizedCost[i][j] = -1
		}
	}
	sd := shareddata{
		jobDifficulty: jobDifficulty,
		memoizedCost:  memoizedCost,
	}
	return sd.minDifficultyRecursive(0, d)
}

type shareddata struct {
	jobDifficulty []int
	memoizedCost  [][]int // remaining index j & number of days d
}

// state = index of remaining jobs, num of remaining days, memoized calculation & jobDifficulty
func (sd *shareddata) minDifficultyRecursive(j int, d int) int {
	// base cases
	if len(sd.jobDifficulty) < d {
		return -1
	} else if d == 1 {
		maxCost := 0
		for i := j; i < len(sd.jobDifficulty); i++ {
			if sd.jobDifficulty[i] > maxCost {
				maxCost = sd.jobDifficulty[i]
			}
		}
		return maxCost
	}

	// memoized cases
	if sd.memoizedCost[j][d-1] != -1 {
		return sd.memoizedCost[j][d-1]
	}

	// recurrence relation
	lowestCost := math.MaxInt
	maxCost := 0
	for i := j; i < len(sd.jobDifficulty)-d+1; i++ {
		if sd.jobDifficulty[i] > maxCost {
			maxCost = sd.jobDifficulty[i]
		}
		currentCost := maxCost + sd.minDifficultyRecursive(i+1, d-1)
		if currentCost < lowestCost {
			lowestCost = currentCost
		}
	}
	sd.memoizedCost[j][d-1] = lowestCost
	return lowestCost
}
