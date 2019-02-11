package day28

import "fmt"

func PrintMatrixClockwisely(matrix [][]int, columns, rows int) {
	if matrix == nil || columns <= 0 || rows <= 0 {
		return
	}

	start := 0

	for columns > start * 2 && rows > start * 2 {
		PrintMatrixInCircle(matrix, columns, rows, start)
		// 下一圈起始坐标
		start++
	}
}

// 循环打印一圈
func PrintMatrixInCircle(matrix [][]int, columns, rows, start int) {
	var number int
	endX := columns - 1 - start
	endY := rows - 1 - start

	// 从左到右打印一行
	for i := start; i <= endX; i++ {
		number = matrix[start][i]
		fmt.Println(number)
	}

	// 从上到下打印一列
	if start < endY {
		for i := start + 1; i <= endY; i++ {
			number = matrix[i][endX]
			fmt.Println(number)
		}
	}

	// 从右到左打印一行
	if start < endX && start < endY {
		for i := endX - 1; i >= start; i-- {
			number = matrix[endY][i]
			fmt.Println(number)
		}
	}

	// 从下到上打印一行
	if start < endX && start < endY - 1 {
		for i := endY - 1; i >= start + 1; i-- {
			number = matrix[i][start]
			fmt.Println(number)
		}
	}
}