package data_structures

// kahn's algorithm for topological sorting
// queue + in-degree map
func minimumSemesters(n int, relations [][]int) int {
	// create adjacency list = pre -> course
	// in-degree = course -> number of pre-requisites
	adjList := make(map[int][]int, 0)
	inDegree := make(map[int]int, n)
	for _, relation := range relations {
		p, c := relation[0], relation[1]
		adjList[p] = append(adjList[p], c)
		inDegree[c] += 1
	}

	// process no pre-requisite courses first
	queue := make([]int, 0)
	for c := 1; c <= n; c++ {
		if inDegree[c] == 0 {
			queue = append(queue, c)
		}
	}
	// process queue
	numSemesters := 0
	numCourses := 0
	for len(queue) > 0 {
		numCoursesInSemester := len(queue)
		for i := 0; i < numCoursesInSemester; i++ {
			// process node
			popped := queue[0]
			queue = queue[1:]
			numCourses++
			// process neighbours
			for _, neighbour := range adjList[popped] {
				inDegree[neighbour]--
				if inDegree[neighbour] == 0 {
					queue = append(queue, neighbour)
				}
			}
		}
		numSemesters++
	}
	if numCourses != n {
		return -1
	}
	return numSemesters
}
