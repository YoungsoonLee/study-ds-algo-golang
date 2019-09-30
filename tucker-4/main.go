package main

import (
	"fmt"

	"github.com/YoungsoonLee/study-ds-algo-golang/tucker-4/datastruct"
)

func main() {
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
