package data_structures

// get(key) = map[key]node O(1)
// put(key) = map[key]node O(1)
// _track ordering = remove & insert linked list O(1)

type LRUCache struct {
	capacity int
	nodeMap  map[int]*Node
	head     *Node
	tail     *Node
}

type Node struct {
	prev  *Node
	next  *Node
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity: capacity, nodeMap: map[int]*Node{}}
}

func (lru *LRUCache) Get(key int) int {
	node, ok := lru.nodeMap[key]
	if ok {
		// most recently used so front of linked list
		lru.remove(node)
		lru.addHead(node)
		return node.value
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	// check if key already taken
	node, ok := lru.nodeMap[key]
	if ok {
		lru.remove(node)
		node.value = value
		lru.addHead(node)
		// capacity unchanged
		return
	} else {
		node := &Node{key: key, value: value}
		if len(lru.nodeMap) >= lru.capacity {
			delete(lru.nodeMap, lru.tail.key)
			lru.remove(lru.tail)
		}
		lru.nodeMap[key] = node
		lru.addHead(node)
	}

}

func (lru *LRUCache) addHead(node *Node) {
	// update node
	node.prev = nil
	node.next = lru.head
	// update head and tail
	if lru.head != nil {
		lru.head.prev = node
	}
	lru.head = node
	if lru.tail == nil {
		lru.tail = node
	}
}

func (lru *LRUCache) addTail(node *Node) {
	// update node
	node.prev = lru.tail
	node.next = nil
	// update head and tail
	if lru.head == nil {
		lru.head = node
	}
	if lru.tail != nil {
		lru.tail.next = node
	}
	lru.tail = node
}

func (lru *LRUCache) remove(node *Node) {
	// update previous node & head
	if node != lru.head {
		node.prev.next = node.next
	} else {
		lru.head = node.next
	}
	// update next node & tail
	if node != lru.tail {
		node.next.prev = node.prev
	} else {
		lru.tail = node.prev
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
