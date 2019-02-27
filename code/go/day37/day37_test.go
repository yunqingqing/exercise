package day37

import (
	"testing"
	"fmt"
)

func TestFunc(t *testing.T) {
	array := []int{4, 3, 8, 6, 11}
	QuickSort(&array, len(array), 0, len(array)-1)
	fmt.Println("sorted array.", array)
}