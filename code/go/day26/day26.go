package day26

import "../utils/binarytree"

func MirrorRecursively(node *binarytree.Node) {
	if node.Left == nil && node.Right == nil {
		return
	}

	node.Left, node.Right = node.Right, node.Left

	if node.Left != nil {
		MirrorRecursively(node.Left)
	}

	if node.Right != nil {
		MirrorRecursively(node.Right)
	}
}
