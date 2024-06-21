package data_structures

type NumArray struct {
	n    int
	tree []int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	tree := make([]int, n*2)
	if n > 0 {
		// add leaves
		for i, j := n, 0; i < 2*n; i, j = i+1, j+1 {
			tree[i] = nums[j]
		}
		// add roots
		for i := n - 1; i > 0; i-- {
			tree[i] = tree[i*2] + tree[i*2+1]
		}

	}
	return NumArray{
		n:    n,
		tree: tree,
	}
}

// O(log n)
func (this *NumArray) Update(index int, val int) {
	index += this.n // go to leaf
	this.tree[index] = val
	for index > 0 {
		left, right := index, index
		if index%2 == 0 { // index is a left hand child
			right = index + 1
		} else { // index is a right hand child
			left = index - 1
		}
		// update parent
		this.tree[index/2] = this.tree[left] + this.tree[right]
		index /= 2
	}
}

// O(log n)
func (this *NumArray) SumRange(left int, right int) int {
	// get sum by [left->][<-right] i.e. until l & r meet
	// leaves
	left, right = left+this.n, right+this.n
	sum := 0
	for left <= right {
		// left is a right child of its parent, don't need parent sum
		if left%2 == 1 {
			sum += this.tree[left]
			left++
		}
		// right is a left child of its parent, don't need parent sum
		if right%2 == 0 {
			sum += this.tree[right]
			right--
		}
		left /= 2
		right /= 2
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
