package recursion

// these are interconnected subproblems combined for a solution -> dynamic programming
// can either do bottom up tabulation or top-down recursion with memoization -> try latter for exp
func minCostClimbingStairs(cost []int) int {
	oneStep, twoSteps := recursiveMinCost(cost)
	if oneStep < twoSteps {
		return oneStep
	}
	return twoSteps
}

// recurse from base cases
// memoize last two numbers only
// return cost of 1 step and cost of 2 steps
func recursiveMinCost(cost []int) (oneStep int, twoSteps int) { //
	// base cases
	if len(cost) == 1 {
		oneStep = cost[0]
		// recursive step
	} else {
		// return cost of 1 step and 2 steps (i.e. skipping 1 step)
		recursiveOneStep, recursiveTwoSteps := recursiveMinCost(cost[1:])
		// question: shall we step cost[0]?
		// if yes
		oneStep = recursiveOneStep + cost[0]
		if recursiveTwoSteps+cost[0] < oneStep {
			oneStep = recursiveTwoSteps + cost[0]
		}
		// if no
		twoSteps = recursiveOneStep // i.e. skip cost[0]
	}
	return oneStep, twoSteps
}
