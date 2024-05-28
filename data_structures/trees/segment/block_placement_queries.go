package data_structures

import "slices"

// clarifying questions
// * number line 0 to inf
// * 1. build obstacle at x 2. place block of size sz anywhere in range 0, x
// approaches
// * for any x it would be good to know the max size it can support
// specifics
// * update tree on query 1 or query 2
// * segment tree with time complexity O(n logm), space O(n)
func getResults(queries [][]int) []bool {
	obs := make([]int, 0, 1<<5)
	res := make([]bool, 0, len(queries))
	st := &segmentTree[int]{
		data: make(map[int]int),
		size: 5 * 1e4,
	}
	for _, q := range queries {
		i, _ := slices.BinarySearch(obs, q[1])
		switch q[0] {
		case 1: // place obs
			obs = slices.Insert(obs, i, q[1])
			prevObs := 0
			if i > 0 {
				prevObs = obs[i-1]
			}
			st.Update(obs[i], obs[i]-prevObs)
			if i < len(obs)-1 {
				st.Update(obs[i+1], obs[i+1]-obs[i])
			}
		case 2:
			maxGap := q[1]
			if i > 0 {
				maxGap = max(q[1]-obs[i-1], st.Query(0, obs[i-1]))
			}
			res = append(res, maxGap >= q[2])
		}
	}
	return res
}

type SegmentTree[T ~int] interface { // T can be any type whose underlying type is int
	Query(l, r int) T
	Update(pos int, val T)
}

type segmentTree[T ~int] struct {
	data map[int]T
	size int
}

func (st *segmentTree[T]) Query(l, r int) T {
	// return sum of left and right queries
	l += st.size
	r += st.size
	res := st.data[r]
	for l < r {
		// if left is a right node
		if l&1 != 0 {
			res = max(res, st.data[l])
		}
		l++ // move inside

		// if right is a right node bring value of left node
		if r&1 != 0 {
			r--
			res = max(res, st.data[r])
		}
		// go to parent
		l >>= 1
		r >>= 1
	}
	return res
}

func (st segmentTree[T]) Update(pos int, val T) {
	pos += st.size // start at leaf
	// update from leaft to root
	st.data[pos] = val
	for pos > 1 {
		pos >>= 1
		st.data[pos] = max(st.data[2*pos], st.data[2*pos+1])
	}
}

func max[T ~int](x, y T) T {
	if x > y {
		return x
	}
	return y
}
