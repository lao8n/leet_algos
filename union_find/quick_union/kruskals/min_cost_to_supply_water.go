package data_structures

import "sort"

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

// actually rather than dynamic programming you can use the cut theorem and prims algorithm where you have a collection of visited nodes and add minimum edges to it but this would require a heap
// or you can use kruskal's algorithm and leverage the union find i already have setup

// key trick (which i missed) is to add a virtual node which is connected to every node which represents cost of well
// func minCostToSupplyWater(n int, wells []int, pipes [][]int) int {
// 	// setup union find with well per house
// 	houses := make([]waterSupply, n)
// 	for i, well := range wells {
// 		houses[i] = waterSupply{
// 			well: i, // house connected to its own well
// 			cost: well,
// 		}
// 	}
// 	uf := unionFind{houses: houses}

// 	// sort pipes
// 	sort.Slice(pipes, func(x, y int) bool {
// 		return pipes[x][2] < pipes[y][2]
// 	})

// 	// try pipes
// 	for _, pipe := range pipes {
// 		x, y, xyCost := pipe[0], pipe[1], pipe[2]
// 		uf.union(x-1, y-1, xyCost) // all the houses are -1 indexed
// 	}

// 	totalCost := 0
// 	for _, house := range uf.houses {
// 		totalCost += house.cost
// 	}

// 	return totalCost
// }

// type unionFind struct {
// 	houses []waterSupply
// }

// type waterSupply struct {
// 	well int // i.e. root
// 	cost int // either pipe or well
// }

// // should find
// func (uf *unionFind) find(x int) int {
// 	if uf.houses[x].well == x {
// 		return x
// 	}
// 	uf.houses[x].well = uf.find(uf.houses[x].well) // root parent well
// 	return uf.houses[x].well
// }

// func (uf *unionFind) connected(x int, y int) bool {
// 	return uf.find(x) == uf.find(y)
// }

// func (uf *unionFind) union(x int, y int, xyCost int) {
// 	xRoot, yRoot := uf.find(x), uf.find(y)
// 	// not connected
// 	if xRoot != yRoot {
// 		// connect using x well
// 		if xyCost < uf.houses[y].cost && uf.houses[x].cost < uf.houses[y].cost {
// 			uf.houses[y].cost = xyCost
// 			// connect using y well
// 		} else if xyCost < uf.houses[x].cost && uf.houses[y].cost < uf.houses[x].cost {
// 			uf.houses[x].cost = xyCost
// 			// don't connect
// 		} else if uf.houses[x].cost < xyCost && uf.houses[y].cost < xyCost {
// 			// do nothing
// 		}
// 	}
func minCostToSupplyWater(n int, wells []int, pipes [][]int) int {
	// setup union find with well per house
	root := make([]int, n+1)
	rank := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		root[i] = i
		rank[i] = 1
	}
	uf := unionFind{root: root, rank: rank}

	// make a new node at 0 which represents well and connects to every house
	for i, well := range wells {
		pipes = append(pipes, []int{0, i + 1, well}) // from 0 well to ith house
	}
	// sort pipes
	sort.Slice(pipes, func(x, y int) bool {
		return pipes[x][2] < pipes[y][2]
	})

	totalCost := 0
	for _, pipe := range pipes {
		h1, h2, cost := pipe[0], pipe[1], pipe[2]
		if !uf.connected(h1, h2) {
			uf.union(h1, h2)
			totalCost += cost
		}
	}

	return totalCost
}

type unionFind struct {
	root []int
	rank []int
}

func (uf *unionFind) find(x int) int {
	if uf.root[x] == x {
		return x
	}
	uf.root[x] = uf.find(uf.root[x])
	return uf.root[x]
}

func (uf *unionFind) connected(x int, y int) bool {
	return uf.find(x) == uf.find(y)
}

func (uf *unionFind) union(x int, y int) {
	xRoot, yRoot := uf.find(x), uf.find(y)
	// not connected
	if xRoot != yRoot {
		if uf.rank[xRoot] < uf.rank[yRoot] {
			uf.root[xRoot] = yRoot
		} else if uf.rank[yRoot] < uf.rank[xRoot] {
			uf.root[yRoot] = xRoot
		} else {
			uf.root[xRoot] = yRoot
			uf.rank[yRoot] += 1
		}
	}
}
