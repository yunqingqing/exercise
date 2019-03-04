package day39

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

func GetLeastNumbers(array []int, length, k int) []int {
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
