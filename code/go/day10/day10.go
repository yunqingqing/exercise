package day10

func HasPath(matrix string, rows, cols int, str string) bool {
	// 非法输入返回false
	if matrix == "" || rows < 1 || cols < 1 || str == "" {
		return false
	}

	visited := make([]bool, rows * cols)
	
	pathLength := 0

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// 起始点
			if hasPathCore(matrix, rows, cols, row, col, pathLength, str, visited) {
				return true
			}
		}
	}
	return false
}

func hasPathCore(matrix string, rows, cols, row, col, pathLength int, str string, visited []bool) bool {
	hasPath := false

	// 长度等于查找字符长度时即推出
	if pathLength == len(str) {
		return true
	}

	if row >= 0 && row < rows && col >= 0 && col < cols && matrix[row*cols+col] == str[pathLength] && !visited[row*cols+col] {
		pathLength++
		visited[row*cols+col] = true

		// 向上下左右前进一步，匹配成功则为true
		hasPath = hasPathCore(matrix, rows, cols, row, col-1, pathLength, str, visited) || hasPathCore(matrix, rows, cols, row-1, col, pathLength, str, visited) || hasPathCore(matrix, rows, cols, row, col+1, pathLength, str, visited) || hasPathCore(matrix, rows, cols, row+1, col, pathLength, str, visited)

		// 向前一步没有匹配的路径，失败则退回
		if !hasPath {
			pathLength--
			visited[row * cols + col] = false
		}
	}

	return hasPath
}
