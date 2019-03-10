package day39

import (
	"fmt"
	"testing"
)

func TestFunc1(t *testing.T) {
	array := []int{4, 5, 1, 6, 2, 7, 3, 8}
	fmt.Println(GetLeastNumbers1(array, len(array), 4))
	fmt.Println("=============")
	fmt.Println(GetLeastNumbers2(array, len(array), 4))
	fmt.Println("======fun1 end=======")
}

func TestFunc2(t *testing.T) {
	array := []int{4, 5, 1, 6, 2, 7, 3, 8}
	fmt.Println(GetLeastNumbers1(array, len(array), len(array)))
	fmt.Println("=============")
	fmt.Println(GetLeastNumbers2(array, len(array), len(array)))
	fmt.Println("=======fun2 end======")
}

func TestFunc3(t *testing.T) {
	array := []int{4, 5, 1, 6, 2, 7, 3, 8}
	fmt.Println(GetLeastNumbers1(array, len(array), len(array)+1))
	fmt.Println("=============")
	fmt.Println(GetLeastNumbers2(array, len(array), len(array)+1))
	fmt.Println("=======fun3 end======")
}

func TestFunc4(t *testing.T) {
	array := []int{4, 5, 1, 6, 2, 7, 3, 8}
	fmt.Println(GetLeastNumbers1(array, len(array), 1))
	fmt.Println("=============")
	fmt.Println(GetLeastNumbers2(array, len(array), 1))
	fmt.Println("======fun4 end=======")
}
