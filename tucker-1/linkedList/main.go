package main

import "fmt"

type Node struct {
	next *Node
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
	l.tail = l.tail.next
	return
}

func (l *LinkedList) RemoveNode(node *Node) {
	if node == l.root {
		l.root = l.root.next
		node.next = nil
		return
	}

	prev := l.root
	for prev.next != node {
		prev = prev.next
	}

	if node == l.tail {
		prev.next = nil
		l.tail = prev
	} else {
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

/*
func AddNode(tail *Node, val int) *Node {
	--*
		var tail *Node
		tail = root
		for tail.next != nil {
			tail = tail.next
		}
	*--

	node := &Node{val: val}
	tail.next = node

	return node
}
*/

/*
func RemoveNode(node *Node, root *Node, tail *Node) (*Node, *Node) {
	if node == root {
		root = root.next
		if root == nil {
			tail = nil
		}
		return root, tail
	}

	prev := root
	for prev.next != node {
		prev = prev.next
	}

	if node == tail {
		prev.next = nil
		tail = prev
	} else {
		// middle
		prev.next = prev.next.next
	}

	return root, tail
}
*/

/*
func PrintNode(root *Node) {
	node := root
	for node.next != nil {
		fmt.Printf("%d ->", node.val)
		node = node.next
	}
	fmt.Printf("%d ->", node.val)
}
*/

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

	/*
		var root *Node
		var tail *Node

		root = &Node{val: 0}
		tail = root

		for i := 1; i < 10; i++ {
			tail = AddNode(tail, i)
		}

		PrintNode(root)

		//node = root
		root, tail = RemoveNode(root.next, root, tail)
		fmt.Println()
		PrintNode(root)

		root, tail = RemoveNode(root, root, tail)
		fmt.Println()
		PrintNode(root)

		root, tail = RemoveNode(tail, root, tail)
		fmt.Println()
		PrintNode(root)

		fmt.Println()
		fmt.Println(tail.val)
	*/
}
