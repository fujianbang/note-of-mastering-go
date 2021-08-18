package main

import (
	"fmt"
)

type QueueNode struct {
	Value int
	Next  *QueueNode
}

var sizeOfQueue = 0
var queue = new(QueueNode)

func Push(t *QueueNode, v int) bool {
	if queue == nil {
		queue = &QueueNode{v, nil}
		sizeOfQueue++
		return true
	}

	t = &QueueNode{v, nil}
	t.Next = queue
	queue = t
	sizeOfQueue++

	return true
}

func Pop(t *QueueNode) (int, bool) {
	if sizeOfQueue == 0 {
		return 0, false
	}

	if sizeOfQueue == 1 {
		queue = nil
		sizeOfQueue--
		return t.Value, true
	}

	temp := t
	for (t.Next) != nil {
		temp = t
		t = t.Next
	}

	v := (temp.Next).Value
	temp.Next = nil

	sizeOfQueue--
	return v, true
}

func traverseQueue(t *QueueNode) {
	if sizeOfQueue == 0 {
		fmt.Println("Empty Queue!")
		return
	}
	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func main() {
	queue = nil
	Push(queue, 10)
	fmt.Println("Size:", sizeOfQueue)
	traverseQueue(queue)

	v, b := Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", sizeOfQueue)

	for i := 0; i < 5; i++ {
		Push(queue, i)
	}
	traverseQueue(queue)
	fmt.Println("Size:", sizeOfQueue)

	v, b = Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", sizeOfQueue)

	v, b = Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", sizeOfQueue)
	traverseQueue(queue)
}
