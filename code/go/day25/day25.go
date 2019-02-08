package day25

type Node struct {
	val   int
	right *Node
	left  *Node
}

func doesTree1HaveTree2(rootNode1, rootNode2 *Node) bool {
	if rootNode2 == nil {
		return true
	}

	if rootNode1 == nil {
		return false
	}

	if rootNode1.val != rootNode2.val {
		return false
	}

	return doesTree1HaveTree2(rootNode1.left, rootNode2.left) && doesTree1HaveTree2(rootNode1.right, rootNode2.right)
}

func HasSubtree(rootNode1, rootNode2 *Node) bool {
	result := false

	if rootNode1 != nil && rootNode2 != nil {
		if rootNode1.val == rootNode2.val {
			result = doesTree1HaveTree2(rootNode1, rootNode2)
		}

		if !result {
			result = HasSubtree(rootNode1.left, rootNode2)
		}

		if !result {
			result = HasSubtree(rootNode1.right, rootNode2)
		}
	}

	return result
}