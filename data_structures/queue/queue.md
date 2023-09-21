**Queue**

Definition = A data structure where elements are stored and removed in a FIFO manner.

Time complexity = `O(n)`

Space complexity = `O(n)`

**Types**

**BFS**
```
// level_order.go
for len(queue) > 0 {
    levelSize := len(queue)
    ...
    for i := 0; i < levelSize; i++ {
        current := queue[0]
        queue = queue[1:]
        ...
        for _, child := range current.Children {
            queue = append(queue, child)
        }
    }
}
```
`level_order.go` is a good example of standard BFS with the added twist of retaining level information, which can be done with a single queue where you use the trick of taking the size of the queue, iterating over that size, but still appending to the end. Other variations like `num_squares.go` iterate over perfect squares, or `all_paths_source_target.go` over neighbours, or `connect_nodes.go` over left and right nodes rather than children but the principal is the same. 

**BFS with visited**
```
// num_islands.go
for len(queue) != 0 {
    pop := queue[0]
    queue = queue[1:]
    for ... {
        if ... grid[x][y] == '1' {
            grid[x][y] = '0'
            queue = append(queue, []int{x, y})
        }
    }
}
```
Sometimes it is necessary to track whether a node has been visited. This can either be done in processing the item off the queue or often it is better to not add it to the queue at all unless it hasn't been visited as with the above example. Examples of the former include `01_matrix.go` or `can_visit_all_rooms.go`. Examples of latter include `find_links.go`

**Dijkstra's algorithm**
```
// minimum_effort_path.go
for i := 0; i < ql; i++ {
    popped := queue[0]
    queue = queue[1:]
    for _, neighbour := range getNeighbours(popped, numRows, numCols){
        ...
        dp[neighbour.x][neighbour.y] = diff
        queue = append(queue, neighbour)
    }
}
```
Dijkstra's algorithm is an approach to finding the single shortest path from a single vertex to all other nodes of a weighted graph with non-negative weights (otherwise need to use a Bellman-ford dynamic programming approach). It is considered greedy because it does not have to calculate all paths, instead it uses a queue (and BFS) to ensure which points are processed first (although typical implementations use a priority queue)

**Kahn's algorithm**
```
// course_schedule.go
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
```
Kahn's algorithm is a way to find the topological ordering of a directed-acyclic graph. It uses 3 main data structures: 
1. adjacency list to maintain neighbours/dependency relationships 
2. an `inDegree := make(map[int]int, numCourses)` map to measure the number of dependencies a node has
3. a queue to process the 0-degree nodes
Some variations like `minimum_semesters.go` do not track an ordering but the `numSemesters` or `find_min_height_tree.go` where rather than looking to process 0-degree nodes instead just trim the leaves with `if inDegreeMap[neighbour] == 1` back until none left.

**Other data structures**
```
// circular_queue.go
this.end = (this.end + 1) % this.k
this.queue[this.end] = value
this.num += 1
```
Queues can be used to make a number of other data structures including `stack_using_queues.go`, `moving_average.go` and `circular_queue.go`