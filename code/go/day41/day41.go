package day41

func FindGreatestSumOfSubArray(nums []int, length int) int {
	if length <= 0 {
		return 0
	}

	curSum := 0        // 当前计算的最大值
	var greatesSum int // 上一次计算到的最大值

	for index, item := range nums {
		if index == 0 {
			greatesSum = item
		}

		if curSum <= 0 {
			curSum = item
		} else {
			curSum += item
		}

		if curSum > greatesSum {
			greatesSum = curSum
		}
	}

	return greatesSum
}
