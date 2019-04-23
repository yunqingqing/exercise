package day52

import (
	"../utils/binarytree"
	"testing"
)

//            1
//         /      \
//        2        3
//       /\         \
//      4  5         6
//        /
//       7
func TestFunc1(t *testing.T) {
	node1 := binarytree.Node{
		Val: 1,
	}
	node2 := binarytree.Node{
		Val: 2,
	}
	node3 := binarytree.Node{
		Val: 3,
	}
	node4 := binarytree.Node{
		Val: 4,
	}
	node5 := binarytree.Node{
		Val: 5,
	}
	node6 := binarytree.Node{
		Val: 6,
	}
	node7 := binarytree.Node{
		Val: 7,
	}

	binarytree.ConnectNode(&node1, &node2, &node3)
	binarytree.ConnectNode(&node2, &node4, &node5)
	binarytree.ConnectNode(&node3, nil, &node6)
	binarytree.ConnectNode(&node5, &node7, nil)

	if TreeDepth(&node1) != 4 {
		t.Errorf("not correct.")
	}
}

//               1
//              /
//             2
//            /
//           3
//          /
//         4
//        /
//       5
func TestFunc2(t *testing.T) {
	node1 := binarytree.Node{
		Val: 1,
	}
	node2 := binarytree.Node{
		Val: 2,
	}
	node3 := binarytree.Node{
		Val: 3,
	}
	node4 := binarytree.Node{
		Val: 4,
	}
	node5 := binarytree.Node{
		Val: 5,
	}

	binarytree.ConnectNode(&node1, &node2, nil)
	binarytree.ConnectNode(&node2, &node3, nil)
	binarytree.ConnectNode(&node3, &node4, nil)
	binarytree.ConnectNode(&node4, &node5, nil)

	if TreeDepth(&node1) != 5 {
		t.Errorf("not correct.")
	}
}

func TestFunc3(t *testing.T) {
	node1 := binarytree.Node{
		Val: 1,
	}

	if TreeDepth(&node1) != 1 {
		t.Errorf("not correct.")
	}
}
