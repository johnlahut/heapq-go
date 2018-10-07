package pq

import (
	"fmt"
	"globals"
	"heap"
)

// PriorityQueue : Implements Heap for data structure
type PriorityQueue struct {
	heap heap.Heap
}

// Enqueue : enqueue the passed in element
func (pq *PriorityQueue) Enqueue(item globals.Item) {
	pq.heap.Push(item)
}

// Dequeue : removes the highest priority element, returning to caller
func (pq *PriorityQueue) Dequeue() globals.Item {
	return pq.heap.Pop()
}

// Look : prints the current state of the queue to the caller
func (pq *PriorityQueue) Look() {
	fmt.Printf("%v\n", pq.heap.Arr)
}

// IsEmpty : returns true if the current queue is empty, false otherwise
func (pq *PriorityQueue) IsEmpty() bool {
	return len(pq.heap.Arr) == 0
}
