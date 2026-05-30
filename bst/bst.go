package bst

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BST struct {
	Root *Node
}

func NewBST() *BST {
	return &BST{Root: nil}
}

func (b *BST) Insert(v int) {
	n := &Node{Value: v, Left: nil, Right: nil}
	if b.Root == nil {
		b.Root = n
		return
	}
	curr := b.Root
	for true {
		if n.Value == curr.Value {
			return
		}
		if n.Value < curr.Value {
			if curr.Left == nil {
				curr.Left = n
				return
			} else {
				curr = curr.Left
			}
		}
		if n.Value > curr.Value {
			if curr.Right == nil {
				curr.Right = n
				return
			} else {
				curr = curr.Right
			}
		}
	}
}
