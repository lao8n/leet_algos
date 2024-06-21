// approaches
// * fenwick tree, also known as a binary index tree
// * time complexity O(n log m), space complexity O(m)
func countSmaller(nums []int) []int {
	offset := 10 * 10 * 10 * 10
	size := 2*offset + 2 // num of values between -10^4 and 10^4 + one dummy
	tree := fenwickTree{
		data: make([]int, size),
		size: size,
	}

	result := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		// get all numbers between min and up to but not including nums[i]
		result[len(nums)-1-i] = tree.query(nums[i] + offset)
		// add number
		tree.update(nums[i]+offset, 1)
	}

	return reverse(result)
}

type fenwickTree struct {
	data []int
	size int
}

func (t *fenwickTree) update(i int, v int) {
	i += 1 // bit index is shifted one to the right
	// update from leaf to root
	for i < t.size {
		t.data[i] += v
		// twos complement -i, where & isolates lowest set bit of i, effectively subtracting
		// the smallest power of two that divides the number
		i += i & -i // increment by the smallest power of 2 that divides it
	}
}

func (t *fenwickTree) query(i int) int {
	// return sum of [0, index)
	result := 0
	for i >= 1 {
		result += t.data[i]
		i -= i & -i
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