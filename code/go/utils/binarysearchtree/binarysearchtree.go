// Binary Search Tree implement.
package binarysearchtree

import "fmt"

const MaxUint = ^uint(0) 
const MaxInt = int(MaxUint >> 1) 
const MinInt = -MaxInt - 1

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
	if node != nil {
		InOrderPrint(node.leftChild)
		fmt.Printf("%+v->", node.data)
		InOrderPrint(node.rightChild)
	}
}

// 先序遍历
func PreOrderPrint(node *Node) {
	if node != nil {
		fmt.Printf("%+v->", node.data)
		InOrderPrint(node.leftChild)
		InOrderPrint(node.rightChild)
	}

}

// 后序遍历
func PostOrderPrint(node *Node) {
	if node != nil {
		InOrderPrint(node.leftChild)
		InOrderPrint(node.rightChild)
		fmt.Printf("%+v->", node.data)
	}

}

// 获得最小节点
func (tree *BinarySearchTree) GetMin() *Node {
	return findmin(tree.root)
}

// 内部查找最小值
func findmin(node *Node) *Node {
	currentNode := node
	if currentNode != nil {
		for currentNode.rightChild != nil {
			currentNode = currentNode.rightChild
		}
	}
	return currentNode
}

// 查找给定值的节点
func (tree *BinarySearchTree) Find(value int) *Node {
	resultNode, _ := find(tree.root, nil, value)
	return resultNode
}

// 内部递归查找方法, 返回查找值对应的节点及期父节点
func find(node, parentNode *Node, value int) (*Node, *Node) {
	var resultNode *Node
	if node != nil {
		if node.data > value {
			resultNode, parentNode = find(node.leftChild, node, value)
		} else if node.data < value {
			resultNode, parentNode = find(node.rightChild, node, value)
		} else {
			resultNode = node
		}
	}
	return resultNode, parentNode
}

// 删除给定值对应的节点
func (tree *BinarySearchTree) Delete(value int) {
	delete(tree.root, nil, value)
}

// 内部删除节点方法
func delete(node, parentNode *Node, value int) {
	deleteNode, parentNode := find(node, parentNode, value)
	if deleteNode.rightChild == nil && deleteNode.leftChild == nil {
		// 没有子节点，直接删除
		parentNode.leftChild = nil
		parentNode.rightChild = nil
	} else if deleteNode.rightChild == nil{
		// 一个子节点, 该节点父节点的指针指向该节点的子节点
		deleteNode = deleteNode.leftChild
	} else if deleteNode.leftChild == nil { 
		// 一个子节点, 该节点父节点的指针指向该节点的子节点
		deleteNode = deleteNode.rightChild
	} else { 
		// 两个子节点,
		rightMinNode := findmin(deleteNode.rightChild)
		deleteNode.data = rightMinNode.data
		delete(deleteNode.rightChild, deleteNode, rightMinNode.data)
	}
	
}

// 判断当前是否为二叉搜索树
func (tree *BinarySearchTree) IsBinarySearchTree() bool {
	return between(tree.root, MinInt, MaxInt)
}

func between(node *Node, min, max int) bool {
	if node == nil {
		return true
	}
	return node.data > min && node.data < max && between(node.leftChild, min, node.data) && between(node.rightChild, node.data, max)
}