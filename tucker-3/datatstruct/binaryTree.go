package datatstruct

type Bnode struct {
	Val   int
	Left  *Bnode
	Right *Bnode
}

type BinaryTree struct {
	Root *Bnode
}

func NewBinaryTree(v int) *BinaryTree {
	tree := &BinaryTree{}
	tree.Root = &Bnode{Val: v}
	return tree
}

func (b *Bnode) AddNode(v int) *Bnode {
	newNode := &Bnode{Val: v}

	if b.Val > v {
		// left
		if b.Left == nil {
			b.Left = newNode
			return b.Left
		} else {
			return b.Left.AddNode(v)
		}
	} else {
		//right

		if b.Right == nil {
			b.Right = newNode
			return b.Right
		} else {
			return b.Right.AddNode(v)
		}
	}
}

func main() {

}
