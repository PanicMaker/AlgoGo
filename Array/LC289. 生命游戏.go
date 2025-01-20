package Array

// https://leetcode.cn/problems/game-of-life/

func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])
	newBoard := make([][]int, len(board))

	for i, row := range board {
		newBoard[i] = make([]int, len(row))
		copy(newBoard[i], row)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			count := countOnesAround(newBoard, i, j)
			if count < 2 {
				board[i][j] = 0
			} else if count == 2 {
				if newBoard[i][j] == 1 {
					continue
				}
			} else if count == 3 {
				if newBoard[i][j] == 1 {
					continue
				}
				if newBoard[i][j] == 0 {
					board[i][j] = 1
				}
			} else {
				board[i][j] = 0
			}
		}
	}
}

func countOnesAround(grid [][]int, x, y int) int {
	// 定义方向数组，表示周围8个方向的偏移量
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1}, // 上一行
		{0, -1} /* 当前行 */, {0, 1}, // 当前行
		{1, -1}, {1, 0}, {1, 1}, // 下一行
	}

	// 初始化计数器
	count := 0

	// 遍历周围8个方向
	for _, dir := range directions {
		nx, ny := x+dir[0], y+dir[1]
		// 检查是否在表格范围内
		if nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid[0]) {
			// 如果该位置的值为1，计数器加1
			if grid[nx][ny] == 1 {
				count++
			}
		}
	}

	// 返回计数结果
	return count
}
