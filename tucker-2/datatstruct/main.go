package main

import (
	"fmt"

	"github.com/YoungsoonLee/study-ds-algo-golang/tucker-2/datatstruct/src"
)

func main() {
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

	/*
		s := src.NewStack()

		for i := 1; i <= 5; i++ {
			s.Push(i)
		}

		for !s.Empty() {
			v := s.Pop()
			fmt.Printf("%d -> ", v)
		}

		fmt.Println()

		q := src.NewQueue()
		for i := 1; i <= 5; i++ {
			q.Push(i)
		}

		for !q.Empty() {
			v := q.Pop()
			fmt.Printf("%d -> ", v)
		}
	*/

	/*
		ll := &src.LinkedList{}
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
	*/

}
