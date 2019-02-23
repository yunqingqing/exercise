package binarytree

import "fmt"

type Node struct {
	Val   int
	Right *Node
	Left  *Node
}

func ConnectNode(root, left, right *Node) {
	root.Left = left
	root.Right = right
}

func PrintTree(rootNode *Node) {
	printTreeNode(rootNode)
	if rootNode != nil {
		if rootNode.Left != nil {
			printTreeNode(rootNode.Left)
		}

		if rootNode.Right != nil {
			printTreeNode(rootNode.Right)
		}
	}
}

func printTreeNode(node *Node) {
	if node != nil {
		fmt.Printf("value of this node is: %d\n", node.Val)
		if node.Left != nil {
			fmt.Printf("value of left child is: %d\n", node.Left.Val)
		} else {
			fmt.Printf("left node is nil.\n")
		}

		if node.Right != nil {
			fmt.Printf("value of right child is: %d\n", node.Right.Val)
		} else {
			fmt.Printf("right node is nil.\n")
		}
	} else {
		fmt.Printf("this node is nil.\n")
	}
	fmt.Printf("")
}