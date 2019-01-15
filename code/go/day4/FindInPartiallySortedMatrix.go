package main

import "fmt"

func find(matrix [][]int, rows, columns, number int) bool {
	if(matrix != nil && rows > 0 && columns >0) {
		row := 0
		column := columns - 1
		for row < rows && column >= 0 {
			if matrix[row][column] == number {
				return true
			} else if matrix[row][column] > number {
				column--
			} else {
				row++
			}
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}

	fmt.Println("8 in matrix: ", find(matrix, 4, 4, 8))
	fmt.Println("16 in matrix: ", find(matrix, 4, 4, 16))
	fmt.Println("1 in matrix: ", find(matrix, 4, 4, 1))
	fmt.Println("15 in matrix: ", find(matrix, 4, 4, 15))
}