package data_structures

// use kahns algorithm - basically a queue
func findOrder(numCourses int, prerequisites [][]int) []int {
	// setup adjacency list & indegree maps
	adjacencyList := make(map[int][]int, numCourses) // out-degree i.e. pre-requisite -> course
	inDegree := make(map[int]int, numCourses)
	for _, prePair := range prerequisites {
		c, p := prePair[0], prePair[1]
		adjacencyList[p] = append(adjacencyList[p], c)
		inDegree[c]++
	}

	// setup queue
	queue := make([]int, 0)
	for vertex := 0; vertex < numCourses; vertex++ {
		if inDegree[vertex] == 0 {
			queue = append(queue, vertex)
		}
	}

	// process zero-degree nodes and neighbours
	solution := make([]int, 0)
	for len(queue) > 0 {
		popped := queue[0]
		queue = queue[1:]
		solution = append(solution, popped)
		for _, neighbour := range adjacencyList[popped] {
			inDegree[neighbour] -= 1
			if inDegree[neighbour] == 0 {
				queue = append(queue, neighbour)
			}
		}
	}

	if len(solution) < numCourses {
		return []int{}
	}
	return solution
}
