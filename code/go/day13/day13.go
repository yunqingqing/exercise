package day13

import "math"

// 动态规划
func MaxProductAfterCyttingSolution1(length int) int {
	if length < 2 {
		return 0
	}
	if length == 2 {
		return 1
	}
	if length == 3 {
		return 2
	}
	
	products := make([]int, length + 1)
	products[0] = 0
	products[1] = 1
	products[2] = 2
	products[3] = 3

	max := 0
	for i := 4; i <= length; i++ {
		max = 0
		for j := 1; j <= i / 2; j++ {
			product := products[j] * products[i - j]
			if max < product {
				max = product
			}

			products[i] = max
		}
	}

	max = products[length]
	return max
}

// 贪婪算法
func MaxProductAfterCyttingSolution2(length int) int {
	if length < 2 {
		return 0
	}	

	if length == 2 {
		return 1
	}

	if length == 3 {
		return 2
	}

	// 尽可能减去长度为3的绳子段
	timesOf3 := length / 3

	// 当绳子最后剩下长度为4的时候,不能在剪去长度为3的绳子段
	// 此时更好的方法是把绳子剪成长度为2的两段
	if length - timesOf3 * 3 == 1 {
		timesOf3 -= 1
	}

	timesOf2 := (length - timesOf3 * 3) / 2

	return int(math.Pow(3, float64(timesOf3)) * math.Pow(2, float64(timesOf2)))
}