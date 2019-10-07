package main

import (
	"fmt"

	"github.com/YoungsoonLee/study-ds-algo-golang/tucker-4/datastruct"
)

func main() {
	h := &datastruct.Heap{}
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

	/*
		tree := datastruct.NewBinaryTree(5)
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
			fmt.Println("found 6, ", cnt)
		} else {
			fmt.Println("not found 6, ", cnt)
		}

		fmt.Println()

		if found, cnt := tree.Search(11); found {
			fmt.Println("found 11, ", cnt)
		} else {
			fmt.Println("not found 11, ", cnt)
		}
	*/

	/*
		tree := datastruct.Tree{}

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

	/*
		stack2 := datastruct.NewStack()

		for i := 1; i <= 5; i++ {
			stack2.Push(i)
		}

		for !stack2.Empty() {
			val := stack2.Pop()
			fmt.Printf("%d -> ", val)
		}
	*/

	/*
		q2 := datastruct.NewStack()

		for i := 1; i <= 5; i++ {
			q2.Push(i)
		}

		for !q2.Empty() {
			val := q2.Pop()
			fmt.Printf("%d -> ", val)
		}
	*/

	/* Linked List
	ls := &datastruct.LinkedList{}
	ls.AddNode(0)

	for i := 1; i < 10; i++ {
		ls.AddNode(i)
	}

	ls.PrintNode()

	ls.RemoveNode(ls.Root.Next)

	ls.PrintNode()

	ls.RemoveNode(ls.Root)

	ls.PrintNode()

	ls.RemoveNode(ls.Tail)

	ls.PrintNode()

	fmt.Printf("%d\n", ls.Tail.Val)

	ls.PrintReverse()
	*/

}
