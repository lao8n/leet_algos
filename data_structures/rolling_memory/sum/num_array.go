package data_structures

// clarifying questions
// * calculate sum between l & r

// approaches
// * segment tree
// * could just do rolling sum where r - [l-1]
type NumArray struct {
	rolling []int
}

func Constructor(nums []int) NumArray {
	rolling := make([]int, len(nums))
	sum := 0
	for i, num := range nums {
		sum += num
		rolling[i] = sum
	}
	return NumArray{
		rolling: rolling,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	leftSum, rightSum := 0, this.rolling[right]
	if left > 0 {
		leftSum = this.rolling[left-1]
	}
	return rightSum - leftSum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
