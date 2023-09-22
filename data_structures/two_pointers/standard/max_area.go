package data_structures

func maxArea(height []int) int {
	maxArea, left, right := 0, 0, len(height)-1
	for left < right {
		width := right - left
		area := min(height[left], height[right]) * width
		maxArea = max(maxArea, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
