package day22

import "../utils/linkedlist"

// 计算环的长度
func LengthOfLoop(ll *linkedlist.List) int {
	head := ll.GetHead()
	p1 := head.GetNext() // 移动慢的指针，一次一步
	p2 := p1.GetNext()   // 移动快的指针，一次两步
	length := 0          // 环的长度
	hasLoop := false

	for p1 != nil && p2 != nil {
		//p2追上p1, 说明存在环
		if p1 == p2 {
			hasLoop = true
			break
		}

		p1 = p1.GetNext()
		p2 = p2.GetNext()
		if p2 != nil {
			p2 = p2.GetNext()
		}
	}

	// 存在环的话，此时p2不动，p1在以1的步速移动。
	// 统计移动次数，直到与p2相遇,移动的次数即为环大小。
	if hasLoop {
		length = 1
		p1 = p1.GetNext()
		for p1 != p2 {
			length++
			p1 = p1.GetNext()
		}
	}
	return length
}

// 寻找环的入口节点
func EntryNodeOfLoop(ll *linkedlist.List) *linkedlist.Node {
	length := LengthOfLoop(ll)

	// 不存在环
	if length == 0 {
		return nil
	}

	// 申明两个指针，之间差了环的长度个节点
	p1 := ll.GetHead()
	p2 := p1
	for i := 0; i < length; i++ {
		p2 = p2.GetNext()
	}

	for p1 != p2 {
		p1 = p1.GetNext()
		p2 = p2.GetNext()
	}

	return p1

}
