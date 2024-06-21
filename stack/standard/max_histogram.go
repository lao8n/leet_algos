package data_structures

// approaches
// * increasing stack - if lower num then use it as right index and compute all rectangles with what's on stack as left side and height
func largestRectangleArea(heights []int) int {
	max := 0
	stack := []data{} // of indices

	for i, h := range heights {
		prevI := i
		for len(stack) > 0 && stack[len(stack)-1].h > h {
			area := stack[len(stack)-1].h * (i - stack[len(stack)-1].i)
			if area > max {
				max = area
			}
			prevI = stack[len(stack)-1].i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, data{prevI, h})
	}

	for _, data := range stack {
		area := data.h * (len(heights) - data.i)
		if area > max {
			max = area
		}
	}
	return max
}

type data struct {
	i int
	h int
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
