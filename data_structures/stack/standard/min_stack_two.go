package data_structures

// problem with tuple of value and rolling min is often the minimum doesn't change so there is a lot of repetition
// one solution is to have a stack of minimums instead where you coordinate between the two stacks by
// 1. when popping also pop from min stack 2. when pushing also push to min stack (including if == value on min stack)
// to deal with popping logic
// then you can go even further by having an index to track number of occurrences of minimum on min stack

type MinStack struct {
	stack       []int
	size        int
	secondStack SecondStack
}

func Constructor() MinStack {
	return MinStack{stack: []int{}, size: 0, secondStack: SecondStackConstructor()}
}

func (this *MinStack) Push(val int) {
	// push to stack of values
	if this.size < len(this.stack) {
		this.stack[this.size] = val
	} else {
		this.stack = append(this.stack, val)
	}
	this.size++
	if this.secondStack.size == 0 || val <= this.secondStack.Top() {
		this.secondStack.Push(val)
	}

}

func (this *MinStack) Pop() {
	if this.Top() == this.secondStack.Top() {
		this.secondStack.Pop()
	}
	this.size--
}

func (this *MinStack) Top() int {
	top := this.stack[this.size-1]
	return top
}

func (this *MinStack) GetMin() int {
	return this.secondStack.Top()
}

type SecondStack struct {
	stack []int
	size  int
}

func SecondStackConstructor() SecondStack {
	return SecondStack{stack: []int{}, size: 0}
}

func (this *SecondStack) Push(val int) {
	if this.size < len(this.stack) {
		this.stack[this.size] = val
	} else {
		this.stack = append(this.stack, val)
	}
	this.size++
}

func (this *SecondStack) Pop() {
	this.size--
}

func (this *SecondStack) Top() int {
	return this.stack[this.size-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
