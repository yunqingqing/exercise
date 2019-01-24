package day12

func NumberOf1Solution1(num int) int {
	count := 0
	flag := 1
	for flag > 0 {
		if num & flag != 0 {
			count++
		}
		flag = flag << 1
	}
	return count
}

func NumberOf1Solution2(num int) int {
	count := 0

	for num != 0 {
		count++
		num = (num -1) & num
	}
	return count
}