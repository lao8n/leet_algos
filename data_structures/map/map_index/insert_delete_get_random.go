package data_structures

// insertion O(1) need a list
// deletion O(1) need a linked list - but cannot use because also need random access
// randomised need length of list which is O(n) but if update after every operation it's just +/-1, not necessary in golang because slice's are pointers to an array with length already maintained

import (
	"math/rand"
)

type RandomizedSet struct {
	set  map[int]int // val -> nums index
	nums []int       // val
}

func Constructor() RandomizedSet {
	return RandomizedSet{set: make(map[int]int), nums: []int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.set[val]; ok {
		return false
	}
	index := len(this.nums) // O(1) for slices
	this.nums = append(this.nums, val)
	this.set[val] = index
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if index, ok := this.set[val]; ok {
		swapIndex := len(this.nums) - 1
		this.nums[index] = this.nums[swapIndex]
		this.set[this.nums[index]] = index
		this.nums = this.nums[:swapIndex] // reduce len by 1
		delete(this.set, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(this.nums))
	return this.nums[index]
}
