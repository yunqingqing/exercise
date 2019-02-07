package day24

import "../utils/linkedlist"

func Merge(l1, l2 *linkedlist.List) linkedlist.List {
	// 对输入的异常情况进行处理
	if l1 == nil && l2 == nil {
		return linkedlist.NewLinkedList()
	} else if l1 == nil {
		return *l2
	} else if l2 == nil {
		return *l1
	}

	p1 := l1.GetHead()
	p2 := l2.GetHead()
	newL := linkedlist.NewLinkedList()

	for {
		if p1 != nil &&  p2 != nil {
			val1 := p1.GetData().(int)
			val2 := p2.GetData().(int)

			if val1 > val2 {
				newL.Append(val2)
				p2 = p2.GetNext()
			} else {
				newL.Append(val1)
				p1 = p1.GetNext()
			}
		} else {
			break
		}
	}

	if p1 == nil {
		for p2 != nil {
			newL.Append(p2.GetData())
			p2 = p2.GetNext()
		}
	}

	if p2 == nil {
		for p1 != nil {
			newL.Append(p1.GetData())
			p1 = p1.GetNext()
		}
	}
	
	return newL
}