package day28

import (
	"testing"
	"fmt"
)

func TestFunc1(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	PrintMatrixClockwisely(matrix, 4, 4)
}

func TestFunc2(t *testing.T) {
	matrix := [][]int{
		{1},
	}
	fmt.Printf("test2: \n")
	PrintMatrixClockwisely(matrix, 1, 1)
}

func TestFunc3(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{5, 6, 7},
		{9, 10, 11},
	}

	fmt.Printf("test3: \n")
	PrintMatrixClockwisely(matrix, 3, 3)
}

func TestFunc4(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{5, 6, 7},
		{9, 10, 11},
		{12, 13, 14},
	}

	fmt.Printf("test4: \n")
	PrintMatrixClockwisely(matrix, 3, 4)
}