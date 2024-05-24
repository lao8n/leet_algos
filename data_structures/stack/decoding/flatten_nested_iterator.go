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
// approach
// * recursive solution? return nested values as list?
// * use a stack to manage recursion?

// specifics
// * after every next need to move i to next valid value so HasNext is valid
// * i recurse into structure but how do i get out of it? we want to recurse a single path but not return everything
// * lots of nested iterators?
// * front of stack is first element, more expensive appends but don't have to reverse
type NestedIterator struct {
	stack []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{
		stack: nestedList,
	}
}

func (this *NestedIterator) Next() int {
	popped := this.stack[0] // is an integer
	this.stack = this.stack[1:]
	return popped.GetInteger()
}

func (this *NestedIterator) HasNext() bool { // could be empty lists though?
	// process top of stack
	if len(this.stack) == 0 {
		return false
	}
	popped := this.stack[0]
	for !popped.IsInteger() {
		this.stack = append(popped.GetList(), this.stack[1:]...)
		if len(this.stack) == 0 {
			return false
		}
		popped = this.stack[0]
	}
	return true
}
