package day27

import (
	"../utils/binarytree"
	"testing"
)

//            8
//        6      6
//       5 7    7 5
func TestFunc1(t *testing.T) {
	node1 := binarytree.Node{
		Val: 8,
	}
	node2 := binarytree.Node{
		Val: 6,
	}
	node3 := binarytree.Node{
		Val: 6,
	}
	node4 := binarytree.Node{
		Val: 5,
	}
	node5 := binarytree.Node{
		Val: 7,
	}
	node6 := binarytree.Node{
		Val: 7,
	}
	node7 := binarytree.Node{
		Val: 5,
	}
	binarytree.ConnectNode(&node1, &node2, &node3)
	binarytree.ConnectNode(&node2, &node4, &node5)
	binarytree.ConnectNode(&node3, &node6, &node7)

	if IsSymmetrical(&node1) != true {
		t.Errorf("not correct.")
	}
}

//            8
//        6      9
//       5 7    7 5
func TestFunc2(t *testing.T) {
	node1 := binarytree.Node{
		Val: 8,
	}
	node2 := binarytree.Node{
		Val: 6,
	}
	node3 := binarytree.Node{
		Val: 9,
	}
	node4 := binarytree.Node{
		Val: 5,
	}
	node5 := binarytree.Node{
		Val: 7,
	}
	node6 := binarytree.Node{
		Val: 7,
	}
	node7 := binarytree.Node{
		Val: 5,
	}
	binarytree.ConnectNode(&node1, &node2, &node3)
	binarytree.ConnectNode(&node2, &node4, &node5)
	binarytree.ConnectNode(&node3, &node6, &node7)

	if IsSymmetrical(&node1) != false {
		t.Errorf("not correct.")
	}
}

// 所有结点都有相同的值，树不对称
//               5
//              / \
//             5   5
//            /     \
//           5       5
//          /       /
//         5       5
func TestFunc3(t *testing.T) {
	node1 := binarytree.Node{
		Val: 5,
	}
	node2 := binarytree.Node{
		Val: 5,
	}
	node3 := binarytree.Node{
		Val: 5,
	}
	node4 := binarytree.Node{
		Val: 5,
	}
	node5 := binarytree.Node{
		Val: 5,
	}
	node6 := binarytree.Node{
		Val: 5,
	}
	node7 := binarytree.Node{
		Val: 5,
	}
	binarytree.ConnectNode(&node1, &node2, &node3)
	binarytree.ConnectNode(&node2, &node4, nil)
	binarytree.ConnectNode(&node3, nil, &node5)
	binarytree.ConnectNode(&node4, &node6, nil)
	binarytree.ConnectNode(&node5, &node7, nil)

	if IsSymmetrical(&node1) != false {
		t.Errorf("not correct.")
	}
}
