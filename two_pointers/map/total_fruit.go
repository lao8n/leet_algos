package data_structures

func totalFruit(fruits []int) int {
	l, r := 0, 0
	basket := make(map[int]int)
	maxLen := 0
	for r < len(fruits) {
		basket[fruits[r]]++
		for len(basket) > 2 {
			basket[fruits[l]]--
			if basket[fruits[l]] == 0 {
				delete(basket, fruits[l])
			}
			l++
		}
		if r-l > maxLen {
			maxLen = r - l
		}
		r++
	}
	return maxLen + 1
}
