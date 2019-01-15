// 从尾到头打印链表
package main

import (
	"../utils/linkedlist"
	"fmt"
	"github.com/golang-collections/collections/stack"
)

// 修改链表结构实现，把指针倒转
func PrintListReversinglyModifyPointer(l *linkedlist.List) {
	var prevNode *linkedlist.Node
	currentNode := l.GetHead()
	l.SetTail(currentNode)
	for currentNode != nil {
		nextNode := currentNode.GetNext()		
		currentNode.SetNext(prevNode)
		prevNode = currentNode
		currentNode = nextNode
	}
	l.SetHead(prevNode)
	l.PrintList()
}

// 利用栈结构迭代实现
func PrintListReversinglyIteratively(l linkedlist.List) {
	nodes := stack.New()  // 申请一个栈
	currentNode := l.GetHead()
	for currentNode != nil {
		nodes.Push(currentNode)  // 迭代链表，节点依次入栈
		currentNode = currentNode.GetNext()
	}

	for nodes.Len() != 0 {
		node := nodes.Pop().(*linkedlist.Node)
		fmt.Printf("%+v->",node.GetData())  // 出栈。打印节点值
	}
}

// 利用递归实现
func PrintNode(node *linkedlist.Node) {
	if node.GetNext() != nil {
		PrintNode(node.GetNext())
	}
	fmt.Printf("%+v->",node.GetData())
}

func PrintListReversinglyRecurisively(l linkedlist.List) {
	currentNode := l.GetHead()
	if currentNode != nil {
		PrintNode(currentNode)
	}
}

func main() {
	ll := linkedlist.NewLinkedList()
	ll.Append(12)
	ll.Append(14)
	ll.Append("a")
	ll.Append(17)
	fmt.Printf("Initially linked list.\n")
	ll.PrintList()
	fmt.Printf("\nAfter iteration, reversed linked list.\n")
	PrintListReversinglyIteratively(ll)
	fmt.Printf("\nAfter Recursive, reversed linked list.\n")
	PrintListReversinglyRecurisively(ll)
	fmt.Printf("\nAfter Modify pointer,reserved linked list.\n")
	PrintListReversinglyModifyPointer(&ll)
}