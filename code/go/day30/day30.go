package day30

import (
	"../utils/binarytree"
	"fmt"
	"gopkg.in/karalabe/cookiejar.v1/collections/deque"
)

func Print(node *binarytree.Node) {
	dq := deque.New()

	dq.PushRight(node)

	nextLevel := 0
	toBePrint := 1

	for dq.Size() != 0 {
		currentNode := dq.PopLeft().(*binarytree.Node)
		fmt.Printf("%d  ", currentNode.Val)
		toBePrint--
		if currentNode.Left != nil {
			dq.PushRight(currentNode.Left)
			nextLevel++
		}
		if currentNode.Right != nil {
			dq.PushRight(currentNode.Right)
			nextLevel++
		}

		if toBePrint == 0 {
			fmt.Printf("\n")
			toBePrint = nextLevel
			nextLevel = 0
		}
	}
}