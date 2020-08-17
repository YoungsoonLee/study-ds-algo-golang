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
	tree := &BinaryTree{}
	tree.Root = &BinaryTreeNode{Val: v}
	return tree
}

func (n *BinaryTreeNode) AddNode(v int) *BinaryTreeNode {
	if n.Val < v {
		if n.Left == nil {
			n.Left = &BinaryTreeNode{Val: v}
			return n.Left
		} else {
			return n.Left.AddNode(v)
		}
	} else {
		if n.Right == nil {
			n.Right = &BinaryTreeNode{Val: v}
			return n.Right
		} else {
			return n.Right.AddNode(v)
		}
	}

}

type depthNode struct {
	depth int
	node  *BinaryTreeNode
}

// BFS는 queue를 사용 해서...
// DFS는 stack을 사용 해서...
func (t *BinaryTree) Print() {
	q := []depthNode{}
	q = append(q, depthNode{depth: 0, node: t.Root})
	currentDepth := 0

	for len(q) > 0 {
		var first depthNode
		first, q = q[0], q[1:]

		if first.depth != currentDepth {
			fmt.Println()
			currentDepth = first.depth
		}

		fmt.Print(first.node.Val, " ")

		if first.node.Left != nil {
			q = append(q, depthNode{depth: currentDepth + 1, node: first.node.Left})
		}

		if first.node.Right != nil {
			q = append(q, depthNode{depth: currentDepth + 1, node: first.node.Right})
		}

	}
}
