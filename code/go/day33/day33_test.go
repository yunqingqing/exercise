package day33

import (
	"../utils/binarytree"
	"testing"
)

func TestFunc(t *testing.T) {
	node1 := binarytree.Node{
		Val: 10,
	}
	node2 := binarytree.Node{
		Val: 5,
	}
	node3 := binarytree.Node{
		Val: 12,
	}
	node4 := binarytree.Node{
		Val: 4,
	}
	node5 := binarytree.Node{
		Val: 7,
	}

	binarytree.ConnectNode(&node1, &node2, &node3)
	binarytree.ConnectNode(&node2, &node4, &node5)

	FindPath(&node1, 22)
}
