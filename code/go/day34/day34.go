package day34

func CloneNodes(head *ComplexListNode) *ComplexListNode {
	cloneNodes(head)
	connectSiblingNodes(head)
	return reconnectNodes(head)
}

// 第一步复制该节点，插入在该节点和该节点的下一个节点之间
func cloneNodes(head *ComplexListNode) {
	node := head
	for node != nil {
		clonedNode := &ComplexListNode{}
		clonedNode.value = node.value
		clonedNode.next = node.next
		clonedNode.sibling = nil

		node.next = clonedNode
		node = clonedNode.next
	}
}

// 第二步，复制链表的兄弟关系
func connectSiblingNodes(head *ComplexListNode) {
	node := head
	for node != nil {
		cloned := node.next
		if node.sibling != nil {
			cloned.sibling = node.sibling.next
		}
		node = cloned.next

	}
}

// 把链表拆成两个
func reconnectNodes(head *ComplexListNode) *ComplexListNode {
	node := head
	var clonedHead, clonedNode *ComplexListNode

	if node != nil {
		clonedHead = node.next
		clonedNode = node.next
		node.next = clonedNode.next
		node = node.next
	}

	for node != nil {
		clonedNode.next = node.next
		clonedNode = clonedNode.next

		node.next = clonedNode.next
		node = node.next
	}

	return clonedHead
}
