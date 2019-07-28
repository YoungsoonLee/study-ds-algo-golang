package main

import (
	"fmt"

	"github.com/YoungsoonLee/study-ds-algo-golang/tucker/dataStruct/src"
)

func main() {
	// heap
	h := src.Heap{}

	h.Push(2)
	h.Push(6)
	h.Push(9)
	h.Push(6)
	h.Push(7)
	h.Push(8)

	h.Print()

	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())

	/* binarytree
	tree := src.NewBinaryTree(5)

	tree.Root.AddNode(3)
	tree.Root.AddNode(2)
	tree.Root.AddNode(4)
	tree.Root.AddNode(8)
	tree.Root.AddNode(7)
	tree.Root.AddNode(6)
	tree.Root.AddNode(10)
	tree.Root.AddNode(9)

	tree.Print()
	fmt.Println()

	if found, cnt := tree.Search(6); found {
		fmt.Println("found 6 ", cnt)
	} else {
		fmt.Println("not found 6 ", cnt)
	}

	if found, cnt := tree.Search(11); found {
		fmt.Println("found 11 ", cnt)
	} else {
		fmt.Println("not found 11 ", cnt)
	}
	*/

	/*
		// Tree
		tree := src.Tree{}

		val := 1
		tree.AddNode(val)
		val++

		for i := 0; i < 3; i++ {
			tree.Root.AddNode(val)
			val++
		}

		for i := 0; i < len(tree.Root.Childs); i++ {
			for j := 0; j < 2; j++ {
				tree.Root.Childs[i].AddNode(val)
				val++
			}
		}

		tree.DFS1()
		fmt.Println()
		tree.DFS2()
		fmt.Println()

		tree.BFS()
	*/
}
