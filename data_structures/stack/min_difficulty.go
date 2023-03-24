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

// don't understand why this works and the commented out implementation does not
func minDifficulty(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
	if n < d {
		return -1
	}

	minDiffPrevDay := make([]int, n)
	minDiffCurrDay := make([]int, n)

	for i := 0; i < n; i++ {
		minDiffPrevDay[i] = math.MaxInt32
	}

	for day := 0; day < d; day++ {
		stack := make([]int, 0)
		for i := day; i < n; i++ {
			if i > 0 {
				minDiffCurrDay[i] = minDiffPrevDay[i-1] + jobDifficulty[i]
			} else {
				minDiffCurrDay[i] = jobDifficulty[i]
			}

			for len(stack) != 0 && jobDifficulty[stack[len(stack)-1]] <= jobDifficulty[i] {
				j := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				diffIncr := jobDifficulty[i] - jobDifficulty[j]
				if minDiffCurrDay[j]+diffIncr < minDiffCurrDay[i] {
					minDiffCurrDay[i] = minDiffCurrDay[j] + diffIncr
				}
			}

			if len(stack) != 0 {
				if minDiffCurrDay[stack[len(stack)-1]] < minDiffCurrDay[i] {
					minDiffCurrDay[i] = minDiffCurrDay[stack[len(stack)-1]]
				}
			}
			stack = append(stack, i)
		}
		minDiffCurrDay, minDiffPrevDay = minDiffPrevDay, minDiffCurrDay
	}
	return minDiffPrevDay[n-1]
}

// func minDifficulty(jobDifficulty []int, d int) int {
// 	if len(jobDifficulty) < d {
// 		return -1
// 	}
// 	// setup data structures
// 	memoizedCostPreviousDay := make([]int, len(jobDifficulty))

// 	// base cases - what we do on first day
// 	maxCost := 0
// 	for j := 0; j < len(jobDifficulty); j++ {
// 		if jobDifficulty[j] > maxCost {
// 			maxCost = jobDifficulty[j]
// 		}
// 		memoizedCostPreviousDay[j] = maxCost
// 	}

// 	// tabulated cases - what we do on the dith day
// 	for di := 1; di < d; di++ {
// 		memoizedCostCurrentDay := make([]int, len(jobDifficulty))
//         difficultyStack := make([]int, 0, len(jobDifficulty))
// 		for j := di; j < len(jobDifficulty); j++ {
//             memoizedCostCurrentDay[j] = memoizedCostPreviousDay[j - 1] + jobDifficulty[j]
// 			// recurrence relation
// 			memoizedCostCurrentDay[j], difficultyStack = recurrenceStack(
// 				j,
// 				jobDifficulty,
// 				memoizedCostPreviousDay,
//                 memoizedCostCurrentDay,
//                 difficultyStack,
// 			)
// 		}
// 		memoizedCostPreviousDay = memoizedCostCurrentDay
// 	}
// 	return memoizedCostPreviousDay[len(jobDifficulty)-1]
// }

// func recurrenceStack(j int, jobDifficulty []int, memoizedCostPreviousDay []int, memoizedCostCurrentDay []int, difficultyStack []int) (int, []int){
//     n := len(difficultyStack)
//     lowestCost := 10000
//     // more difficult job do today?
//     for n > 0 && jobDifficulty[difficultyStack[n-1]] <= jobDifficulty[j] {
//         i := difficultyStack[n-1]
//         difficultyStack = difficultyStack[:n-1]
//         n = len(difficultyStack)
//         diffIncrease := jobDifficulty[j] - jobDifficulty[i]
//         lowestCost = memoizedCostCurrentDay[j]
//         if memoizedCostCurrentDay[i] + diffIncrease < lowestCost {
//             lowestCost = memoizedCostCurrentDay[i] + diffIncrease
//         }
//     }
//     if memoizedCostCurrentDay[j] != 0 && memoizedCostCurrentDay[j] < lowestCost {
//         lowestCost = memoizedCostCurrentDay[j]
//     }
//     if n > 0 && lowestCost > memoizedCostCurrentDay[difficultyStack[n-1]] {
//         lowestCost = memoizedCostCurrentDay[difficultyStack[n-1]]
//     }
//     difficultyStack = append(difficultyStack, j)
// 	return lowestCost, difficultyStack
// }