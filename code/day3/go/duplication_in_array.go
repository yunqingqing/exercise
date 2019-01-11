package main

import ("fmt")

func findDuplicationNumberInArray(numbers []int, length int) int {
	if (length <=0) {
		return -1
	}

	for i:=0; i<length; i++ {
		if (numbers[i] < 0 || numbers[i] > length -1) {
			return -1
		}
	}

	for i:=0; i<length; i++ {
		for i != numbers[i] {
			if numbers[i] == numbers[numbers[i]] {
				return numbers[i]
			}
			numbers[i], numbers[numbers[i]] = numbers[numbers[i]], numbers[i]
		}
	}

	return -1
}

func main () {
	var numbers = []int{2, 3, 1, 0, 5, 2, 3}
	fmt.Println("duplication number: ", findDuplicationNumberInArray(numbers, len(numbers)))
	var no_duplication_numbers = []int{2, 3, 1, 0, 5, 6, 4}
	fmt.Println("duplication number: ", findDuplicationNumberInArray(no_duplication_numbers, len(no_duplication_numbers)))
}