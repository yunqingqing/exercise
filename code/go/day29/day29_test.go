package day29

import (
	"../utils/binarytree"
	"testing"
)

//            8
//        6      7
//       3 4    1 2
func TestFunc1(t *testing.T) {
	node1 := binarytree.Node{
		Val: 8,
	}
	node2 := binarytree.Node{
		Val: 6,
	}
	node3 := binarytree.Node{
		Val: 7,
	}
	node4 := binarytree.Node{
		Val: 3,
	}
	node5 := binarytree.Node{
		Val: 4,
	}
	node6 := binarytree.Node{
		Val: 1,
	}
	node7 := binarytree.Node{
		Val: 2,
	}
	binarytree.ConnectNode(&node1, &node2, &node3)
	binarytree.ConnectNode(&node2, &node4, &node5)
	binarytree.ConnectNode(&node3, &node6, &node7)

	PrintFromTopToBottom(&node1)
}
