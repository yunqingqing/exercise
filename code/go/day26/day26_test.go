package day26

import (
	"testing"
	"../utils/binarytree"
)

func TestFunc(t *testing.T) {
	node1 := binarytree.Node{
		Val: 8,
	}
	node2 := binarytree.Node{
		Val: 6,
	}
	node3 := binarytree.Node{
		Val: 10,
	}
	node4 := binarytree.Node{
		Val: 5,
	}
	node5 := binarytree.Node{
		Val: 7,
	}
	node6 := binarytree.Node{
		Val: 9,
	}
	node7 := binarytree.Node{
		Val: 11,
	}
	binarytree.ConnectNode(&node1, &node2, &node3)
	binarytree.ConnectNode(&node2, &node4, &node5)
	binarytree.ConnectNode(&node3, &node6, &node7)
	MirrorRecursively(&node1)

	if node1.Left != &node3 {
		t.Errorf("not correct.")
	}
	if node1.Right != &node2 {
		t.Errorf("not correct.")
	}
	if node2.Left != &node5 {
		t.Errorf("not correct.")
	}
	if node2.Right != &node4 {
		t.Errorf("not correct.")
	}
	if node3.Left != &node7 {
		t.Errorf("not correct.")
	}
	if node3.Right != &node6 {
		t.Errorf("not correct.")
	}
}
