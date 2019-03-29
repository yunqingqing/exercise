package day46

func GetMaxValue(values [][]int, rows, cols int) int {
	if values == nil || rows <= 0 || cols <= 0 {
		return 0
	}

	maxValues := make([][]int, rows)
	for i := 0; i < rows; i++ {
		maxValues[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			left := 0
			up := 0
			if i > 0 {
				up = maxValues[i - 1][j]
			}

			if j > 0 {
				left = maxValues[i][j - 1]
			}

			// f(i,j) = max(f(i-1, j), f(i,j-1)) + gift[i,j]
			maxValues[i][j] = max(left, up) + values[i][j]
		}
	}
	maxValue := maxValues[rows - 1][cols - 1]

	return maxValue
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}