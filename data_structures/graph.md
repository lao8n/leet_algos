3 types of graphs
* Undirected graphs = edges between two vertices do not have a direction indicating a two-way relationship
* Directed graphs = edges between two vertices do have a direction indicating a one-way relationship
* Weighted graph = edges between two vertices have an associated weight

Spanning tree = is a connected subgraph in an undirected graph where all vertices are connected with the minimum number of edges
Minimum spanning tree = a spanning tree with the minimum possible total edge weight in a weighted undirected graph
Cut = is a partition of vertices in a graph into two disjoint subsets.
Crossing = is an edge that connects a vertex in one set with a vertex in the other set.
Cut property = For any cut C of the graph, if the weight of an edge E in the cut-set of C is strictly smaller than the weights of all other edges of the cut-set of C, then this edge belongs to all MSTs of the graph

Similarities
* all use greedy approaches which can be proven by proof by contradiction (if non-greedy path was shorter)

Kruskal's algorithm = sorted edges + union find
1. sort all edges by weight
2. add minimum edge that doesn't produce cycle
3. repeat until V-1 edges added 
Time complexity = O(E logE) with O(E logE) for sort, O(Ea(V)) for cycle (connected components) check
Space complexity = O(V) V

Prim's algorithm = heap
1. divide nodes into visited and unvisited
2. of the visited set add its minimum edge
3. repeat until all vertices are added
Time complexity (Binary heap) = O(E logV) with O(V + E) to traverse vertices, extracting min element is O(logV) so overall is O(V + E) * O(logV) = O(E log V)
Time complexity (Fibonacci heap) = O(E + V logV)  with O(logV) to extract min element and O(1) key decreasing operation
Space complexity = O(V) vertices in data structure
Note = don't need deletable heap, as can just keep adding new nodes

Dijkstra's algorithm = bfs with heap
1. build adjacency list of neighbours
2. build map of cost of path initialized to MAXINT
3. build heap of neighbours 
4. do bfs on 
Time complexity = O(E + E logV) for binary heap rather than O(E V) for bfs
Space complexity = O(E + V)
Note = only works on weighted graphs with non-negative weights

Bellman-Ford = 
Note = if n nodes and no negative weight cycle shortest path between two nodes is n-1, if negative weight cycle no solution