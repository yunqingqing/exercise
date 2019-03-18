package day42

func NumberOf1Between1AndN(num int) int {
	if num <= 0 {
		return 0
	}
	number := 0
	for i := 1; i <= num; i++ {
		number += NumberOf1(i)
	}
	return number
}

func NumberOf1(num int) int {
	number := 0
	for num != 0 {
		if num%10 == 1 {
			number++
		}
		num = num / 10
	}
	return number
}
