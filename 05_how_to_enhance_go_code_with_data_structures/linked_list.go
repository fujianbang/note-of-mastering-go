package main

import "fmt"

type LinkedListNode struct {
	Value int
	Next  *LinkedListNode
}

var root = new(LinkedListNode)

func addLinkedListNode(t *LinkedListNode, v int) int {
	if root == nil {
		t = &LinkedListNode{v, nil}
		root = t
		return 0
	}

	if v == t.Value {
		fmt.Println("LinkedListNode already exists:", v)
		return -1
	}

	if t.Next == nil {
		t.Next = &LinkedListNode{v, nil}
		return -2
	}

	return addLinkedListNode(t.Next, v)
}

func traverseLinkedList(t *LinkedListNode) {
	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func lookupNode(t *LinkedListNode, v int) bool {
	if root == nil {
		t = &LinkedListNode{v, nil}
		return false
	}

	if v == t.Value {
		return true
	}

	if t.Next == nil {
		return false
	}

	return lookupNode(t.Next, v)
}

func size(t *LinkedListNode) int {
	if t == nil {
		fmt.Println("-> Empty list!")
		return 0
	}
	i := 0
	for t != nil {
		i++
		t = t.Next
	}
	return i
}

func main() {
	fmt.Println(root)
	root = nil
	traverseLinkedList(root)
	addLinkedListNode(root, 1)
	addLinkedListNode(root, -1)
	traverseLinkedList(root)
	addLinkedListNode(root, 10)
	addLinkedListNode(root, 5)
	addLinkedListNode(root, 45)
	addLinkedListNode(root, 5)
	addLinkedListNode(root, 5)
	traverseLinkedList(root)
	addLinkedListNode(root, 100)
	traverseLinkedList(root)

	if lookupNode(root, 100) {
		fmt.Println("LinkedListNode exists!")
	} else {
		fmt.Println("LinkedListNode does not exist!")
	}

	if lookupNode(root, -100) {
		fmt.Println("LinkedListNode exists!")
	} else {
		fmt.Println("LinkedListNode does not exist!")
	}
}
