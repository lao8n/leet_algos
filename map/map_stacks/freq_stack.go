package data_structures

type FreqStack struct {
    freq map[int]int // int to freq
    stacks map[int][]int // freq to stack
    maxFreq int 
}


func Constructor() FreqStack {
    return FreqStack{
        freq: make(map[int]int),
        stacks: make(map[int][]int),
        maxFreq: 0,
    }
}


func (this *FreqStack) Push(val int)  {
    this.freq[val] += 1
    if this.freq[val] > this.maxFreq {
        this.maxFreq = this.freq[val]
    }
    this.stacks[this.freq[val]] = append(this.stacks[this.freq[val]], val)
}


func (this *FreqStack) Pop() int {
    maxFreqStack := this.stacks[this.maxFreq]
    popped := maxFreqStack[len(maxFreqStack) - 1]
    this.stacks[this.maxFreq] = maxFreqStack[:len(maxFreqStack) - 1]
    this.freq[popped] -= 1
    if len(this.stacks[this.maxFreq]) == 0 {
        this.maxFreq -= 1
    }
    return popped
}


/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */