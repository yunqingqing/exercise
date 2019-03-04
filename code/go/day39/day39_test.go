package day39

import (
	"fmt"
	"testing"
)

func TestFunc1(t *testing.T) {
	array := []int{4, 5, 1, 6, 2, 7, 3, 8}
	fmt.Println(GetLeastNumbers(array, len(array), 4))
}

func TestFunc2(t *testing.T) {
	array := []int{4, 5, 1, 6, 2, 7, 3, 8}
	fmt.Println(GetLeastNumbers(array, len(array), len(array)))
}

func TestFunc3(t *testing.T) {
	array := []int{4, 5, 1, 6, 2, 7, 3, 8}
	fmt.Println(GetLeastNumbers(array, len(array), len(array)+1))
}

func TestFunc4(t *testing.T) {
	array := []int{4, 5, 1, 6, 2, 7, 3, 8}
	fmt.Println(GetLeastNumbers(array, len(array), 1))
}
