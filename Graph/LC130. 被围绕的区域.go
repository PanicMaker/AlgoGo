package graph

// https://leetcode.cn/problems/surrounded-regions/

func solve(board [][]byte) {
	m, n := len(board), len(board[0])

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= m || j >= n || board[i][j] != 'O' {
			return
		}
		board[i][j] = '-' // 标记为未被包围
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}

	// 处理边界
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			dfs(i, 0)
		}
		if board[i][n-1] == 'O' {
			dfs(i, n-1)
		}
	}
	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			dfs(0, j)
		}
		if board[m-1][j] == 'O' {
			dfs(m-1, j)
		}
	}

	// 翻转和恢复
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch board[i][j] {
			case 'O':
				board[i][j] = 'X'
			case '-':
				board[i][j] = 'O'
			}
		}
	}
}
