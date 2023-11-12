package data_structures

// clarification questions
// * increment bottom k elements by val
// approaches
// * brute force is loop over stack elements and actually increment but costs O(k)
// * bottom elements is interesting as top elements come and go but bottom are not changed so often -> could have a rule where loop over rules like if index in stack i < some value then add x- crucially after popping element we need to be clever and update rules
type CustomStack struct {
	stack   []int
	maxSize int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		stack:   []int{},
		maxSize: maxSize,
	}
}

func (this *CustomStack) Push(x int) {
	if len(this.stack) < this.maxSize {
		this.stack = append(this.stack, x)
	}
}

func (this *CustomStack) Pop() int {
	if len(this.stack) > 0 {
		popped := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		return popped
	}
	return -1
}

func (this *CustomStack) Increment(k int, val int) {
	for i := 0; i < k && i < len(this.stack); i++ {
		this.stack[i] += val
	}
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
