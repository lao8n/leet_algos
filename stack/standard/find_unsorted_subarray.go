package data_structures

// approach = use stack to track where to put index
func findUnsortedSubarray(nums []int) int {
	// find increasing nums -> then min
	incStack := make(stack, 0)
	leftIndex := len(nums)
	for i, num := range nums {
		for len(incStack) > 0 && num < nums[incStack.Peek()] {
			if popped := incStack.Pop(); popped < leftIndex {
				leftIndex = popped
			}
		}
		incStack.Push(i)
	}
	// find decreasing nums -> max
	decStack := make(stack, 0)
	rightIndex := 0
	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]
		for len(decStack) > 0 && num > nums[decStack.Peek()] {
			if popped := decStack.Pop(); popped > rightIndex {
				rightIndex = popped
			}
		}
		decStack.Push(i)
	}
	if rightIndex-leftIndex <= 0 {
		return 0
	}
	return rightIndex - leftIndex + 1
}

type stack []int

func (s *stack) Push(x int) {
	*s = append(*s, x)
}
func (s *stack) Pop() int {
	popped := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return popped
}
func (s stack) Peek() int {
	return s[len(s)-1]
}
