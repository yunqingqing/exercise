// Binary Search Tree implement.
package binarysearchtree

import "fmt"

type Node struct {
	leftChild  *Node
	rightChild *Node
	data       int
}

type BinarySearchTree struct {
	root *Node
}

// 初始化树
func NewBinarySearchTree() BinarySearchTree {
	return BinarySearchTree{}
}

// 设置根节点
func (tree *BinarySearchTree) SetRootNode(node *Node) {
	tree.root = node
}

// 往节点插入值
func InsertNode(node, newNode *Node) {
	if newNode.data > node.data {
		if node.rightChild == nil {
			node.rightChild = newNode
		} else {
			InsertNode(node.rightChild, newNode)
		}
	} else if newNode.data < node.data {
		if node.leftChild == nil {
			node.leftChild = newNode
		} else {
			InsertNode(node.leftChild, newNode)
		}
	} else {
		fmt.Printf("node value repeat.")
	}
}

// 给树插入值
func (tree *BinarySearchTree) Insert(value int) {
	node := Node{
		data: value,
	}
	if tree.root == nil {
		tree.SetRootNode(&node)
	} else {
		InsertNode(tree.root, &node)
	}
}

// 中序遍历
func InOrderPrint(node *Node) {
	for node != nil {
		InOrderPrint(node.leftChild)
		fmt.Printf("%+v->", node.data)
		InOrderPrint(node.rightChild)
	}
}

// 先序遍历
func PreOrderPrint(node *Node) {
	for node != nil {
		fmt.Printf("%+v->", node.data)
		InOrderPrint(node.leftChild)
		InOrderPrint(node.rightChild)
	}

}

// 后序遍历
func PostOrderPrint(node *Node) {
	for node != nil {
		InOrderPrint(node.leftChild)
		InOrderPrint(node.rightChild)
		fmt.Printf("%+v->", node.data)
	}

}

// 查找给定值的节点
func (tree *BinarySearchTree) Find(value int) *Node {
	return find(tree.root, value)
}

// 内部递归查找方法
func find(node *Node, value int) *Node {
	var resultNode *Node
	if node != nil {
		if node.data > value {
			find(node.leftChild, value)
		} else if node.data < value {
			find(node.rightChild, value)
		} else {
			resultNode = node
		}
	}
	return resultNode
}

// 删除给定值对应的节点
func (tree *BinarySearchTree) Delete(value int) {
	node := tree.Find(value)
	if node.rightChild != nil && node.leftChild != nil {
		// 两个子节点,

	} else if node.rightChild == nil || node.leftChild == nil { 
		// 一个子节点, 该节点父节点的指针指向该节点的子节点

	} else { 
		// 没有子节点，直接删除

	}
}
