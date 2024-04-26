package data_structures

// approach
// * kahn's topological sorting algorithm - either with bfs or dfs
func canFinish(numCourses int, prerequisites [][]int) bool {
	// adj list & in-degree map
	inDegree := make([]int, numCourses) // course num -> count
	adjList := make(map[int][]int, numCourses)
	for _, pr := range prerequisites {
		course, pre := pr[0], pr[1]
		inDegree[course]++
		adjList[pre] = append(adjList[pre], course)
	}

	// setup queue and add 0 prerequisite courses
	queue := make([]int, 0)
	for course, count := range inDegree {
		if count == 0 {
			queue = append(queue, course)
		}
	}
	coursesTaken := 0
	// process queue
	for len(queue) > 0 {
		popped := queue[0]
		queue = queue[1:]
		coursesTaken++
		for _, neighbour := range adjList[popped] {
			inDegree[neighbour]-- // we have taken prerequisite
			if inDegree[neighbour] == 0 {
				queue = append(queue, neighbour)
			}
		}
	}

	return coursesTaken == numCourses
}
