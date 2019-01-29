package day16

import "../utils/linkedlist"


func Delete(l *linkedlist.List, pToBeDeleted *linkedlist.Node) {
	pHead := l.GetHead()
	var nilNode *linkedlist.Node

	if pToBeDeleted.GetNext() != nil {
		next := pToBeDeleted.GetNext()
		pToBeDeleted.SetData(next.GetData())
		pToBeDeleted.SetNext(next.GetNext())
	} else if pHead.GetData() == pToBeDeleted.GetData() {
		l.SetHead(nilNode)
	} else {
		cur := pHead
		for cur.GetNext().GetData() != pToBeDeleted.GetData() {
			cur = cur.GetNext()
		}
		cur.SetNext(nilNode)
		l.SetTail(cur)
	}
}