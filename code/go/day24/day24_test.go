package day24

import (
	"os"
	"testing"
	"../utils/linkedlist"
)

var ll linkedlist.List = linkedlist.NewLinkedList()
var ll1 linkedlist.List = linkedlist.NewLinkedList()

func setUp() {
	ll.Append(1)
	ll.Append(3)
	ll.Append(5)
	ll.Append(11)
	ll.Append(12)

	ll1.Append(2)
	ll1.Append(4)
	ll1.Append(6)
	ll1.Append(10)
}

func TestFunc(t *testing.T) {
	result := Merge(&ll, &ll1)
	result.PrintList()

	result1 := Merge(nil, nil)
	result1.PrintList()
}

func TestMain(m *testing.M) {
	setUp()
	retCode := m.Run()
	os.Exit(retCode)
}
