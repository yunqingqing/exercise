// 二叉树给定一个节点获取它在中序遍历序列中的下一个节点
package main

import "fmt"

type Node struct {
	value int
	left *Node
	right *Node
	parent *Node
}

func GetNext(node *Node) *Node {
	if node == nil {
		return nil
	}
	var next *Node
	if node.right != nil {
		next = node.right
		for next.left != nil {
			next = next.left
		}
	} else if node.parent != nil {
		cur := node
		parent := node.parent
		for parent != nil && cur == parent.right {
			cur = parent
			parent = parent.parent
		}
		next = parent
	}
	return next
}

func inorder(root *Node, res []int) []int {
	if root == nil {
		return res
	}
	res = inorder(root.left, res)
	res = append(res, root.value)
	res = inorder(root.right, res)
	return res
}

// 帮助构造测试二叉树
func ConnectNodes(parent, left, right *Node) {
	parent.left = left
	parent.right = right
	if left != nil {
		left.parent = parent
	}
	if right != nil {
		right.parent = parent
	}
}


func main() {
	var nodes []*Node
	for i:=1; i < 11; i++ {
		nodes = append(nodes, &Node{value: i})
	}

	root := nodes[0]
	ConnectNodes(root, nodes[1], nodes[2])
	ConnectNodes(nodes[1], nodes[3], nodes[4])
	ConnectNodes(nodes[2], nodes[5], nodes[6])
	ConnectNodes(nodes[4], nodes[7], nodes[8])
	ConnectNodes(nodes[6], nodes[9], nil)

	// tree like:
	//          1
	//		 2       3
  	//    4   5   6    7
    //       8 9     10    
	

	fmt.Printf("inorders=%+v\n", inorder(root, []int{}))
	fmt.Printf("1 next was %+v\n", GetNext(nodes[0]))
	fmt.Printf("2 next was %+v\n", GetNext(nodes[1]))
	fmt.Printf("9 next was %+v\n", GetNext(nodes[8]))
	fmt.Printf("7 next was %+v\n", GetNext(nodes[6]))
	fmt.Printf("inorders=%+v\n", inorder(root, []int{}))
}