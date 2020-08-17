package datastruct

import "fmt"

type TreeNode struct {
	Val    int
	Childs []*TreeNode
}

// !!!!!
type Tree struct {
	Root *TreeNode
}

func (t *TreeNode) AddNode(val int) {
	newTree := &TreeNode{Val: val}
	t.Childs = append(t.Childs, newTree)
}

func (t *Tree) AddNode(val int) {
	newTree := &TreeNode{Val: val}
	if t.Root == nil {
		t.Root = newTree
	} else {
		t.Root.Childs = append(t.Root.Childs, newTree)
	}
}

func (t *Tree) DFS1() {
	DFS1(t.Root)
}

// !!!
func DFS1(node *TreeNode) {
	fmt.Printf("%d ->", node.Val)

	for i := 0; i < len(node.Childs); i++ {
		DFS1(node.Childs[i])
	}
}

func (t *Tree) DFS2() {
	s := []*TreeNode{}
	s = append(s, t.Root)

	for len(s) > 0 {
		var last *TreeNode
		last, s = s[len(s)-1], s[:len(s)-1]

		fmt.Printf("%d ->", last.Val)

		for i := len(last.Childs) - 1; i >= 0; i-- {
			s = append(s, last.Childs[i])
		}
	}
}

func (t *Tree) BFS() {
	q := []*TreeNode{}
	q = append(q, t.Root)

	for len(q) > 0 {
		var first *TreeNode
		first, q = q[0], q[1:]

		// 연산
		fmt.Printf("%d ->", first.Val)

		for i := 0; i < len(first.Childs); i++ {
			q = append(q, first.Childs[i])
		}
	}
}
