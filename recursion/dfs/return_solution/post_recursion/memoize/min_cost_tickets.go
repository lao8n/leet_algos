// interdependent problems combined for a solution -> dynamic programming

// 2 approaches
// 1. top down recursion with memoization
// 2. bottom up iteration with tabulation -> probs easier

// either ways need to consider
// 1. base cases
// 2. state = index, day paid for until -> cost
// 3. recurrence = choose between do nothing - or if buy a ticket which ticket to buy

// time complexity O(d x n)
// space complexity O(d x n)
func mincostTickets(days []int, costs []int) int {
	lastDay := days[len(days)-1] + 29
	memoized := make([][]int, len(days))
	for i, _ := range memoized {
		memoized[i] = make([]int, lastDay)
	}
	s := state{days: days, costs: costs, memoized: memoized}
	return s.recursiveMinCostTickets(0, -1)
}

type state struct {
	days     []int
	costs    []int
	memoized [][]int
}

func (s *state) recursiveMinCostTickets(dayIndex int, tillDay int) int {
	// base cases
	if dayIndex == len(s.days) {
		return 0
	}

	// memoized
	if tillDay >= 0 && s.memoized[dayIndex][tillDay] != 0 {
		return s.memoized[dayIndex][tillDay]
	}

	// recurrence
	holdCost, buy1Cost, buy7Cost, buy30Cost := math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt
	today := s.days[dayIndex]
	if tillDay >= today {
		// hold
		holdCost = s.recursiveMinCostTickets(dayIndex+1, tillDay)
	} else {
		buy1Cost = s.costs[0] + s.recursiveMinCostTickets(dayIndex+1, today)
		buy7Cost = s.costs[1] + s.recursiveMinCostTickets(dayIndex+1, today+6)
		buy30Cost = s.costs[2] + s.recursiveMinCostTickets(dayIndex+1, today+29)
	}
	lowestCost := holdCost

	if buy1Cost < lowestCost {
		lowestCost = buy1Cost
	}
	if buy7Cost < lowestCost {
		lowestCost = buy7Cost
	}
	if buy30Cost < lowestCost {
		lowestCost = buy30Cost
	}
	if tillDay != -1 {
		s.memoized[dayIndex][tillDay] = lowestCost
	}

	return lowestCost
}