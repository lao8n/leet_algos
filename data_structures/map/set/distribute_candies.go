package data_structures

// clarifying questions
// * slice of candies - different nums denote different types
// * n always even so don't need to handle odd case
// * is input always sorted? it is not
// approaches
// * for loop up to n / 2 where increment types if find new - and increment pointer if not new
// * could use map - but if sorted no need to just track prev
func distributeCandies(candyType []int) int {
	// init
	count, seen := 0, map[int]bool{}
	// for loop
	for _, candy := range candyType {
		if count == len(candyType)/2 {
			return count
		}
		if !seen[candy] {
			count++
			seen[candy] = true
		}
	}
	return count
}
