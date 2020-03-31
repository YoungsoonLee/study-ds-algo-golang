package main

type Node struct {
	val  int
	next *Node
	prev *Node
}

type LinkedList struct {
	root *Node
	tail *Node
}

func (l *LinkedList) AddNode(val int) {
	newNode := &Node{val: val}

	if l.root == nil {
		// first
		l.root = newNode
		l.tail = l.root
	}

	l.tail.next = newNode
	prev := l.tail
	l.tail = l.tail.next
	l.tail.prev = prev
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

func main() {

}
