package data_structures

// approaches
// * set O(1) need to use a map
// * randomised behaviours - problem is map is not ordered (but not randomly ordered either).. have a slice with index? golang runtime actually randomizes order of keys.. could rely on this - how update slice with remove in O(1) time? swap and remove end...

import "math/rand"

type RandomizedSet struct {
	nums []int
	set  map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nums: make([]int, 0),
		set:  make(map[int]int), // num -> index in nums
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.set[val]; ok {
		return false // item was present so cannot insert
	} else {
		this.set[val] = len(this.nums) // insert val at index len(nums)
		this.nums = append(this.nums, val)
		return true
	}
}

func (this *RandomizedSet) Remove(val int) bool {
	if i, ok := this.set[val]; ok {
		// move element to the end of slice
		lastIndex := len(this.nums) - 1
		lastElement := this.nums[lastIndex]
		this.set[lastElement] = i                                               // update map
		this.nums[i], this.nums[lastIndex] = this.nums[lastIndex], this.nums[i] // swap

		// remove
		delete(this.set, val)
		this.nums = this.nums[:lastIndex] // unappend
		return true
	} else {
		// item not present
		return false
	}
}

func (this *RandomizedSet) GetRandom() int {
	r := rand.Intn(len(this.nums)) // [0, n)
	return this.nums[r]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
