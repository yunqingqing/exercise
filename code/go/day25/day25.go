package day25

import "../utils/binarytree"

// 判断子树结构是否一致
func doesTree1HaveTree2(rootNode1, rootNode2 *binarytree.Node) bool {
	if rootNode2 == nil {
		return true
	}

	if rootNode1 == nil {
		return false
	}

	if rootNode1.Val != rootNode2.Val {
		return false
	}

	return doesTree1HaveTree2(rootNode1.Left, rootNode2.Left) && doesTree1HaveTree2(rootNode1.Right, rootNode2.Right)
}

func HasSubtree(rootNode1, rootNode2 *binarytree.Node) bool {
	result := false

	if rootNode1 != nil && rootNode2 != nil {

		// 根节点相同，判断下子树是否一致
		if rootNode1.Val == rootNode2.Val {
			result = doesTree1HaveTree2(rootNode1, rootNode2)
		}

		// 根节点不同，看子树中是否还有相同的根节点
		if !result {
			result = HasSubtree(rootNode1.Left, rootNode2)
		}

		if !result {
			result = HasSubtree(rootNode1.Right, rootNode2)
		}
	}

	return result
}