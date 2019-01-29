package day16

import (
	"testing"
	"fmt"
	"../utils/linkedlist"
)


func TestFunc(t *testing.T) {
	ll := linkedlist.NewLinkedList()
	node1 := ll.Append(12)
	node2 := ll.Append(14)
	node3 := ll.Append("a")
	node4 := ll.Append(17)
	fmt.Printf("Initially linked list.\n")
	ll.PrintList()
	fmt.Printf("\ndelete 'a'\n")
	Delete(&ll, node3)
	ll.PrintList()
	fmt.Printf("\ndelete 17\n")
	Delete(&ll, node4)
	ll.PrintList()
	fmt.Printf("\ndelete 12\n")
	Delete(&ll, node1)
	ll.PrintList()
	fmt.Printf("\ndelete 14\n")
	Delete(&ll, node2)
	ll.PrintList()
	fmt.Printf("\n")
}