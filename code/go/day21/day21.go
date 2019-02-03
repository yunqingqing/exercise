package day21

import (
	"../utils/linkedlist"
)

func FindKthToTail(ll *linkedlist.List, k int) interface{} {
	pFirst := ll.GetHead()
	pSecond := pFirst
	for i := 0; i < k-1; i++ {
		if pFirst.GetNext() != nil {
			pFirst = pFirst.GetNext()
		} else {
			return nil
		}
	}

	for pFirst.GetNext() != nil {
		pFirst = pFirst.GetNext()
		pSecond = pSecond.GetNext()
	}

	return pSecond.GetData()
}
