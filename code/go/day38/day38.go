package day38

func checkMoreThanHalf(array []int, length, number int) bool {
	times := 0
	for _, num := range array {
		if num == number {
			times++
		}
	}

	isMoreThanHalf := true
	if times*2 < length {
		isMoreThanHalf = false
	}
	return isMoreThanHalf
}

func MoreThanHalfNum(array []int, length int) int {
	result := array[0]
	times := 1
	for _, num := range array {
		if times == 0 {
			result = num
			times = 1
		} else if num == result {
			times++
		} else {
			times--
		}
	}
	if !checkMoreThanHalf(array, length, result) {
		result = 0
	}

	return result
}
