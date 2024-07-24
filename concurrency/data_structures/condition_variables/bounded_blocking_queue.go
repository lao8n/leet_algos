// package main

// this approach uses one condition per state
// tradeoff is that could wake up goroutine that is not waiting for that condition
// advantage is that only one condition
package data_structures

import (
	"fmt"
	"sync"
	"time"
)

// a channel is a bounded blocking queue
type BoundedBlockingQueue struct {
	capacity int
	queue    []int
	mutex    *sync.Mutex
	cond     *sync.Cond
}

func Constructor(capacity int) BoundedBlockingQueue {
	m := sync.Mutex{}
	c := sync.NewCond(&m)
	queue := BoundedBlockingQueue{
		capacity: capacity,
		queue:    make([]int, 0, capacity),
		mutex:    &m,
		cond:     c,
	}
	return queue
}

func (q *BoundedBlockingQueue) Enqueue(element int) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	for q.Size() == q.capacity { // full cannot enqueue
		q.cond.Wait() // unlocks and suspends goroutine
		// can only be awoken by Broadcast or Signal
		// when awoken need to check condition again hence for loop
	}
	q.queue = append(q.queue, element)
	q.cond.Broadcast()
}

func (q *BoundedBlockingQueue) Dequeue() int {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	for q.Size() == 0 { // empty cannot dequeue
		q.cond.Wait()
	}
	element := q.queue[0]
	q.queue = q.queue[1:]
	q.cond.Broadcast()
	return element
}

// only called safely by one goroutine
func (q *BoundedBlockingQueue) Size() int {
	return len(q.queue)
}

func main() {
	tests := []struct {
		name         string
		threads      string
		capacity     int
		functions    []string
		args         [][]int
		expected     []int
		expectedSize int
	}{
		{
			name:         "example 1",
			threads:      "one_each",
			capacity:     2,
			functions:    []string{"enqueue", "dequeue", "dequeue", "enqueue", "enqueue", "enqueue", "enqueue", "dequeue"},
			args:         [][]int{{1}, nil, nil, {0}, {2}, {3}, {4}, nil},
			expected:     []int{1, 0, 2},
			expectedSize: 2,
		},
		{
			name:         "example 2",
			threads:      "one_per",
			capacity:     3,
			functions:    []string{"enqueue", "enqueue", "enqueue", "dequeue", "dequeue", "dequeue", "enqueue"},
			args:         [][]int{{1}, {0}, {2}, nil, nil, nil, {3}},
			expected:     []int{1, 0, 2},
			expectedSize: 1,
		},
	}

	for _, test := range tests {
		fmt.Println(test.name)
		queue := Constructor(test.capacity)
		solution := []int{}
		if test.threads == "one_each" {
			go func() {
				for i, function := range test.functions {
					if function == "enqueue" {
						if test.args[i] != nil {
							queue.Enqueue(test.args[i][0])
						}
					}
				}
			}()
			go func() {
				for _, function := range test.functions {
					if function == "dequeue" {
						solution = append(solution, queue.Dequeue())
					}
				}
			}()
		} else if test.threads == "one_per" {
			for i, function := range test.functions {
				if function == "enqueue" {
					if test.args[i] != nil {
						j := i
						go func() {
							queue.Enqueue(test.args[j][0])
						}()
					}
					time.Sleep(1 * time.Millisecond)
				} else if function == "dequeue" {
					go func() {
						solution = append(solution, queue.Dequeue())
					}()
				}
			}
		}
		time.Sleep(1 * time.Second)
		if len(solution) != len(test.expected) {
			fmt.Printf("different lengths %v %v\n", solution, test.expected)
		} else {
			for i := 0; i < len(test.expected); i++ {
				if solution[i] != test.expected[i] {
					fmt.Printf("different %v %v\n", solution, test.expected)
					break
				}
				if i == len(test.expected)-1 {
					fmt.Println("test passed")
				}
			}
		}
		if test.expectedSize != queue.Size() {
			fmt.Printf("queue is %v size not %v", queue.Size(), test.expectedSize)
		} else {
			fmt.Println("queue size test passes")
		}
	}
}
