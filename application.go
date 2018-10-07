package main

import (
	"fmt"
	"globals"
	"pq"
)

func main() {
	queue := new(pq.PriorityQueue)
	queue.Enqueue(globals.Item{4, "job1"})
	queue.Enqueue(globals.Item{3, "job2"})
	queue.Dequeue()
	queue.Enqueue(globals.Item{0, "job3"})
	queue.Enqueue(globals.Item{1, "job4"})
	queue.Dequeue()

	fmt.Printf("%v\n", queue.Dequeue().Key)
	fmt.Printf("%v\n", queue.Dequeue().Key)

}
