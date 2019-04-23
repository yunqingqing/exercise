package day52

import (
	"../utils/binarytree"
)

func TreeDepth(node *binarytree.Node) int {
	if node == nil {
		return 0
	}

	nLeft := TreeDepth(node.Left)
	nRight := TreeDepth(node.Right)

	if nLeft > nRight {
		return nLeft + 1
	} else {
		return nRight + 1
	}
}
