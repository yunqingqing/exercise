package binarysearchtree

import (
	"os"
	"testing"
	"fmt"
)

var tree BinarySearchTree

func Setup() {
	tree = NewBinarySearchTree()
	tree.Insert(12)
	tree.Insert(4)
	tree.Insert(2)
	tree.Insert(13)
	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(17)
	tree.Insert(6)
	tree.Insert(7)
}

func TestInsert(t *testing.T) {
	tree.Insert(8)
	if !tree.IsBinarySearchTree() {
		t.Errorf("not binary tree")
	}
	InOrderPrint(tree.root)
	fmt.Printf("\n")
	tree.Delete(8)
	tree.Delete(2)
	tree.Delete(6)
	if !tree.IsBinarySearchTree() {
		t.Errorf("not binary tree")
	}
	InOrderPrint(tree.root)
}



func TestMain(m *testing.M) {
	Setup()
	retCode := m.Run()
	os.Exit(retCode)
}
