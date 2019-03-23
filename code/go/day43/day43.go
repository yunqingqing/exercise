package day43

import "math"

func DigitAtIndex(index int) int {
	if index < 0 {
		return -1
	}

	digits := 1
	for {
		numbers := countOfIntegers(digits)
		if index < numbers * digits {
			// numbers*m m位数总共占了多少位
			// 输入的数小于总位数的话，说明目标就在这里面
			return digitAtIndex(index, digits)
		}

		// 不再m位数中，把前面的跳过。直接从m+1位的数里开始
		index -= digits * numbers

		// 位数增加
		digits++
	}

}

// 得到m位数总共有多少个
func countOfIntegers(digits int) int {
	if digits == 1 {
		return 10
	}

	count := math.Pow10(digits - 1)
	return int(9 * count)
}

// 当确定目标为m位数时，通过这个方法找出具体是哪个
func digitAtIndex(index, digits int) int {
	//index 811 digits:3
	number := beginNumber(digits) + index / digits
	indexFromRight := digits - index % digits  //2
	for i:=1; i < indexFromRight; i++ {
		number /= 10
	}
	return number % 10
}

// 确定m位数的第一个数字，第一个两位数是10
// 第一个三位数是100
func beginNumber(digits int) int {
	if digits == 1 {
		return 0
	}
	return int(math.Pow10(digits - 1))
}