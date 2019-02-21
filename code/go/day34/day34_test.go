package day34

import "testing"

//          -----------------
//         \|/              |
//  1-------2-------3-------4-------5
//  |       |      /|\             /|\
//  --------+--------               |
//          -------------------------
func TestFunc1(t *testing.T) {
	node1 := CreateNode(1)
	node2 := CreateNode(2)
	node3 := CreateNode(3)
	node4 := CreateNode(4)
	node5 := CreateNode(5)

	BuildNodes(node1, node2, node3)
	BuildNodes(node2, node3, node5)
	BuildNodes(node3, node4, nil)
	BuildNodes(node4, node5, node2)

	clonedHead := CloneNodes(node1)
	PrintList(clonedHead)
}

//          -----------------
//         \|/              |
//  1-------2-------3-------4-------5
//         |       | /|\           /|\
//         |       | --             |
//         |------------------------|
func TestFunc2(t *testing.T) {
	node1 := CreateNode(1)
	node2 := CreateNode(2)
	node3 := CreateNode(3)
	node4 := CreateNode(4)
	node5 := CreateNode(5)

	BuildNodes(node1, node2, nil)
	BuildNodes(node2, node3, node5)
	BuildNodes(node3, node4, node3)
	BuildNodes(node4, node5, node2)

	clonedHead := CloneNodes(node1)
	PrintList(clonedHead)
}

//          -----------------
//         \|/              |
//  1-------2-------3-------4-------5
//          |              /|\
//          |               |
//          |---------------|
func TestFunc3(t *testing.T) {
	node1 := CreateNode(1)
	node2 := CreateNode(2)
	node3 := CreateNode(3)
	node4 := CreateNode(4)
	node5 := CreateNode(5)

	BuildNodes(node1, node2, nil)
	BuildNodes(node2, node3, node4)
	BuildNodes(node3, node4, nil)
	BuildNodes(node4, node5, node2)

	clonedHead := CloneNodes(node1)
	PrintList(clonedHead)
}

// 只有一个节点
func TestFunc4(t *testing.T) {
	node1 := CreateNode(1)

	BuildNodes(node1, nil, node1)

	clonedHead := CloneNodes(node1)
	PrintList(clonedHead)
}
