package main

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

}

func (l *LinkedList) RemoveNode(node *Node) {
	if node == l.Root {
		l.Root = l.Root.Next
		l.Root.Prev = nil
		node.Next = nil

		if l.Root == nil {
			l.Tail = nil
		}
		return
	}

	/*
		prev := l.Root
		for prev.Next != node {
			prev = prev.Next
		}
	*/
	prev := node.Prev

	if node == l.Tail {
		prev.Next = nil
		l.Tail.Prev = nil
		l.Tail = prev
	} else {
		node.Prev = nil
		prev.Next = prev.Next.Next
		prev.Next.Prev = prev // !!!
	}
	node.Next = nil

}

func (l *LinkedList) PrintNode() {
	node := l.Root
	for node.Next != nil {
		fmt.Printf("%d ->", node.Val)
		node = node.Next
	}
	fmt.Printf("%d\n", node.Val)
}

func (l *LinkedList) PrintRever() {
	node := l.Tail
	for node.Prev != nil {
		fmt.Printf("%d ->", node.Val)
		node = node.Prev
	}
	fmt.Printf("%d\n", node.Val)
}

func main() {
	ll := &LinkedList{}
	ll.AddNode(0)

	for i := 1; i < 10; i++ {
		ll.AddNode(i)
	}

	ll.PrintNode()

	ll.RemoveNode(ll.Root.Next)
	ll.PrintNode()

	ll.RemoveNode(ll.Root)
	ll.PrintNode()

	ll.RemoveNode(ll.Tail)
	ll.PrintNode()

	ll.PrintRever()

	/*
		var root *Node
		var tail *Node

		root = &Node{val: 0}
		tail = root

		for i := 1; i < 10; i++ {
			tail = AddNode(tail, i)
		}

		//node := root
		PrintNode(root)
		root, tail = RemoveNode(root.next, root, tail)
		PrintNode(root)

		root, tail = RemoveNode(root, root, tail)
		PrintNode(root)

		root, tail = RemoveNode(tail, root, tail)
		PrintNode(root)
	*/

}
