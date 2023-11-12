package data_structures

// clarifying questions
// * if num of children == 2 then you have a boundary where int64(n + 1 - 2 * max(n - limit, 0))
// approaches

func distributeCandies(n int, limit int) int64 {
	limit = min(limit, n)
	var ans int64
	for i := 0; i <= limit; i++ {
		// to the 1st child we can give between 0 and limit candies
		// if the remaining number of candies is > limit * 2 then this is invalid as even if we give max to 2nd and
		// 3rd child we will still have some left over
		if n-i > limit*2 {
			continue
		}
		// we have two children remaining amongst which we must distribute n - i candies...
		// to the 2nd child we give between 0 and minimum of limit and n - i candies... and then to the third child
		// we can give between 0 and n - i - limit candies and 0
		ans += int64(min(limit, n-i) - max(0, n-i-limit) + 1)
	}
	return ans
}

func min(x1, x2 int) int {
	if x1 < x2 {
		return x1
	}
	return x2
}

func max(x1, x2 int) int {
	if x1 > x2 {
		return x1
	}
	return x2
}
