/*
Package heap :
Author: John Lahut
Date: 10/6/2018
Project: Project #2: Priority Queue
Purpose: Implements a minheap data structure
	Heap is slice based, indexed at zero
	Assume the slice is of type int
Filename: heap.go
Package: heap
*/
package heap

import (
	"fmt"
	"globals"
)

// -- public functions --

// Heap : represents a heap in memory
type Heap struct {
	Arr []globals.Item
}

// BuildHeap : constructor for Heap, inits the Heap struct with arr, heapifies it
func (heap *Heap) BuildHeap(arr []globals.Item) {

	// copy into struct
	heap.Arr = make([]globals.Item, len(arr))
	copy(heap.Arr, arr)

	// heapify starting at first non-leaf node
	for i := len(arr) / 2; i >= 0; i-- {
		heap.Heapify(i)
	}
}

// Pop : returns the root node to caller, heap is now one element smaller
func (heap *Heap) Pop() globals.Item {
	// min element is always root
	root := heap.Arr[0]

	// replace root with right most child, cut slice down by one
	heap.Arr[0] = heap.Arr[len(heap.Arr)-1]
	heap.Arr = heap.Arr[:len(heap.Arr)-1]

	// re-heapify at root
	heap.Heapify(0)
	return root
}

// Push : inserts the element into the heap, re builds the heap
func (heap *Heap) Push(i globals.Item) {
	heap.Arr = append(heap.Arr, i)
	heap.BuildHeap(heap.Arr)
}

// Look : prints current state of heap to console
func (heap *Heap) Look() {
	fmt.Printf("%v\n", heap.Arr)
}

// Sort : sorts the passed in list in ascending order (min heap)
func Sort(arr []globals.Item) {

	// build a heap, and pop each one off, replacing in given list -> BOOM! sorted
	heap := new(Heap)
	heap.BuildHeap(arr)
	for i := 0; i < len(arr); i++ {
		arr[i] = heap.Pop()
	}
}

// Heapify : validate the heap based on the root passed in, it is assumed the heap was valid until the most recent operation
func (heap *Heap) Heapify(root int) {
	// assume the max node is the root, get the children
	min := root
	length := len(heap.Arr)
	left, right := left(root), right(root)

	// if we are in the bounds of array, and a child is greater than parent, mark it
	if left < length && heap.Arr[min].Pri > heap.Arr[left].Pri {
		min = left
	}
	if right < length && heap.Arr[min].Pri > heap.Arr[right].Pri {
		min = right
	}
	// check to see if child is greater than parent
	if root != min {
		// if so, swap the child and parent and re-heapify
		heap.Arr[root], heap.Arr[min] = heap.Arr[min], heap.Arr[root]
		heap.Heapify(min)
	}
}

// -- private functions --

// gets the left child of a zero-indexed array repr. of heap
func left(i int) int {
	return 2*i + 1
}

// gets the right child of a zero-indexed array repr. of head
func right(i int) int {
	return 2*i + 2
}
