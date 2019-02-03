package day21

import (
	"../utils/linkedlist"
	"os"
	"testing"
)

var ll linkedlist.List = linkedlist.NewLinkedList()

func setUp() {
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	ll.Append(4)
	ll.Append(5)
}

func TestFunc(t *testing.T) {
	if FindKthToTail(&ll, 2) != 4 {
		t.Errorf("not correct.")
	}

	if FindKthToTail(&ll, 1) != 5 {
		t.Errorf("not correct.")
	}

	if FindKthToTail(&ll, 5) != 1 {
		t.Errorf("not correct.")
	}
}

func TestMain(m *testing.M) {
	setUp()
	retCode := m.Run()
	os.Exit(retCode)
}
