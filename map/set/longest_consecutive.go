package data_structures

// clarifying questions
// * we just want the length not the sequence itself
// * no space constraint
// * there can also be duplicates

// approaches
// * sort is O(n logn) -> doesn't work
// * cycle sort doesn't work because not between fixed range of 1 to n (max is 10^9)
// * heap doesn't work because though heapify is O(n) extracting is O(n logn)
// * map with each number - but also check number above and number below -> create seq length

// specifics
// * problem is if have a seq where legnth is 3, 3, 3 -> 1 need to loop through full seq length updating all
// * how get updated information - only the start and finish are guaranteed to be updated?
func longestConsecutive(nums []int) int {
	set := make(map[int]bool, 0)
	for _, num := range nums {
		set[num] = true
	}

	longestStreak := 0
	for num, _ := range set {
		if !set[num-1] {
			currentNum := num
			currentStreak := 1
			for set[currentNum+1] { // only loop for the beginning
				currentNum++
				currentStreak++
			}
			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		}
	}
	return longestStreak
}
