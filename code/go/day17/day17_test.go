package day17

import (
	"testing"
	"fmt"
	"../utils/linkedlist"
)

// 某些节点重复
func TestFunc1(t *testing.T) {
	fmt.Printf("某些节点重复\n")
	ll := linkedlist.NewLinkedList()
	ll.Append(11)
	ll.Append(12)
	ll.Append(14)
	ll.Append(14)
	ll.Append(17)
	DeleteDuplication(&ll)
	ll.PrintList()
}

// 没有节点重复
func TestFunc2(t *testing.T) {
	fmt.Printf("\n没有节点重复\n")
	ll := linkedlist.NewLinkedList()
	ll.Append(11)
	ll.Append(12)
	ll.Append(13)
	ll.Append(14)
	ll.Append(17)
	DeleteDuplication(&ll)
	ll.PrintList()
}

// 所有节点重复
func TestFunc3(t *testing.T) {
	fmt.Printf("\n所有节点重复\n")
	ll := linkedlist.NewLinkedList()
	ll.Append(11)
	ll.Append(11)
	ll.Append(11)
	ll.Append(11)
	ll.Append(11)
	DeleteDuplication(&ll)
	ll.PrintList()
}

// 所有节成对出现
func TestFunc4(t *testing.T) {
	fmt.Printf("\n所有节成对出现\n")
	ll := linkedlist.NewLinkedList()
	ll.Append(11)
	ll.Append(11)
	ll.Append(14)
	ll.Append(14)
	DeleteDuplication(&ll)
	ll.PrintList()
}