package day20

import (
	"fmt"
	"testing"
)

func TestFunc1(t *testing.T) {
	array1 := []int{2, 5, 1, 4}
	fmt.Println("re order array", RecorderOddEven1(array1))
	array2 := []int{1}
	fmt.Println("re order array", RecorderOddEven1(array2))
	array3 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("re order array", RecorderOddEven1(array3))
}

func TestFunc2(t *testing.T) {
	array1 := []int{2, 5, 1, 4}
	fmt.Println("method2 re order array", RecorderOddEven2(array1))
	array2 := []int{1}
	fmt.Println("method2 re order array", RecorderOddEven2(array2))
	array3 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("method2 re order array", RecorderOddEven2(array3))
}
