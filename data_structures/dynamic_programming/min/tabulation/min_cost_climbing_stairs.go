package data_structures

// approaches
// * dp - tabulation or recursion
// * rather than array can just have last two steps
// * go forward or go backwards? forwards
func minCostClimbingStairs(cost []int) int {
	curStepCost, prevOneStepCost, prevTwoStepCost := 0, 0, 0
	for i := 0; i < len(cost); i++ {
		curStepCost = min(prevOneStepCost+cost[i], prevTwoStepCost+cost[i])
		prevTwoStepCost = prevOneStepCost
		prevOneStepCost = curStepCost
	}
	return min(curStepCost, prevTwoStepCost)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
