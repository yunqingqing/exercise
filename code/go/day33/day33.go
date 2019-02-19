package day33

import (
	"../utils/binarytree"
	"fmt"
)

func FindPath(node *binarytree.Node, sum int) {
	if node == nil {
		return
	}

	currentSum := 0
	path := []int{}
	findPath(node, sum, path, &currentSum)
}

func findPath(node *binarytree.Node, expectedSum int, path []int, currentSum *int) {
	*currentSum += node.Val
	path = append(path, node.Val)

	// 如果是叶结点，并且路径上结点的和等于输入的值
	// 打印出这条路径
	isLeaf := node.Left == nil && node.Right == nil
	if *currentSum == expectedSum && isLeaf {
		for _, num := range path {
			fmt.Printf(" %d", num)
		}
		fmt.Printf("\n")
	}

	// 如果不是叶结点，则遍历它的子结点
	if node.Left != nil {
		findPath(node.Left, expectedSum, path, currentSum)
	}
	if node.Right != nil {
		findPath(node.Right, expectedSum, path, currentSum)
	}

	// 在返回到父结点之前，在路径上删除当前结点，
	// 并在currentSum中减去当前结点的值
	*currentSum -= node.Val
	i := len(path) - 1
	// 删除slice最后一个元素
	path = append(path[:i], path[i+1:]...)
}
