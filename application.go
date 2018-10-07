package main

import (
	"fmt"
	"heap"
)

func main() {
	var l = []int{4, 8, 9, 1, 2, 5, 7, 6, 3, 0}
	// fmt.Printf("%v\n", l)

	// var h = new(heap.Heap)

	// heap.BuildHeap(l)
	// fmt.Printf("%v\n", l)
	// fmt.Printf("%v\n", heap.Pop(l))
	// fmt.Printf("%v\n", l)

	h := new(heap.Heap)
	h.BuildHeap(l)

	h.Look()
	h.Pop()
	h.Look()

	heap.Sort(l)

	fmt.Printf("%v\n", l)

	// fmt.Printf("%v\n", l)

}
