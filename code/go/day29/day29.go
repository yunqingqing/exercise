package day29

import (
	"../utils/binarytree"
	"fmt"
	"gopkg.in/karalabe/cookiejar.v1/collections/deque"
)

func PrintFromTopToBottom(node *binarytree.Node) {
	dq := deque.New()

	dq.PushRight(node)

	for dq.Size() != 0 {
		currentNode := dq.PopLeft().(*binarytree.Node)
		fmt.Println(currentNode.Val)
		if currentNode.Left != nil {
			dq.PushRight(currentNode.Left)
		}
		if currentNode.Right != nil {
			dq.PushRight(currentNode.Right)
		}
	}
}
