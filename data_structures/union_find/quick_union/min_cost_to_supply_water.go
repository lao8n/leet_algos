package data_structures

// WARNING - this doesn't work just storing it here for posterity
// approach - use union find until have a single component
// assumption - probably cannot just greedily choose which wells & pipes to lay - maybe need to use dynamic programming?
// union find - quick union with 1. memoized find 2. union by rank
// dynamic programming - 1. top down with recursion & memoization 2. bottom up iteration with tabulation
// if we do dynamic programming there are 3 questions
// 1. base case
// 2. state - n house well or pipe? - might need to lay pipe for well that isn't connected yet?
// 3. recurrence case - choice between 1. add a well 2. add a pipe for the nth house given n-1 houses optimised
// probably just have an array of W x P i.e. number of wells * number of pipes
// choice
// 1. dp first 2. union find where use union find to track whether done
// 1. union find 2. dp -> probably better where expand model to include total cost
// could include cost but what is cost of a pipe? - whatever is downstream of well
// maybe rather than slowly adding wells and pipes you start with a well for every house and then see where you can save with pipes?
// rather than union by rank - i guess we want to union by minimum cost?
// key problem i'm running into is how to systematically try every combination
func minCostToSupplyWater(n int, wells []int, pipes [][]int) int {
	// setup union find with well per house
	houses := make([]waterSupply, n)
	for i, well := range wells {
		houses[i] = waterSupply{
			well: i, // house connected to its own well
			cost: well,
		}
	}
	uf := unionFind{houses: houses}
	// try pipes
	for _, pipe := range pipes {
		x, y, xyCost := pipe[0], pipe[1], pipe[2]
		uf.union(x-1, y-1, xyCost) // all the houses are -1 indexed
	}

	totalCost := 0
	for _, ws := range uf.houses {
		totalCost += ws.cost
	}
	return totalCost
}

type unionFind struct {
	houses []waterSupply
}

type waterSupply struct {
	well int // i.e. root
	cost int // either pipe or well
}

// should find
func (uf *unionFind) find(x int) int {
	if uf.houses[x].well == x {
		return x
	}
	uf.houses[x].well = uf.find(uf.houses[x].well) // root parent well
	return uf.houses[x].well
}

func (uf *unionFind) union(x int, y int, xyCost int) {
	xRoot, yRoot := uf.find(x), uf.find(y)
	// not connected
	if xRoot != yRoot {
		// choice:
		// keep houses with separate wells i.e. not connected
		// connect two houses with a pipe
		// just choose cheapest if either costs (doesn't matter if well or pipe)
		if xyCost < uf.houses[x].cost || xyCost < uf.houses[y].cost {
			if uf.houses[x].cost < uf.houses[y].cost {
				// replace more expensive house costs with cheaper xyCost
				uf.houses[y].cost = xyCost
			} else {
				uf.houses[x].cost = xyCost
			}
		}
		// else better to keep the two components separate
	}
}
