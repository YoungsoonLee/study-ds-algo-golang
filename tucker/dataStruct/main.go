package main

import (
	"fmt"

	"github.com/YoungsoonLee/study-ds-algo-golang/tucker/datatStruct/src"
)

func main() {
	fmt.Println("abcde := ", src.Hash("abcde"))
	fmt.Println("abcde := ", src.Hash("abcde"))

	fmt.Println("abcdf := ", src.Hash("abcdf"))
	fmt.Println("tbcde := ", src.Hash("tbcde"))

	mp := src.CreateMap()
	mp.Add("AAA", "0107777777")
	mp.Add("BBB", "0108888888")
	mp.Add("CCCSKHDKSFHKDF", "0101111111")
	mp.Add("CCC", "97938479347329847")

	fmt.Println(mp.Get("AAA"))
	fmt.Println(mp.Get("CCC"))
	fmt.Println(mp.Get("DDD"))

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

		h.Print()
	*/

	/*
		// heap 문제: [-1,3,-1,5,4], 2번째 큰 값
		h := &src.Heap{}
		nums := []int{-1, 3, -1, 5, 4}
		for i := 0; i < len(nums); i++ {
			h.MinPush(nums[i])
			if h.Count() > 2 {
				h.MinPop()
			}
		}

		fmt.Println(h.MinPop())
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

		// tree.BFS() have a bug
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
		fmt.Println()

		if found, cnt := tree.Search(6); found {
			fmt.Println("found 6, cnt: ", cnt)
		} else {
			fmt.Println("not found 6, cnt: ", cnt)
		}

		if found, cnt := tree.Search(11); found {
			fmt.Println("found 11, cnt: ", cnt)
		} else {
			fmt.Println("not found 11, cnt: ", cnt)
		}
	*/

	/*
		list := &src.LinkedList{}
		list.AddNode(0)

		for i := 1; i < 10; i++ {
			list.AddNode(i)
		}

		list.PrintNode()

		list.RemoveNode(list.Root.Next)

		list.PrintNode()

		list.RemoveNode(list.Root)

		list.PrintNode()

		list.RemoveNode(list.Tail)

		list.PrintNode()

		list.PrintReverse()

		stack2 := src.NewStack()

		for i := 1; i <= 5; i++ {
			stack2.Push(i)
		}

		for !stack2.Empty() {
			val := stack2.Pop()
			fmt.Printf("%d ->", val)
		}

		queue2 := src.NewQueue()
		for i := 1; i <= 5; i++ {
			queue2.Push(i)
		}
		fmt.Println()
		for !queue2.Empty() {
			val := queue2.Pop()
			fmt.Printf("%d ->", val)
		}
	*/
}
