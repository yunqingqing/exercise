package day15

import "fmt"

func printNumber(number []int) {
	isBeginning := true
	nLength := len(number)

	for i := 0; i < nLength; i++ {
		if isBeginning && number[i] != 0 {
			isBeginning = false
		}
		if !isBeginning {
			fmt.Printf("%d", number[i])
		}
	}
	fmt.Printf("\n")
}

func Print1ToMaxOfDigits(n int) {
	if n <= 0 {
		return
	}
	number := make([]int, n)

	for i := 0; i < 10; i++ {
		number[0] = i
		Print1ToMaxOfNDigitsRecursively(number, n, 0)
	}

}

func Print1ToMaxOfNDigitsRecursively(number []int, length, index int) {
	if index == length - 1{
		printNumber(number)
		return
	}

	for i:= 0; i<10; i++ {
		number[index + 1] = i
		Print1ToMaxOfNDigitsRecursively(number, length, index + 1)
	}
}


