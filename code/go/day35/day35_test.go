package day35

import (
	"testing"
	"fmt"
	"../utils/binarytree"
)

func printLinkedList(headNode *binarytree.Node) {
	node := headNode

	fmt.Printf("The node from left to right:\n")
	for node != nil {
		fmt.Printf("%d\t", node.Val)
		if node.Right == nil {
			break
		}
		node = node.Right
	}
	fmt.Printf("\n")

	fmt.Printf("The node from right to left:\n")
	for node != nil {
		fmt.Printf("%d\t", node.Val)
		if node.Left == nil {
			break
		}
		node = node.Left
	}
	fmt.Printf("\n")
}

func TestFunc(t *testing.T) {
	node1 := binarytree.Node{
		Val: 10,
	}
	node2 := binarytree.Node{
		Val: 6,
	}
	node3 := binarytree.Node{
		Val: 14,
	}
	node4 := binarytree.Node{
		Val: 4,
	}
	node5 := binarytree.Node{
		Val: 8,
	}
	node6 := binarytree.Node{
		Val: 12,
	}
	node7 := binarytree.Node{
		Val: 16,
	}
	binarytree.ConnectNode(&node1, &node2, &node3)
	binarytree.ConnectNode(&node2, &node4, &node5)
	binarytree.ConnectNode(&node3, &node6, &node7)

	binarytree.PrintTree(&node1)
	listHead := Convert(&node1)
	printLinkedList(listHead)
}
