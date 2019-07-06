package main

import (
	"fmt"
	"io"
	"os"
)

type BinaryTree struct {
	root *BinaryNode
}

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int
}

func (t *BinaryTree) add(data int) *BinaryTree {
	if t.root == nil {
		t.root = &BinaryNode{left: nil, right: nil, data: data}
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *BinaryNode) insert(data int) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &BinaryNode{left: nil, right: nil, data: data}
		} else {
			n.left.insert(data)
		}
	} else if data >= n.data {
		if n.right == nil {
			n.right = &BinaryNode{left: nil, right: nil, data: data}
		} else {
			n.right.insert(data)
		}
	}
}

func (t *BinaryTree) remove(data int) {
	delete(t.root, data)
	/*
		if node != nil {
			t.root = node
		}
	*/
}

// !!!!!!
func delete(node *BinaryNode, data int) *BinaryNode {
	if node == nil {
		return nil
	}

	if data < node.data {
		node.left = delete(node.left, data)
	} else if data > node.data {
		node.right = delete(node.right, data)
	} else {
		fmt.Println(node, node.left, node.right)
		if node.left == nil {
			newNode := node.right
			return newNode
		} else if node.right == nil {
			newNode := node.left
			return newNode
		} else {
			fmt.Println("change: ", node, node.left, node.right)
			ex := node.left
			for ex.right != nil {
				ex = ex.right
			}
			temp := node.data
			node.data = ex.data
			ex.data = temp
			node.left = delete(node.left, ex.data)
		}
	}
	return node
}

func get(node *BinaryNode, data int) *BinaryNode {
	if node == nil {
		return nil
	}
	fmt.Println(node.data)
	if node.data < data {
		fmt.Println("right: ", node.right)
		get(node.right, data)
	} else if node.data > data {
		fmt.Println("left: ", node.left)
		get(node.left, data)
	}

	return node
}

func print(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}

func main() {
	tree := &BinaryTree{}
	tree.add(100).
		add(-20).
		add(-50).
		add(-15).
		add(-60).
		add(50).
		add(60).
		add(55).
		add(85).
		add(15).
		add(5).
		add(-10)
	print(os.Stdout, tree.root, 0, 'M')
	//get(tree.root, 5)
	tree.remove(-20)
	print(os.Stdout, tree.root, 0, 'M')
	//print(os.Stdout, tree.get(5), 0, 'M')
	//tree.remove(100)
	//print(os.Stdout, tree.root, 0, 'M')
}
