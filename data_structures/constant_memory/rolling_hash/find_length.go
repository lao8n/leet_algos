package data_structures

import "fmt"

// interconnected problems combined for a solution -> dynamic programming

// 2 approaches
// 1. top down recursion with memoization
// 2. bottom up iteration with tabulation

// 3 components
// 1. base cases
// 2. state -> index in nums1, index in nums2 - longest array starting at n1, n2 or longest array from n1, n2?
// 3. recurrence ->

// binary search - is there a subarray with length l common to A & B?
// for loop for all substrings of length l in A and then manually check B
// but we can do better by memoizing substrings we've already looked for with a rolling hash (robin-karp algorithm)

// time complexity O(m+n * log(min(m, n))) where m = len(nums1) & n = len(nums2)
// log factor from binary search and O(m+n) for creating the rolling hashes
// space complexity

// doesn't work - not sure what i'm doing wrong...
func findLength(nums1 []int, nums2 []int) int {
	prime := 13
	mod := 107
	d := data{
		nums1: nums1,
		nums2: nums2,
		prime: prime,
		mod:   mod,
	}

	// binary search
	lo, hi := 0, min(len(nums1), len(nums2))
	for lo < hi {
		mi := (lo + hi) / 2
		if d.check(mi) {
			lo = mi + 1
		} else {
			hi = mi
		}
	}
	return lo
}

type data struct {
	nums1        []int
	nums2        []int
	prime        int
	mod          int
	primeInverse int
}

// is there a subarray with length l common to A & B
func (d *data) check(guessLength int) bool {
	rollingHashes := make(map[int][]int, len(d.nums1)) // hash
	for i, _ := range rollingHashes {
		rollingHashes[i] = []int{}
	}

	// nums1 hashes for each possible starting location
	for possFinish, hash := range d.rollingHash(d.nums1, guessLength) {
		rollingHashes[hash] = append(rollingHashes[hash], possFinish)
	}

	for nums2PossFinish, hash := range d.rollingHash(d.nums2, guessLength) {
		if nums1PossFinishes, ok := rollingHashes[hash]; ok {
			for _, nums1PossFinish := range nums1PossFinishes {
				// fmt.Printf("%v vs %v\n", d.nums1[nums1PossStart:nums1PossStart + guessLength], d.nums2[nums2PossStart: nums2PossStart + guessLength])
				if compare(d.nums1[nums1PossFinish-guessLength:nums1PossFinish], d.nums2[nums2PossFinish-guessLength:nums2PossFinish]) {
					return true
				}
			}
		}
	}
	return false
}

func (d *data) rollingHash(nums []int, guessLength int) []int {
	// e.g. if len is 3 on slice of len 10 there are 8 possible starting points
	if guessLength == 0 {
		return []int{}
	}
	hash, power := 0, 1
	finishIndex := guessLength - 1
	numsRolling := make([]int, len(nums))
	powers := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		// building up hash
		powers[i] = power
		if i < finishIndex {
			hash = (hash*power + nums[i]) % d.mod
		} else {
			hash = (hash - numsRolling[i-finishIndex]*powers[i]/powers[i-finishIndex]) % d.mod
		}
		for hash < 0 {
			hash += d.mod
		}
		numsRolling[i] = hash
		power = (power * d.prime) % d.mod
	}
	fmt.Println(guessLength, nums, numsRolling, power)
	return numsRolling[finishIndex:]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func compare(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
