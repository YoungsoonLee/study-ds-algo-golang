package main

import "fmt"

type Node struct {
	Next *Node
	Val  int
}

func main() {
	var root, tail *Node

	root = &Node{Val: 0}
	tail = root

	for i := 1; i < 10; i++ {
		tail = AddNode(tail, i)
	}
	PrintNode(root)

	root, tail = RemoveNode(root.Next, root, tail)

	PrintNode(root)

	root, tail = RemoveNode(root, root, tail)

	PrintNode(root)

	root, tail = RemoveNode(tail, root, tail)

	PrintNode(root)
}

func PrintNode(root *Node) {
	node := root
	for node.Next != nil {
		fmt.Printf("%d ->", node.Val)
		node = node.Next
	}
	fmt.Printf("%d\n", node.Val)
}

func AddNode(tail *Node, val int) *Node {
	// var tail *Node
	// tail = root

	// for tail.Next != nil {
	// 	tail = tail.Next
	// }

	node := &Node{Val: val}
	tail.Next = node
	return node
}

func RemoveNode(node *Node, root *Node, tail *Node) (*Node, *Node) {
	if node == root {
		root = root.Next
		if root == nil {
			tail = nil
		}
		return root, tail
	}

	prev := root
	for prev.Next != node {
		prev = prev.Next
	}

	if node == tail {
		prev.Next = nil
		tail = prev
	} else {
		prev.Next = prev.Next.Next
	}

	return root, tail
}
