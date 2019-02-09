package day26

type Node struct {
	val   int
	right *Node
	left  *Node
}

func MirrorRecursively(node *Node) {
	if node.left == nil && node.right == nil {
		return
	}

	node.left, node.right = node.right, node.left

	if node.left != nil {
		MirrorRecursively(node.left)
	}

	if node.right != nil {
		MirrorRecursively(node.right)
	}
}
