package data_structures

// queue is fifo, stack is lifo
// clue is use you stacks and stack has push to top and pop from top
// if push x, y, z into stack [x, y, z] with z on top
// when want to pop could then pop into new stack where reverse order for [z, y, x]
// problem is although push is O(1), pop is always O(n)
// [x, y, z] -> [z, y, x]
// [x, y] -> [y, x]
type MyQueue struct {
	pushStack []int
	popStack  []int
}

func Constructor() MyQueue {
	return MyQueue{
		pushStack: []int{},
		popStack:  []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.pushStack = append(this.pushStack, x)
}

func (this *MyQueue) Pop() int {
	popped := this.Peek()
	if len(this.popStack) > 0 {
		this.popStack = this.popStack[:len(this.popStack)-1]
	}
	return popped
}

func (this *MyQueue) Peek() int {
	if len(this.popStack) > 0 {
		peeked := this.popStack[len(this.popStack)-1]
		return peeked
	}
	// need to transfer
	topPushStack := len(this.pushStack) - 1
	this.popStack = make([]int, len(this.pushStack))
	for i := topPushStack; i >= 0; i-- {
		this.popStack[topPushStack-i] = this.pushStack[i]
	}
	peeked := this.pushStack[0]
	this.pushStack = make([]int, 0)
	return peeked
}

func (this *MyQueue) Empty() bool {
	return len(this.pushStack) == 0 && len(this.popStack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
