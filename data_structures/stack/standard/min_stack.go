package data_structures

// stack = can get all O(1) operations with an array
// min stack
// cannot use a heap, bst etc. because not O(1)
// key idea is use a stack with a rolling min

// O(1) for push, pop, top, getMin
// O(n) space complexity if all push
type MinStack struct {
	stack []valueMinPair // value & rolling min
	size  int
}

type valueMinPair struct {
	value int
	min   int
}

func Constructor() MinStack {
	return MinStack{stack: []valueMinPair{}, size: 0}
}

func (this *MinStack) Push(val int) {
	// get new min
	stackMin := val
	if this.size > 0 {
		min := this.GetMin()
		if min < stackMin {
			stackMin = min
		}
	}
	// push to stack
	newEntry := valueMinPair{value: val, min: stackMin}
	if len(this.stack) > this.size {
		this.stack[this.size] = newEntry
	} else {
		this.stack = append(this.stack, newEntry)
	}
	this.size++
}

func (this *MinStack) Pop() {
	this.size--
}

func (this *MinStack) Top() int {
	top := this.stack[this.size-1]
	return top.value
}

func (this *MinStack) GetMin() int {
	top := this.stack[this.size-1]
	return top.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
