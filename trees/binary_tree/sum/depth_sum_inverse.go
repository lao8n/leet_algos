package data_structures

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */

// approaches
// * recursively open up NestedInteger with accumulated depth
// * return a list then calculate on list - but how keep track of depths?
// specifics
// * do we have to calculate weight first? or can we do it greedily? (m - d + 1) * v = mv - dv + v.. perhaps we can calculate sum of dvs - then once we have max at end sum mvs and v - this is m (v + 1)
func depthSumInverse(nestedList []*NestedInteger) int {
	d := data{}
	dv, v := 0, 0
	for _, nI := range nestedList {
		idv, iv := d.recurse(nI, 1)
		dv += idv
		v += iv
	}
	// fmt.Println(d.maxDepth, v, dv)
	return (d.maxDepth+1)*v - dv
}

// sum(w * v) = sum((m - d + 1) * v)
// sum(mv - dv + v) = sum(v(m + 1) - dv)
// m * sum(v + 1) - sum(dv)
type data struct {
	maxDepth int
}

func (d *data) recurse(nestedInteger *NestedInteger, depth int) (dv int, v int) {
	if nestedInteger == nil {
		return 0, 0
	}
	if depth > d.maxDepth {
		d.maxDepth = depth
	}

	// 2 options
	// 1. integer
	if nestedInteger.IsInteger() {
		v = nestedInteger.GetInteger()
		dv = v * depth
	} else {
		// 2. list
		vs := nestedInteger.GetList()
		if len(vs) != 0 {
			for _, nI := range vs {
				idv, iv := d.recurse(nI, depth+1)
				dv += idv
				v += iv
			}
		}
	}
	return dv, v
}
