package main

import "fmt"

type StackNode struct {
	Value int
	Next  *StackNode
}

var sizeOfStack = 0
var stack = new(StackNode)

func PushStack(v int) bool {
	if stack == nil {
		stack = &StackNode{v, nil}
		sizeOfStack = 1
		return true
	}

	temp := &StackNode{v, nil}
	temp.Next = stack
	stack = temp
	sizeOfStack++
	return true
}

func PopStack(t *StackNode) (int, bool) {
	if sizeOfStack == 0 {
		return 0, false
	}

	if sizeOfStack == 1 {
		sizeOfStack = 0
		stack = nil
		return t.Value, true
	}

	stack = stack.Next
	sizeOfStack--
	return t.Value, true
}

func traverseStack(t *StackNode) {
	if sizeOfStack == 0 {
		fmt.Println("Empty Stack!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func main() {
	stack = nil
	v, b := PopStack(stack)
	if b {
		fmt.Print(v, " ")
	} else {
		fmt.Println("Pop() failed!")
	}

	PushStack(100)
	traverseStack(stack)
	PushStack(200)
	traverseStack(stack)

	for i := 0; i < 10; i++ {
		PushStack(i)
	}

	for i := 0; i < 15; i++ {
		v, b := PopStack(stack)
		if b {
			fmt.Print(v, " ")
		} else {
			break
		}
	}
	fmt.Println()
	traverseStack(stack)
}
