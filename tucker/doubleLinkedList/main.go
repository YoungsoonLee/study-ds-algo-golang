package main

import "fmt"

type Node struct {
	next *Node
	prev *Node
	val  int
}

type LinkedList struct {
	root *Node
	tail *Node
}

func (l *LinkedList) AddNode(val int) {
	node := &Node{val: val}

	if l.root == nil {
		l.root = node
		l.tail = l.root
		return
	}

	l.tail.next = node
	prev := l.tail
	l.tail = l.tail.next
	l.tail.prev = prev
	return
}

func (l *LinkedList) RemoveNode(node *Node) {
	if node == l.root {
		l.root = l.root.next
		l.root.prev = nil
		node.next = nil
		return
	}

	prev := node.prev

	if node == l.tail {
		prev.next = nil
		l.tail.prev = nil
		l.tail = prev
	} else {
		node.prev = nil
		prev.next = prev.next.next
	}

	node.next = nil
}

func (l *LinkedList) PrintNode() {
	node := l.root
	for node.next != nil {
		fmt.Printf("%d ->", node.val)
		node = node.next
	}
	fmt.Printf("%d \n", node.val)
}

func (l *LinkedList) PrintReverse() {
	node := l.tail
	for node.prev != nil {
		fmt.Printf("%d ->", node.val)
		node = node.prev
	}
	fmt.Printf("%d \n", node.val)
}

func main() {
	list := &LinkedList{}
	list.AddNode(0)
	for i := 1; i < 10; i++ {
		list.AddNode(i)
	}
	list.PrintNode()

	list.RemoveNode(list.root.next)
	list.PrintNode()

	list.RemoveNode(list.root)
	list.PrintNode()

	list.RemoveNode(list.tail)
	list.PrintNode()

	fmt.Println(list.tail.val)
	list.PrintNode()

	list.PrintReverse()

}
