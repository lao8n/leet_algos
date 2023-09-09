package data_structures

type MyCircularQueue struct {
	k     int
	queue []int
	num   int
	start int
	end   int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		k:     k,
		queue: make([]int, k),
		num:   0,
		start: 0,
		end:   -1,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.end = (this.end + 1) % this.k
	this.queue[this.end] = value
	this.num += 1
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.start = (this.start + 1) % this.k
	this.num -= 1
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.start]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.end]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.num == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.num == this.k
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
