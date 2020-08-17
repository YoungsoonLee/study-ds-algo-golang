package datastruct

type TreeNode struct {
	Val    int
	Childs []*TreeNode
}

type Tree struct {
	Root *TreeNode
}

func newNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func (t *Tree) AddNode(val int) {
	if t.Root == nil {
		t.Root = newNode(val)
	} else {
		t.Root.Childs = append(t.Root.Childs, newNode(val))
	}
}
