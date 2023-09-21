**Heap**

Definition = A priority queue is an abstract data type where element has a priority associated with it, where elements with high priority are served before elements with low-priority. Implementations of priority queues are usually with heaps

Approaches for Kth largest element
* 1. Min heap of size n -> remove n-k elements (equivalently until heap of size k) -> kth largest
* 2. Min heap of size k -> only insert if larger than minimum (harder with negative numbers)
* 3. Max heap of size n -> remove k elements

Time complexity = Creation of a heap is `O(n)` because rather than inserting all the elements one-on-one one can heapify. Insertion of individual elements is `O(n logn)`. Max/min are be `O(1)`

Space complexity = `O(n)` for all the elements in the heap

**Types**

**Max heap**
```
// find_k_largest.go
h := make(Heap, len(nums))
copy(h, nums)
heap.Init(&h)
for ; k > 1; k-- {
    heap.Pop(&h)
}
return h[0]
...
type Heap[]int
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int))}
func (h *Heap) Pop() interface{} {
    previous, n, popped := *h, h.Len(), 0
    *h, popped = previous[0: n-1], previous[n-1]
    return popped
}
func (h Heap) Len() int { return len(h) } 
func (h Heap) Less(i, j int) bool { return h[i] > h[j] }
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
```
The above has all the main elements of heap questions including `heap.Init` on a copy of the input data and defining the 5 heap methods of the interface which are `Push` and `Pop` plus the `sort.Interface` ones. The approach above to find the `k`th largest involves popping the first `k-1` elements off first. Others like `top_k_frequent_nums` or `top_k_frequent_words` do not ignore those elements but store them in an array instead. Crucially in all cases the `Less` function is defined backwards such that `return h[i] > h[j]`

**Min heap**
```
// kth_largest_stream.go
func (this *KthLargest) Add(val int) int {
    heap.Push(&this.heap, val)
    for this.heap.Len() > this.k {
        heap.Pop(&this.heap)
    }
    return this.heap.Peek()
}
```
Some questions like `kth_smallest.go` are essentially the min heap equivalent of the max heap examples the only difference being that `Less` is defined as `return h[i].val < h[j].val` as you would expect. 

Sometimes though there is a more subtle approach where in the `kth_largest_stream` we actually want the largest elements but we want to replace the smallest so we still use a min heap of a fixed size `k` where we can either have a condition to not insert unless larger, or we insert anyway and then pop until the size is correct.

Other questions like `min_meeting_rooms.go` use heaps to maintain a list of end times, or `furthest_building.go` where have a min heap of ladders using them on the largest jumps. `median_finder.go` goes further by having both a min heap and a max heap where push and pop numbers between the two heaps to keep their sizes the same.

**Min Heap: Dijkstra's**
```
// network_delay_time.go
for h.Len() > 0 {
    popped := heap.Pop(&h).(edge)
    prevTime := minimumTime[popped.source]
    if popped.time + prevTime < minimumTime[popped.target] {
        minimumTime[popped.target] = popped.time + prevTime
        for _, neighbour := range adjList[popped.target] {
            heap.Push(&h, neighbour)
        }
    }
}
```
Dijkstra's is about finding the shortest path from a source node on a weighted graph to all other nodes. In addition to BFS a min heap can be used to greedily update the path times. 

**Min Heap: Prim's**
```
// min_cost_connect_points.go
for h.Len() > 0 {
    ...
    minEdge := heap.Pop(&h).(edge)
    ...
    visited[newPoint] = true
    visitedCount += 1
    rollingCost += minEdge.cost
    for _, neighbourEdge := range adjacencyMap[newPoint]{
        if !(visited[neighbourEdge.i] && visited[neighbourEdge.j]) {
            heap.Push(&h, neighbourEdge)
        }
    }
}
```
To find the minimum spanning tree there are two approaches. Kruskal's adds the minimum edge they all connect. Prim's does the opposite adding only the connected edges, although greedily picking the minimum with a heap