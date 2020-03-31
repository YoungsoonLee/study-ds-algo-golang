package src

import "fmt"

type Node struct {
	Next *Node
	Prev *Node
	Val  int
}

type LinkedList struct {
	Root *Node
	Tail *Node
}

func (l *LinkedList) AddNode(val int) {
	node := &Node{Val: val}

	if l.Root == nil {
		l.Root = node
		l.Tail = l.Root
		return
	}

	l.Tail.Next = node
	prev := l.Tail
	l.Tail = l.Tail.Next
	l.Tail.Prev = prev
	return
}

func (l *LinkedList) RemoveNode(node *Node) {
	if node == l.Root {
		l.Root = l.Root.Next
		l.Root.Prev = nil
		node.Next = nil
		return
	}

	prev := node.Prev

	if node == l.Tail {
		prev.Next = nil
		l.Tail.Prev = nil
		l.Tail = prev
	} else {
		node.Prev = nil
		prev.Next = prev.Next.Next
	}

	node.Next = nil
}

func (l *LinkedList) Back() int {
	if l.Tail != nil {
		return l.Tail.Val
	}
	return 0
}

func (l *LinkedList) PopBack() {
	if l.Tail == nil {
		return
	}
	l.RemoveNode(l.Tail)
}

func (l *LinkedList) PrintNode() {
	node := l.Root
	for node.Next != nil {
		fmt.Printf("%d ->", node.Val)
		node = node.Next
	}
	fmt.Printf("%d \n", node.Val)
}

func (l *LinkedList) PrintReverse() {
	node := l.Tail
	for node.Prev != nil {
		fmt.Printf("%d ->", node.Val)
		node = node.Prev
	}
	fmt.Printf("%d \n", node.Val)
}
