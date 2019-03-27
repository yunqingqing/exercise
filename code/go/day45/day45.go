package day45

import "strconv"

func GetTranslationCount(number int) int {
	if number < 0 {
		return 0
	}
	stringNumber := strconv.Itoa(number)
	return getTranslationCount(stringNumber)
}

func getTranslationCount(strNumber string) int {
	length := len(strNumber)
	counts := make([]int, length)  // 用来存放计算过的情况的次数
	count := 0

	for i := length - 1; i >= 0; i-- {
		count = 0
		if i < (length - 1) {
			count = counts[i + 1]
		} else {
			count = 1
		}
		if i < (length - 1) {
			digit1, _ := strconv.Atoi(string([]rune(strNumber)[i]))
			digit2, _ := strconv.Atoi(string([]rune(strNumber)[i+1]))
			converted := digit1 * 10 + digit2
			if converted >= 10 && converted <= 25 {
				if i < (length - 2) {
					count += counts[i + 2]
				} else {
					count += 1
				}
			}
		}
		counts[i] = count

	}
	count = counts[0]

	return count
}