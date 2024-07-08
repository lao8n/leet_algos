package data_structures

// clarifying questions
// * next greater element -> first greater element to the right - are they sorted?
// * nums1 is a subset of nums2 - ordered subset? no.

// approaches
// * naive is have map(v -> index) then search for larger value for that index.
// * maybe can do it in one go with a decreasing stack so if reach a value if greater than stack then add it..
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int) // v -> next greater element index
	stack := make([]int, 0)
	for i, num2 := range nums2 {
		// pop from stack first
		for len(stack) > 0 && num2 > nums2[stack[len(stack)-1]] {
			popped := stack[len(stack)-1]
			v := nums2[popped]
			m[v] = i // num2 is next greater element
			stack = stack[:len(stack)-1]
		}
		// then add to stack
		stack = append(stack, i)
	}
	// fmt.Println(m)
	// we override original slice
	solution := make([]int, len(nums1))
	for i, num1 := range nums1 {
		if j, ok := m[num1]; ok {
			solution[i] = nums2[j]
		} else {
			solution[i] = -1
		}
	}
	return solution
}
