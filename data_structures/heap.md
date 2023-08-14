Priority Queue = abstract data type where each element has a priority associated with it, where elements with high priority are served before elements with low priority.

Implementations of Priority Queues
* Arrays and Linked Lists with O(1) and O(n) for insertion and deletion (or vice-versa)
* Creation of heap is O(n) time and space complexity. Rather than inserting all elements which is O(n logn) instead heapifying is Sum_i=0^log N i / 2^(i + 1) = n x Sum_i=0^logN i / 2^(i+1) = n. Size is O(1) because we have a single variable to track
* Heap O(log n) for both insertion and deletion. Max/min can be obtained with O(1)

Definition
* Complete (all levels filled except lowest - and then from left) binary (max two children per node) tree
* For Min Heap the value of each node must be <= the value of its child nodes. For Max Heap >=

Insertion 
* Add node as left-most child
* If max/min then swap recursively

Deletion
* If deleting max/min move left-most child to top, then remove top element (which is now in left-most child spot)
* If max/min then swap recursively - if min always exchange with smallest child and vice-versa

Implementing a heap as an array
* How to find parent node of n? n/2 (rounding down) and keep 0th index empty
* How to find the child node of n? n * 2 for left and n * 2 + 1 for right
* Is a node a leaf? If n > N / 2 

Applications of Heap
1. Heap Sort = selection sort is O(n^2) as it involves finding n minimum values and moving it to the front. 
Process: 1. Bottom up heapification O(n) 2. Removal of n elements which is O(n logn) (removal can be in place with heap over array with size n-1). 
However heapsort is undesirable because it 1. Is not a stable sort, 2. Bad cache locality properties because it swaps elements based on locations in heaps with read operations to access indices in random order
2. Top-K Problem = 1. Bottom up heapification O(n) 2. Removal of k elements which is O(k logn). So combined complexity is O(k logn + n). There is also a variation where if you want smallest you build a max heap (and vice-versa) of size k elements and compare remaining n-k elements with top of heap where 1. Bottom up heapification is O(k) 2. Removal of up to (n-k) elements which is O((n-k) log k + k) = O(n log k)
3. K-th element = Exactly the same as Top-K

Approaches for Kth largest element
* 1. Min heap of size n -> remove n-k elements (equivalently until heap of size k) -> kth largest
* 2. Min heap of size k -> only insert if larger than minimum (harder with negative numbers)
* 3. Max heap of size n -> remove k elements

Approaches to Go Heap
```
type Interface interface {
		sort.Interface
		Push(x interface{})
		Pop() interface{}
}

type Heap []freq

func (h *Heap) Push(x interface{}) { *h = append(*h, x.(freq)) }
// note in popped we take n-1 value as end because the source code does a swap first between 0 and n-1
// if we want min value need to go to 0
func (h *Heap) Pop() interface{} {
	previous, n, popped := *h, h.Len(), freq{}
	*h, popped = previous[0:n-1], previous[n-1]
	return popped
}
func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].freq < h[j].freq } // make it greater than
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
```