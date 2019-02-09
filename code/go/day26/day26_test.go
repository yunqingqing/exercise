package day26

import (
	"fmt"
	"testing"
)

// TODO: refact, create binarytree package.
func connectNode(root, left, right *Node) {
	root.left = left
	root.right = right
}

func printTree(root *Node) {
	fmt.Println(root.val)
}

func TestFunc(t *testing.T) {
	node1 := Node{
		val: 8,
	}
	node2 := Node{
		val: 6,
	}
	node3 := Node{
		val: 10,
	}
	node4 := Node{
		val: 5,
	}
	node5 := Node{
		val: 7,
	}
	node6 := Node{
		val: 9,
	}
	node7 := Node{
		val: 11,
	}
	connectNode(&node1, &node2, &node3)
	connectNode(&node2, &node4, &node5)
	connectNode(&node3, &node6, &node7)
	MirrorRecursively(&node1)

	if node1.left != &node3 {
		t.Errorf("not correct.")
	}
	if node1.right != &node2 {
		t.Errorf("not correct.")
	}
	if node2.left != &node5 {
		t.Errorf("not correct.")
	}
	if node2.right != &node4 {
		t.Errorf("not correct.")
	}
	if node3.left != &node7 {
		t.Errorf("not correct.")
	}
	if node3.right != &node6 {
		t.Errorf("not correct.")
	}
}
