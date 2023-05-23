package data_structures

type MovingAverage struct {
	maxSize    int
	rollingSum int
	queue      []int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		maxSize:    size,
		rollingSum: 0,
		queue:      make([]int, 0),
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.queue) == this.maxSize {
		this.rollingSum -= this.queue[0]
		this.queue = this.queue[1:]
	}
	this.rollingSum += val
	this.queue = append(this.queue, val)
	return float64(this.rollingSum) / float64(len(this.queue))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
