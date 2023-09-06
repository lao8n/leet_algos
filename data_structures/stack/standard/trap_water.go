package data_structures

func trap(height []int) int {
	// 2 approaches
	// 1. sum vertical strips = at each point look for maximum either side (using two pointers)
	// 2. sum horizontal strips = store heights down valley (in stack), and then create strips up valley
	var sum int
	var stack []int // down-valley heights
	for stripBottom, stripRight := 0, 0; stripRight < len(height); stripRight++ {
		for len(stack) > 0 && height[stripRight] > height[stripBottom] {
			stack = stack[:len(stack)-1] // pop
			if len(stack) > 0 {
				stripLeft := stack[len(stack)-1] // top
				stripWidth := stripRight - stripLeft - 1
				stripHeight := min(height[stripRight], height[stripLeft]) - height[stripBottom]
				sum += stripWidth * stripHeight
				stripBottom = stack[len(stack)-1]
			}
		}
		stack = append(stack, stripRight) // push
		stripBottom = stack[len(stack)-1]
	}
	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
