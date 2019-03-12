package day40

import "container/heap"

// 大顶堆
type IntMaxHeap []int

func (h IntMaxHeap) Len() int { return len(h) }

func (h IntMaxHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h IntMaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 小顶堆
type IntMinHeap []int

func (h IntMinHeap) Len() int { return len(h) }

func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h IntMinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntMinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 要确保大顶堆里的数据都比小顶堆的要小
// 两个堆数据量之差不超过一
var max = &IntMaxHeap{}
var min = &IntMinHeap{}

func init() {
	heap.Init(max)
	heap.Init(min)
}

func Insert(num int) {
	if ((min.Len() + max.Len()) & 1) == 0 {
		// 偶数插入小顶堆
		if max.Len() == 0 {
			heap.Push(min, num)
			return
		}
		tmp := heap.Pop(max)
		if max.Len() > 0 && num < tmp.(int) {
			// 比大顶堆的根小，则把数据插入大顶堆
			// 然后从大顶堆的最大值插入到小顶堆
			heap.Push(max, tmp)
			heap.Push(max, num)
			heap.Push(min, heap.Pop(max))
		} else {
			heap.Push(max, tmp)
			heap.Push(min, num)
		}
	} else {
		// 奇数插入大顶堆
		if min.Len() == 0 {
			heap.Push(max, num)
			return
		}
		tmp := heap.Pop(min)
		if min.Len() > 0 && tmp.(int) < num {
			// 比小顶堆的根大，则把数据插入小顶堆
			// 然后从小顶堆的最大值插入到大顶堆
			heap.Push(min, tmp)
			heap.Push(min, num)
			// x := heap.Pop(min)
			// fmt.Println("xxxxx", x)
			heap.Push(max, heap.Pop(min))
		} else {
			heap.Push(min, tmp)
			heap.Push(max, num)
		}
	}
}

func GetMedian() int {
	size := max.Len() + min.Len()
	if size == 0 {
	}
	median := 0
	if (size & 1) == 1 {
		median = heap.Pop(min).(int)
		heap.Push(min, median)
	} else {
		tmp1 := heap.Pop(min)
		tmp2 := heap.Pop(max)
		median = (tmp1.(int) + tmp2.(int)) / 2
		heap.Push(min, tmp1)
		heap.Push(max, tmp2)
	}
	return median
}
