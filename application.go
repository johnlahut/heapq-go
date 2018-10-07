package main

import (
	"fmt"
	"globals"
	"pq"
)

func main() {
	queue := new(pq.PriorityQueue)
	queue.Enqueue(globals.Item{Pri: 4, Key: "job1"})
	queue.Enqueue(globals.Item{Pri: 3, Key: "job2"})
	queue.Dequeue()
	queue.Enqueue(globals.Item{Pri: 0, Key: "job3"})
	queue.Enqueue(globals.Item{Pri: 1, Key: "job4"})
	queue.Dequeue()

	fmt.Printf("%v\n", queue.Dequeue().Key)
	fmt.Printf("%v\n", queue.Dequeue().Key)

}
