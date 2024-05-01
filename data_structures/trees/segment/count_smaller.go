package data_structures

// links
// https://leetcode.com/articles/a-recursive-approach-to-segment-trees-range-sum-queries-lazy-propagation/
// https://leetcode.com/problems/count-of-smaller-numbers-after-self/

// approaches
// * loop through for every number O(n^2)
// * stack? just store increasing right->left [6, 1] then get to 2 and we pop 6 and size of stack is number [2, 1]
// * dp? if larger than 2 know that [5,2,6,1] [2,1,1,0]
// * sort the list O(n logn), [6, 5, 2, 1] -> [, 1, 0] -> insert in order?
// * maybe just a heap where keep popping to see where number is?
// * if numbers were decreasing then it would just be [3, 2, 1, 0].. problem is not just a local decision becuase if have 2, 6, 1 - 2 is smaller than 6 but how do we know it is 1 not 0?
// * the only thing that really helps is the next number up or down. if we have 5 and we know 6 and its position? but we don't know the gap.. if you know the index of the next biggest you just have to check until there.. rest of the work is done. if you know the index of the next smallest you also can stop there.

// * segment tree - nums[i] to the right but must be in the range -inf, nums[i] - 1 then just want to
// sum the counts of those numbers. just store in buckets in a map from -10^4 to 10^4
// * time complexity O(n log m) where m is max(num) - min(num) and n is length of nums
// * space complexity need an array of at most 2M + 2 to store the segment tree
func countSmaller(nums []int) []int {
	offset := 10 * 10 * 10 * 10
	size := 2*offset + 1 // num of values between -10^4 and 10^4
	tree := segmentTree{
		data: make([]int, 2*size),
		size: size,
	}

	result := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		// get all numbers between min and up to but not including nums[i]
		result[len(nums)-1-i] = tree.query(0, nums[i]+offset)
		// add number
		tree.update(nums[i]+offset, 1)
	}

	return reverse(result)
}

type segmentTree struct {
	data []int
	size int
}

func (t *segmentTree) update(i int, v int) {
	i += t.size // shift index to leaf
	// update from leaf to root
	t.data[i] += v
	for i > 1 {
		i /= 2
		t.data[i] = t.data[i*2] + t.data[i*2+1]
	}
}

func (t *segmentTree) query(l, r int) int {
	// return sum of left and right
	l += t.size
	r += t.size

	// build up query of all indices between l & r
	result := 0
	for l < r {
		// if left is a right node
		if l%2 == 1 {
			result += t.data[l]
			l += 1 // move inside
		}
		// go to parent
		l /= 2

		// if right is a right node bring value of left node (we want to exclude right node)
		if r%2 == 1 {
			r -= 1
			result += t.data[r]
		}
		// go to parent
		r /= 2
	}
	return result
}

func reverse(result []int) []int {
	l, r := 0, len(result)-1
	for l < r {
		result[l], result[r] = result[r], result[l]
		l++
		r--
	}
	return result
}
