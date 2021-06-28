For dynamic programming https://leetcode.com/problems/longest-increasing-path-in-a-matrix/solution/
with a transition function we can avoid the need for depth first search

However, the problem is you need to have the dependencies in order (known as topological order)
Topological sorting = for DAG is a linear ordering of vertices such that for every directed
edge (u,v) vertex u comes before vertex v in the ordering

The vertices that don't depend upon others are called leaves - which we put in a list and remove 
them from the dag, we do this repeatedly until we have a valid topological order for the vertices. 