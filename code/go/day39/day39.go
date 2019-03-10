package day39

import "container/heap"

func partition(data *[]int, start, end int) int {
	pivot := start
	index := pivot + 1
	for i := index; i <= end; i++ {
		if (*data)[i] < (*data)[pivot] {
			(*data)[i], (*data)[index] = (*data)[index], (*data)[i]
			index++
		}
	}
	(*data)[pivot], (*data)[index-1] = (*data)[index-1], (*data)[pivot]
	return index - 1
}

// partition解法
func GetLeastNumbers1(array []int, length, k int) []int {
	if k > length || array == nil || k < 0 {
		return nil
	}

	start := 0
	end := length - 1
	index := partition(&array, start, end)
	output := []int{}

	for index != k-1 {
		if index > k-1 {
			end = index - 1
			index = partition(&array, start, end)
		} else {
			start = index + 1
			index = partition(&array, start, end)
		}
	}

	for i := 0; i < k; i++ {
		output = append(output, array[i])
	}
	return output
}


// heap解法
type IntHeap []int

func (h IntHeap) Len() int { return len(h) }

func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func GetLeastNumbers2(array []int, length, k int) []int {
	if k > length || array == nil || k < 0 {
		return nil
	}
	
	h := &IntHeap{}
	// 初始化堆
	heap.Init(h)
	for _, item := range(array) {
		if h.Len() == k {
			max := heap.Pop(h)
			if max.(int) > item {
				heap.Push(h, item)
			} else {
				heap.Push(h, max)
			}
		} else {
			heap.Push(h, item)
		}
	}
	output := []int{}
	for h.Len() != 0 {
		output = append(output, heap.Pop(h).(int))
	}
	return output
}