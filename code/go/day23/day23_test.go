package day23

import (
	"../utils/linkedlist"
	"fmt"
	"os"
	"testing"
)

var ll linkedlist.List = linkedlist.NewLinkedList()
var ll1 linkedlist.List = linkedlist.NewLinkedList()

func setUp() {
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	ll.Append(4)
	ll.Append(5)

	ll1.Append(1)
}

func TestFunc(t *testing.T) {
	ReverseList(&ll)
	ll.PrintList()

	fmt.Printf("\n===\n")

	ReverseList(&ll1)
	ll1.PrintList()
	fmt.Printf("\n")
}

func TestMain(m *testing.M) {
	setUp()
	retCode := m.Run()
	os.Exit(retCode)
}
