package day31

import (
	"../utils/binarytree"
	"fmt"
	"github.com/golang-collections/collections/stack"
)

func Print(node *binarytree.Node) {
	stacks := [2]*stack.Stack{
		stack.New(),
		stack.New(),
	}

	current := 0
	next := 1

	stacks[current].Push(node)

	for stacks[0].Len() != 0 || stacks[1].Len() != 0 {
		currentNode := stacks[current].Pop().(*binarytree.Node)
		fmt.Printf("%d  ", currentNode.Val)

		if current == 0 {
			if currentNode.Left != nil {
				stacks[next].Push(currentNode.Left)
			}

			if currentNode.Right != nil {
				stacks[next].Push(currentNode.Right)
			}
		} else {
			if currentNode.Right != nil {
				stacks[next].Push(currentNode.Right)
			}

			if currentNode.Left != nil {
				stacks[next].Push(currentNode.Left)
			}

		}

		if stacks[current].Len() == 0 {
			fmt.Printf("\n")
			current = 1 - current
			next = 1 - next
		}
	}
}
