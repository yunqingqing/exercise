package binarysearchtree

import (
	"os"
	"testing"
)

var tree BinarySearchTree

func Setup() {
	tree = NewBinarySearchTree()
}

func TestInsert(t *testing.T) {
	tree.Insert(12)
	tree.Insert(4)
	tree.Insert(2)
	tree.Insert(13)
	tree.Insert(1)
	tree.Insert(6)
}

func TestMain(m *testing.M) {
	Setup()
	retCode := m.Run()
	os.Exit(retCode)
}
