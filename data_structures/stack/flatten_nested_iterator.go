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

// Space complexity = O(N + L) where we have stack being expanded list of integers and numbers

// Time O(1) - no need to reverse as just maintain topIndex pointer to the front
func Constructor(nestedList []*NestedInteger) *NestedIterator {
	ni := &NestedIterator{
		stack:    nestedList,
		topIndex: 0,
	}
	ni.setIntegerAtTopOfStack()
	return ni
}

// Time O(1) for integer or O(L/N) where L is the total number of lists and O(N).
// where the complexity comes from O(N + L) where need to see every list and number once
// and then amoritized this is O(N + L) / O(N) = O(1 + L/N) = O(L/N)
func (this *NestedIterator) setIntegerAtTopOfStack() {
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
		this.setIntegerAtTopOfStack()
	}
	// if integer is top of stack work is done
}

// O(1) except if need to call setIntegerAtTopOfStack so then O(L/N)
func (this *NestedIterator) Next() int {
	// we have already guaranteed this is top of stack
	topOfStack := this.stack[this.topIndex]
	this.topIndex += 1
	this.setIntegerAtTopOfStack()
	return topOfStack.GetInteger()
}

// O(1)
func (this *NestedIterator) HasNext() bool {
	return this.topIndex < len(this.stack)
}
