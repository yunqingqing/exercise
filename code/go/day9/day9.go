package day9

func Min(numbers []int, length int) int {
	index1 := 0
	index2 := length - 1
	indexMid := 0
	for numbers[index1] >= numbers[index2] {
		if index2 - index1 == 1 {
			indexMid = index2
			break
		}

		indexMid = (index1 + index2) / 2
		if numbers[index1] == numbers[index2] && numbers[indexMid] == numbers[index1] {
			return MinInOrder(numbers, index1, index2)
		}

		if numbers[indexMid] >= numbers[index1] {
			// 中间数大于第一个指针，说明在后面的区间
			index1 = indexMid
		} else if numbers[indexMid] <= numbers[index2]{
			index2 = indexMid
		}
	}
	return numbers[indexMid]
}

// 顺序遍历区间，返回最小值
func MinInOrder(numbers []int, index1, index2 int) int {
	result := numbers[index1]
	for i := index1 + 1; i < index2; i++ {
		if result > numbers[i] {
			result = numbers[i]
		}
	}
	return result
}