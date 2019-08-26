package datatstruct

// double linked list
type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

type LinkedList struct {
	Root *Node
	Tail *Node
}

func (ll *LinkedList) AddNode(val int) {
	newNode := &Node{Val: val}

	if ll.Root == nil {
		ll.Root = newNode
		ll.Tail = ll.Root
	} else {
		ll.Tail.Next = newNode
		prev := ll.Tail
		ll.Tail = ll.Tail.Next
		ll.Tail.Prev = prev
	}
}

func (ll *LinkedList) RemoveNode(node *Node) {
	if ll.Root == node {
		ll.Root = ll.Root.Next

		if ll.Root != nil {
			ll.Root.Prev = nil
		}

		node.Next = nil

		if ll.Root == nil {
			ll.Tail = nil
		}
		return
	}

	prev := node.Prev

	if node == ll.Tail {
		prev.Next = nil
		ll.Tail.Prev = nil
		ll.Tail = prev
	} else {
		node.Prev = nil
		prev.Next = prev.Next.Next
		prev.Next.Prev = prev
	}
	node.Next = nil
}
