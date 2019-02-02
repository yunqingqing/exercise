package day20

func RecorderOddEven1(array []int) []int {
	length := len(array)
	if length == 0 {
		return array
	}
	pLeft := 0           // 指向数组头
	pRight := length - 1 //  指向数组尾

	for pLeft < pRight {
		// 两个指针相遇后退出

		// 指向头的指针一直向右移，直到遇到第一个偶数
		for pLeft < pRight && (array[pLeft]&0x1) != 0 {
			pLeft++
		}

		// 指向尾的指针一直向左移，直到遇到第一个奇数
		for pRight > pLeft && (array[pRight]&0x1) == 0 {
			pRight--
		}

		// 把两个指针遇到到的数值交换
		if pLeft < pRight {
			array[pLeft], array[pRight] = array[pRight], array[pLeft]
		}
	}
	return array
}

// 把判断指针移动的的条件单独抽取出来，增加程序的可扩展性
func RecorderOddEven2(array []int) []int {
	return reorder(array, isEven)
}

type fn func(int) bool

func reorder(array []int, judgeFunc fn) []int {
	length := len(array)
	if length == 0 {
		return array
	}

	pLeft := 0           // 指向数组头
	pRight := length - 1 //  指向数组尾

	for pLeft < pRight {
		// 两个指针相遇后退出

		// 指向头的指针一直向右移，直到遇到第一个偶数
		for pLeft < pRight && !judgeFunc(array[pLeft]) {
			pLeft++
		}

		// 指向尾的指针一直向左移，直到遇到第一个奇数
		for pRight > pLeft && judgeFunc(array[pRight]) {
			pRight--
		}

		// 把两个指针遇到到的数值交换
		if pLeft < pRight {
			array[pLeft], array[pRight] = array[pRight], array[pLeft]
		}
	}
	return array
}

// 判断是否为奇数
func isEven(n int) bool {
	return (n & 0x1) == 0
}
