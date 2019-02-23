package day35

import "../utils/binarytree"

func Convert(rootNode *binarytree.Node) *binarytree.Node {
	var lastNodeInList *binarytree.Node

	convertNode(rootNode, &lastNodeInList)
	headOfList := lastNodeInList
	for headOfList != nil && headOfList.Left != nil {
		headOfList = headOfList.Left
	}
	
	return headOfList
}

func convertNode(node *binarytree.Node, lastNodeInList **binarytree.Node) {
	current := node
	
	if current.Left != nil {
		convertNode(current.Left, lastNodeInList)
	}
	
	// 当前节点加一个指向上一个节点的指针
	current.Left = *lastNodeInList

	if *lastNodeInList != nil {
		(*lastNodeInList).Right = current
	}

	*lastNodeInList = current
	if current.Right != nil {
		convertNode(current.Right, lastNodeInList)
	}
}