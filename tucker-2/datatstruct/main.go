package main

import (
	"fmt"
	"sort"

	"github.com/YoungsoonLee/study-ds-algo-golang/tucker-2/datatstruct/src"
)

func main() {
	fmt.Println("abcde = ", src.Hash("abcde"))

	a := []int{-1, 3, -1, 5, 4}
	sort.Ints(a)
	fmt.Println(a)

	sort.Sort(sort.Reverse(sort.IntSlice(a))) // !!!
	fmt.Println(a)

	/*
		h := &src.Heap{}

		nums := []int{-1, 3, -1, 5, 4}
		for i := 0; i < len(nums); i++ {
			h.Push(nums[i])
			if h.Count() > 2 {
				h.Pop()
			}
		}
		fmt.Println(h.Pop())
	*/

	/*
		h := &src.Heap{}

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
	*/

	/*
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
	*/

	/*
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
