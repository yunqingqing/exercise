package day37

func partition(data *[]int, length, start, end int) int {
	pivot := start
	index := pivot + 1

	for i := index; i <= end; i++ {
		if (*data)[i] < (*data)[pivot] {
			(*data)[i], (*data)[index] = (*data)[index], (*data)[i]
			index++
		}
	}
	(*data)[pivot], (*data)[index - 1] = (*data)[index - 1], (*data)[pivot]
	return index - 1
}

func QuickSort(data *[]int, length, start, end int) {
	if start == end {
		return
	}
	index := partition(data, length, start, end)

	if index > start {
		QuickSort(data, length, start, index - 1)
	}
	if index < end {
		QuickSort(data, length, index + 1, end)
	}
}