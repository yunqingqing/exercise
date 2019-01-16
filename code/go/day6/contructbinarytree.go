package main

import "fmt"

type Node struct {
	value int
	left *Node
	right *Node
}

// 获得左子树元素
func getLeftVals(inorder []int, root int) []int {
	idx := 0
	for inorder[idx] != root {
		idx++
	}
	if idx == 0 {
		return nil
	}

	return inorder[:idx]
}

// 获得右子树元素
func getRightVals(inorder []int, root int) []int {
	idx := 0
	for inorder[idx] != root {
		idx++
	}
	if idx == len(inorder)-1 {
		return nil
	}
	return inorder[idx+1:]
}

func ConstructTree(preorder, inorder []int) *Node {
	if preorder == nil || inorder == nil || len(preorder) <= 0 || len(inorder) <=0 {
		return nil
	}

	if len(preorder) == 1 && len(inorder) == 1 { // 递归推出条件，构建只有一个节点的二叉树
		return &Node{value: preorder[0]}
	}

	rootValue := preorder[0]
	root := &Node{value: rootValue}

	leftValues := getLeftVals(inorder, rootValue)
	rightValues := getRightVals(inorder, rootValue)
	leftSubTree := ConstructTree(preorder[1:], leftValues)
	rightSubTree := ConstructTree(preorder[len(leftValues)+1:], rightValues)

	root.left = leftSubTree
	root.right = rightSubTree
	return root
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

func preorder(root *Node, res []int) []int {
	if root == nil {
		return res
	}
	res = append(res, root.value)
	res = preorder(root.left, res)
	res = preorder(root.right, res)
	return res
}

func main() {
	preorderArray := []int {1, 2, 4, 7, 3, 5, 6, 8}
	inorderArray := []int {4, 7, 2, 1, 5, 3, 8, 6}
	
	root := ConstructTree(preorderArray, inorderArray)
	fmt.Printf("preorder=%+v\n", preorder(root, []int{}))
	fmt.Printf("inorders=%+v\n", inorder(root, []int{}))
}