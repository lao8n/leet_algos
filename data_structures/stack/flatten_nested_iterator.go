package data_structures

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */
// use a stack to slowly unnest each layer
// we want to always leave the stack with its top value as an integer ready to go
type NestedIterator struct {
	stack    []*NestedInteger
	topIndex int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	ni := &NestedIterator{
		stack:    nestedList,
		topIndex: 0,
	}
	ni.getIntegerAtTopOfStack()
	return ni
}

func (this *NestedIterator) getIntegerAtTopOfStack() {
	// assume len stack > 0 as has next already called
	if this.topIndex >= len(this.stack) {
		return
	}
	topOfStack := this.stack[this.topIndex]
	if !topOfStack.IsInteger() {
		nestedList := topOfStack.GetList()
		// we use this opportunity to reset stack
		this.stack = append(nestedList, this.stack[this.topIndex+1:]...)
		// check integer is top of stack again
		this.topIndex = 0
		this.getIntegerAtTopOfStack()
	}
	// if integer is top of stack work is done
}

func (this *NestedIterator) Next() int {
	// we have already guaranteed this is top of stack
	topOfStack := this.stack[this.topIndex]
	this.topIndex += 1
	this.getIntegerAtTopOfStack()
	return topOfStack.GetInteger()
}

func (this *NestedIterator) HasNext() bool {
	return this.topIndex < len(this.stack)
}
