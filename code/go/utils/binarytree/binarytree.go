package binarytree

type Node struct {
	Val   int
	Right *Node
	Left  *Node
}

func ConnectNode(root, left, right *Node) {
	root.Left = left
	root.Right = right
}