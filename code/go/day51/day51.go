package day51

func GetNumberOfK(data []int, length, k int) int {
	number := 0

	if length > 0 {
		first := getFirstK(data, length, k, 0, length - 1)
		last := getLastK(data, length, k, 0, length - 1)

		if first > -1 && last > -1 {
			number = last - first + 1
		}
	}

	return number
}

func getFirstK(data []int, length, k, start, end int) int {
	if start > end {
		return -1
	}

	middleIndex := (start + end) / 2
	middleData := data[middleIndex]

	if middleData == k {
		if (middleIndex > 0 && data[middleIndex - 1] != k) || middleIndex == 0 {
			return middleIndex
		} else {
			end = middleIndex - 1
		}
	} else if middleData > k {
		end = middleIndex - 1
	} else {
		start = middleIndex + 1
	}

	return getFirstK(data, length, k, start, end)
}

func getLastK(data []int, length, k, start, end int) int {
	if start > end {
		return -1
	}

	middleIndex := (start + end) / 2
	middleData := data[middleIndex]

	if middleData == k {
		if (middleIndex < length - 1 && data[middleIndex + 1] != k) || middleIndex == length - 1 {
			return middleIndex
		} else {
			start = middleIndex + 1
		}
	} else if (middleData < k) {
		start = middleIndex + 1
	} else {
		end = middleIndex - 1
	}

	return getLastK(data, length, k, start, end)
}