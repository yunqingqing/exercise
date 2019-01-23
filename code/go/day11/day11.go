package day11

func movingCount(k, rows, cols int) int {
	startRow := 0
	startCol := 0
	visited := []bool{} // 标记是否访问过

	// 初始化数组
	for i := 0; i < rows*cols; i++ {
		visited = append(visited, false)
	}

	// 回溯法求解
	count := movingCountCore(k, rows, cols, startRow, startCol, visited)

	return count
}

func movingCountCore(k, rows, cols, row, col int, visited []bool) int {
	count := 0
	if check(k, rows, cols, row, col, visited) {
		visited[row*cols+col] = true

		// 进入下一歌
		count = 1 + movingCountCore(k, rows, cols, row-1, col, visited) + movingCountCore(k, rows, cols, row, col-1, visited) + movingCountCore(k, rows, cols, row+1, col, visited) + movingCountCore(k, rows, cols, row, col+1, visited)
	}

	return count
}

// 判断是否能进入下一格
func check(k, rows, cols, row, col int, visited []bool) bool {
	if row >= 0 && row < rows && col >= 0 && col < cols && getSum(row)+getSum(col) <= k && !visited[row*cols+col] {
		return true
	}
	return false
}

// 求number数位和
// number=12 return 1+2=3
func getSum(number int) int {
	sum := 0
	for number > 0 {
		sum += number % 10
		number /= 10
	}

	return sum
}
