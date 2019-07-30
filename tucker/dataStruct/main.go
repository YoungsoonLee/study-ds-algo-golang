package main

import (
	"fmt"

	"github.com/YoungsoonLee/study-ds-algo-golang/tucker/dataStruct/src"
)

func main() {
	fmt.Println("abcde := ", src.Hash("abcde"))

	m := src.CreateMap()
	m.Add("AAA", "0107777777")
	m.Add("BBB", "0108888888")
	m.Add("CKHDKHSKDHUIY", "0111111111")
	m.Add("CCC", "00381038038 ")

	fmt.Println(m.Get("AAA"))

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
}
