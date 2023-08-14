3 types of graphs
* Undirected graphs = edges between two vertices do not have a direction indicating a two-way relationship
* Directed graphs = edges between two vertices do have a direction indicating a one-way relationship
* Weighted graph = edges between two vertices have an associated weight

Spanning tree = is a connected subgraph in an undirected graph where all vertices are connected with the minimum number of edges
Minimum spanning tree = a spanning tree with the minimum possible total edge weight in a weighted undirected graph
Cut = is a partition of vertices in a graph into two disjoint subsets.
Crossing = is an edge that connects a vertex in one set with a vertex in the other set.
Cut property = For any cut C of the graph, if the weight of an edge E in the cut-set of C is strictly smaller than the weights of all other edges of the cut-set of C, then this edge belongs to all MSTs of the graph

Kruskal's algorithm = heap + union find
1. sort all edges by weight
2. add minimum edge that doesn't produce cycle
3. repeat until V-1 edges added 
Time complexity O(E logE) with O(E logE) for sort, O(Ea(V)) for cycle (connected components) check
Space complexity O(V) V