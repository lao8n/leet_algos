package recursion

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
	data := data{
		jobDifficulty: jobDifficulty,
		memoizedCost:  memoizedCost,
	}
	return data.minDifficultyRecursive(0, d)
}

type data struct {
	jobDifficulty []int
	memoizedCost  [][]int // remaining index j & number of days d
}

// state = index of remaining jobs, num of remaining days, memoized calculation & jobDifficulty
func (data *data) minDifficultyRecursive(j int, d int) int {
	// base cases
	if len(data.jobDifficulty) < d {
		return -1
	} else if d == 1 {
		maxCost := 0
		for i := j; i < len(data.jobDifficulty); i++ {
			if data.jobDifficulty[i] > maxCost {
				maxCost = data.jobDifficulty[i]
			}
		}
		return maxCost
	}

	// memoized cases
	if data.memoizedCost[j][d-1] != -1 {
		return data.memoizedCost[j][d-1]
	}

	// recurrence relation
	lowestCost := math.MaxInt
	maxCost := 0
	for i := j; i < len(data.jobDifficulty)-d+1; i++ {
		if data.jobDifficulty[i] > maxCost {
			maxCost = data.jobDifficulty[i]
		}
		currentCost := maxCost + data.minDifficultyRecursive(i+1, d-1)
		if currentCost < lowestCost {
			lowestCost = currentCost
		}
	}
	data.memoizedCost[j][d-1] = lowestCost
	return lowestCost
}
