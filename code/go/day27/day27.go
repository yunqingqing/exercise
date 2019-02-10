package day27

import "../utils/binarytree"

func IsSymmetrical(root *binarytree.Node) bool {
	return isSymmetrical(root, root)
}

func isSymmetrical(root1, root2 *binarytree.Node) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	if root1.Val != root2.Val {
		return false
	}

	return isSymmetrical(root1.Left, root2.Right) && isSymmetrical(root1.Right, root2.Left)
}
