package data_structures

// clarifying questions
// * permutations = order matters. want a random sorting
// * just use a random number generator? how prevent calls where repeated?

// approaches
// * n swaps
// * copy in each number 1 by 1 randomly choosing index -> problem of choosing index already taken? - could reduce the size of the array but costs O(n) to copy

import "math/rand"

type Solution struct {
	orig []int
	cur  []int
}

func Constructor(nums []int) Solution {
	cur := make([]int, len(nums))
	copy(cur, nums)
	return Solution{
		orig: nums,
		cur:  cur,
	}
}

func (this *Solution) Reset() []int {
	newCur := make([]int, len(this.orig))
	copy(newCur, this.orig)
	this.cur = newCur
	return this.cur
}

func (this *Solution) Shuffle() []int {
	for i := 0; i < len(this.cur); i++ {
		j := rand.Intn(len(this.cur))
		// fmt.Println(i, j)
		this.cur[i], this.cur[j] = this.cur[j], this.cur[i]
	}
	// fmt.Println(this.cur, this.orig)
	return this.cur
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
