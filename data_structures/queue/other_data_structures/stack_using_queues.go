package data_structures

// queue is fifo, stack is lifo
// clue is to use two queues
// push -> add to queue 1
// pop -> if queue 2 has element

// stack [x, y] + [z] -> pop would yield z
// queue1 [x, y, z] -> x? O(1)
// queue2 - but we want z so then pop O(n) to other queue - back and forth?
// can even go further and just use one queue.

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{
		queue: make([]int, 0),
	}
}

// queue access 0th element and add to n element
func (this *MyStack) Push(x int) {
	// we push all elements to the back of the queue
	// this.queue = append(x, this.queue...)
	this.queue = append(this.queue, x)
	for i := 0; i < len(this.queue)-1; i++ {
		popped := this.Pop()
		this.queue = append(this.queue, popped)
	}
}

func (this *MyStack) Pop() int {
	popped := this.queue[0]
	this.queue = this.queue[1:]
	return popped
}

func (this *MyStack) Top() int {
	top := this.queue[0]
	return top
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
