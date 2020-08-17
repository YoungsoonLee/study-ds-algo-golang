package src

import "fmt"

type TreeNode struct {
	Val    int
	Childs []*TreeNode
}

type Tree struct {
	Root *TreeNode
}

func (t *TreeNode) AddNode(val int) {
	t.Childs = append(t.Childs, &TreeNode{Val: val})
}

func (t *Tree) AddNode(val int) {
	if t.Root == nil {
		t.Root = &TreeNode{Val: val}
	} else {
		t.Root.Childs = append(t.Root.Childs, &TreeNode{Val: val})
	}
}

func (t *Tree) DFS1() {
	DFS1(t.Root)
}

func DFS1(node *TreeNode) {
	fmt.Printf("%d->", node.Val)

	for i := 0; i < len(node.Childs); i++ {
		DFS1(node.Childs[i])
	}
}

func (t *Tree) DFS2() {
	s := []*TreeNode{}
	s = append(s, t.Root)

	for len(s) > 0 { // loop
		var top *TreeNode
		top, s = s[len(s)-1], s[:len(s)-1]

		fmt.Printf("%d->", top.Val) // 연산

		// for i := 0; i < len(top.Childs); i++ {
		for i := len(top.Childs) - 1; i >= 0; i-- { // 자식 널기
			s = append(s, top.Childs[i])
		}
	}
}

func (t *Tree) BFS() {
	q := []*TreeNode{}
	q = append(q, t.Root)

	for len(q) > 0 { // loop
		var first *TreeNode
		first, q = q[0], q[1:]
		fmt.Printf("%d->", first.Val) // 연산

		for i := 0; i < len(first.Childs); i++ {
			q = append(q, first.Childs[i])
		}
	}
}
