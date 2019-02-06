package day23

import "../utils/linkedlist"

func ReverseList(l *linkedlist.List) {
	var prev, next *linkedlist.Node
	cur := l.GetHead()

	// 把当前的头设置为尾
	l.SetTail(cur)
	for cur != nil {
		next = cur.GetNext()
		// 掉转指针
		cur.SetNext(prev)
		prev = cur
		cur = next
	}

	// 把当前的尾设置为头
	l.SetHead(prev)
}
