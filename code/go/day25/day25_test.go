package day25

import "testing"

func connectNode(root, left, right *Node) {
	root.left = left
	root.right = right
}

func TestFunc1(t *testing.T) {
	node1 := Node{
		val: 8,
	}
	node2 := Node{
		val: 8,
	}
	node3 := Node{
		val: 7,
	}
	node4 := Node{
		val: 9,
	}
	node5 := Node{
		val: 2,
	}
	node6 := Node{
		val: 4,
	}
	node7 := Node{
		val: 7,
	}
	connectNode(&node1, &node2, &node3)
	connectNode(&node2, &node4, &node5)
	connectNode(&node5, &node6, &node7)

	// tree2
	node8 := Node{
		val: 8,
	}
	node9 := Node{
		val: 9,
	}
	node10 := Node{
		val: 2,
	}
	connectNode(&node8, &node9, &node10)

	if HasSubtree(&node1, &node8) != true {
		t.Errorf("not correct.")
	}
}


func TestFunc2(t *testing.T) {
	node1 := Node{
		val: 8,
	}
	node2 := Node{
		val: 8,
	}
	node3 := Node{
		val: 7,
	}
	node4 := Node{
		val: 9,
	}
	node5 := Node{
		val: 3,
	}
	node6 := Node{
		val: 4,
	}
	node7 := Node{
		val: 7,
	}
	connectNode(&node1, &node2, &node3)
	connectNode(&node2, &node4, &node5)
	connectNode(&node5, &node6, &node7)
	// tree2
	node8 := Node{
		val: 8,
	}
	node9 := Node{
		val: 9,
	}
	node10 := Node{
		val: 2,
	}
	connectNode(&node8, &node9, &node10)

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
	node1 := Node{
		val: 8,
	}
	node2 := Node{
		val: 8,
	}
	node3 := Node{
		val: 9,
	}
	node4 := Node{
		val: 2,
	}
	node5 := Node{
		val: 5,
	}
	connectNode(&node1, &node2, nil)
	connectNode(&node2, &node3, nil)
	connectNode(&node3, &node4, nil)
	connectNode(&node4, &node5, nil)

	node8 := Node{
		val: 8,
	}
	node9 := Node{
		val: 9,
	}
	node10 := Node{
		val: 2,
	}
	connectNode(&node8, &node9, nil)
	connectNode(&node9, &node10, nil)

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
	node1 := Node{
		val: 8,
	}
	node2 := Node{
		val: 8,
	}
	node3 := Node{
		val: 9,
	}
	node4 := Node{
		val: 3,
	}
	node5 := Node{
		val: 5,
	}
	connectNode(&node1, &node2, nil)
	connectNode(&node2, &node3, nil)
	connectNode(&node3, &node4, nil)
	connectNode(&node4, &node5, nil)

	node8 := Node{
		val: 8,
	}
	node9 := Node{
		val: 9,
	}
	node10 := Node{
		val: 2,
	}
	connectNode(&node8, &node9, nil)
	connectNode(&node9, &node10, nil)

	if HasSubtree(&node1, &node8) != false {
		t.Errorf("not correct.")
	}
}