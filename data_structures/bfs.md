Time Complexity = O(E + V) - traverse every edge and check every vertex
Space Complexity = O(V) - max size of queue

Use when:
* Order level traversal of vertices
* Shortest path between two points - only for unweighted graphs
* Deeper graphs (as opposed to wider)

Steps
1. Build adjacency list & memoized set
2. Setup queue and loop through until empty
3. Check for base cases and neighbours