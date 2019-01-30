package day17

import "../utils/linkedlist"


func DeleteDuplication(l *linkedlist.List) {
	// TODO 后续要考虑下内存释放，以及逻辑是否能够在整理下
	pPreNode := &linkedlist.Node{}
	cur := l.GetHead()
	curHead := cur
	needDelete := false
	isHeadChanged := false
	for cur != nil {
		next := cur.GetNext()
		if next == nil {
			if needDelete {
				pPreNode.SetNext(next)
				if isHeadChanged {
					l.SetHead(next)
					isHeadChanged = false
					curHead = next
				}
			}
		} else if cur.GetData() != next.GetData() {
			if needDelete {
				pPreNode.SetNext(next)
				if isHeadChanged {
					l.SetHead(next)
					isHeadChanged = false
					curHead = next
				}
			} else {
				needDelete = false
			}
			pPreNode = cur
		} else {
			needDelete = true
			if cur == curHead {
				isHeadChanged = true
			}
		}
		cur = next
	}
}