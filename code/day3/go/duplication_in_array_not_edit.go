package main

import "fmt"

func countRange(numbers []int, length int, start, end int) int {
	count := 0
	for _, num := range numbers {
		if (num >= start && num <= end) {
			count++
		}
	}
	return count
}

func findDuplicationNumberInArray(numbers []int, length int) int {
	if (length <= 0) {
		return -1
	}

	for i:=0; i < length; i++ {
		if (numbers[i] < 0 || numbers[i] > length -1) {
			return -1
		}
	}

	start := 1
	end := length - 1
	for start <= end {
		middle := ((end - start) >> 1) + start
		count := countRange(numbers, length, start, middle)
		if end == start {
			if count > 1 {
				return start
			} else {
				break
			}
		}

		if count > (middle- start +1) {
			end = middle
		} else {
			start = middle + 1
		}
	}
	return -1
}

func main() {
	var numbers = []int{2, 3, 1, 6, 5, 2, 3}
	fmt.Println("duplication number: ", findDuplicationNumberInArray(numbers, len(numbers)))
	var no_duplication_numbers = []int{2, 3, 1, 7, 5, 6, 4}
	fmt.Println("duplication number: ", findDuplicationNumberInArray(no_duplication_numbers, len(no_duplication_numbers)))
}