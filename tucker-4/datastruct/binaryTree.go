package datastruct

import "fmt"

type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

type BinaryTree struct {
	Root *BinaryTreeNode
}

func NewBinaryTree(v int) *BinaryTree {
	tree := &BinaryTree{Root: &BinaryTreeNode{}}
	tree.Root = &BinaryTreeNode{Val: v}
	return tree
}

func (n *BinaryTreeNode) AddNode(val int) *BinaryTreeNode {
	newNode := &BinaryTreeNode{Val: val}

	if n.Val < val {
		if n.Right == nil {
			n.Right = newNode
			return n.Right
		} else {
			return n.Right.AddNode(val)
		}
	} else {
		if n.Left == nil {
			n.Left = newNode
			return n.Left
		} else {
			return n.Left.AddNode(val)
		}
	}
}

type DepthNode struct {
	depth int
	node  *BinaryTreeNode
}

func (t *BinaryTree) Print() {
	q := []DepthNode{}
	q = append(q, DepthNode{depth: 0, node: t.Root})
	curDepth := 0

	for len(q) > 0 {
		var first DepthNode
		first, q = q[0], q[1:]

		if first.depth != curDepth {
			fmt.Println()
			curDepth = first.depth
		}

		fmt.Print(first.node.Val, " ")

		if first.node.Left != nil {
			q = append(q, DepthNode{depth: curDepth + 1, node: first.node.Left})
		}
		if first.node.Right != nil {
			q = append(q, DepthNode{depth: curDepth + 1, node: first.node.Right})
		}
	}
}

func (t *BinaryTree) Search(val int) (bool, int) {
	return t.Root.Search(val, 1)
}

func (t *BinaryTreeNode) Search(val, cnt int) (bool, int) {
	if t.Val == val {
		return true, cnt
	} else if t.Val > val {
		if t.Left != nil {
			return t.Left.Search(val, cnt+1)
		}
		return false, cnt
	} else {
		if t.Right != nil {
			return t.Right.Search(val, cnt+1)
		}
		return false, cnt
	}

}
