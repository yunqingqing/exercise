package day25

import (
	"testing"
	"../utils/binarytree"
)

func TestFunc1(t *testing.T) {
	node1 := binarytree.Node{
		Val: 8,
	}
	node2 := binarytree.Node{
		Val: 8,
	}
	node3 := binarytree.Node{
		Val: 7,
	}
	node4 := binarytree.Node{
		Val: 9,
	}
	node5 := binarytree.Node{
		Val: 2,
	}
	node6 := binarytree.Node{
		Val: 4,
	}
	node7 := binarytree.Node{
		Val: 7,
	}
	binarytree.ConnectNode(&node1, &node2, &node3)
	binarytree.ConnectNode(&node2, &node4, &node5)
	binarytree.ConnectNode(&node5, &node6, &node7)

	// tree2
	node8 := binarytree.Node{
		Val: 8,
	}
	node9 := binarytree.Node{
		Val: 9,
	}
	node10 := binarytree.Node{
		Val: 2,
	}
	binarytree.ConnectNode(&node8, &node9, &node10)

	if HasSubtree(&node1, &node8) != true {
		t.Errorf("not correct.")
	}
}


func TestFunc2(t *testing.T) {
	node1 := binarytree.Node{
		Val: 8,
	}
	node2 := binarytree.Node{
		Val: 8,
	}
	node3 := binarytree.Node{
		Val: 7,
	}
	node4 := binarytree.Node{
		Val: 9,
	}
	node5 := binarytree.Node{
		Val: 3,
	}
	node6 := binarytree.Node{
		Val: 4,
	}
	node7 := binarytree.Node{
		Val: 7,
	}
	binarytree.ConnectNode(&node1, &node2, &node3)
	binarytree.ConnectNode(&node2, &node4, &node5)
	binarytree.ConnectNode(&node5, &node6, &node7)
	// tree2
	node8 := binarytree.Node{
		Val: 8,
	}
	node9 := binarytree.Node{
		Val: 9,
	}
	node10 := binarytree.Node{
		Val: 2,
	}
	binarytree.ConnectNode(&node8, &node9, &node10)

	if HasSubtree(&node1, &node8) != false {
		t.Errorf("not correct.")
	}
}


//            树A               树B
//             8                8
//            /                /
//           8                9
//          /                /
//         9                2
//        /
//       2
//      /
//     5
func TestFunc3(t *testing.T) {
	node1 := binarytree.Node{
		Val: 8,
	}
	node2 := binarytree.Node{
		Val: 8,
	}
	node3 := binarytree.Node{
		Val: 9,
	}
	node4 := binarytree.Node{
		Val: 2,
	}
	node5 := binarytree.Node{
		Val: 5,
	}
	binarytree.ConnectNode(&node1, &node2, nil)
	binarytree.ConnectNode(&node2, &node3, nil)
	binarytree.ConnectNode(&node3, &node4, nil)
	binarytree.ConnectNode(&node4, &node5, nil)

	node8 := binarytree.Node{
		Val: 8,
	}
	node9 := binarytree.Node{
		Val: 9,
	}
	node10 := binarytree.Node{
		Val: 2,
	}
	binarytree.ConnectNode(&node8, &node9, nil)
	binarytree.ConnectNode(&node9, &node10, nil)

	if HasSubtree(&node1, &node8) != true {
		t.Errorf("not correct.")
	}
}


//            树A               树B
//             8                8
//            /                /
//           8                9
//          /                /
//         9                2
//        /
//       3
//      /
//     5
func TestFunc4(t *testing.T) {
	node1 := binarytree.Node{
		Val: 8,
	}
	node2 := binarytree.Node{
		Val: 8,
	}
	node3 := binarytree.Node{
		Val: 9,
	}
	node4 := binarytree.Node{
		Val: 3,
	}
	node5 := binarytree.Node{
		Val: 5,
	}
	binarytree.ConnectNode(&node1, &node2, nil)
	binarytree.ConnectNode(&node2, &node3, nil)
	binarytree.ConnectNode(&node3, &node4, nil)
	binarytree.ConnectNode(&node4, &node5, nil)

	node8 := binarytree.Node{
		Val: 8,
	}
	node9 := binarytree.Node{
		Val: 9,
	}
	node10 := binarytree.Node{
		Val: 2,
	}
	binarytree.ConnectNode(&node8, &node9, nil)
	binarytree.ConnectNode(&node9, &node10, nil)

	if HasSubtree(&node1, &node8) != false {
		t.Errorf("not correct.")
	}
}