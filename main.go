package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

var size = 0
var stack = new(Node)

func main() {
	fmt.Println("Stack data structure")
}

func push(v int) bool {
	if stack == nil {
		stack = &Node{v, nil}
		size = 1
		return true
	}

	temp := &Node{v, nil}
	temp.Next = stack
	stack = temp
	size++
	return true
}

func pop(t *Node) (int, bool) {
	if size == 0 {
		return 0, false
	}

	if size == 1 {
		size = 0
		stack = nil
		return t.Value, true
	}

	stack = stack.Next
	size--
	return t.Value, true
}

func traverse(t *Node) bool {
	if size == 0 {
		fmt.Println("Empty Stack!")
		return false
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
	return true
}
