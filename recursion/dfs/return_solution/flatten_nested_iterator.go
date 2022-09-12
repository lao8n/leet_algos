package recursion

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
type NestedIterator struct {
	flattenedList []int
	index         int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{
		flattenedList: flattenList(nestedList),
		index:         0,
	}
}

func flattenList(nestedList []*NestedInteger) []int {
	// base case
	if len(nestedList) == 0 {
		return []int{}
	} else if len(nestedList) == 1 {
		if nestedList[0].IsInteger() {
			return []int{nestedList[0].GetInteger()}
		} else {
			return flattenList(nestedList[0].GetList())
		}
	}
	// recursive
	return append(flattenList([]*NestedInteger{nestedList[0]}), flattenList(nestedList[1:])...)
}

func (this *NestedIterator) Next() int {
	next := this.flattenedList[this.index]
	this.index += 1
	return next
}

func (this *NestedIterator) HasNext() bool {
	return this.index < len(this.flattenedList)
}
