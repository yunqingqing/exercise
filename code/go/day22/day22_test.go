package day22

import (
	"../utils/linkedlist"
	"os"
	"testing"
)

var ll linkedlist.List = linkedlist.NewLinkedList()
var ll1 linkedlist.List = linkedlist.NewLinkedList()

func setUp() {
	// TODO: 构建更多的测试用例
	// 构造一个存在环的链表
	ll.Append(1)
	ll.Append(2)
	node3 := ll.Append(3)
	ll.Append(4)
	node5 := ll.Append(5)
	node5.SetNext(node3)

	// 构造一个不存在的环的链表
	ll1.Append(1)
	ll1.Append(2)
	ll1.Append(3)
	ll1.Append(4)
	ll1.Append(5)
}

func TestFunc(t *testing.T) {
	if EntryNodeOfLoop(&ll).GetData() != 3 {
		t.Errorf("not correct.")
	}

	if EntryNodeOfLoop(&ll1) != nil {
		t.Errorf("not correct.")
	}
}

func TestMain(m *testing.M) {
	setUp()
	retCode := m.Run()
	os.Exit(retCode)
}
