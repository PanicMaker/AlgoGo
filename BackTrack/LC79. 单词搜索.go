package BackTrack

// https://leetcode.cn/problems/word-search/

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	length := len(word)
	// 是否访问过该字符
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	// 是否存在该字符串
	found := false

	var backtrack func(i int, j int, idx int)
	backtrack = func(i int, j int, idx int) {
		if found {
			return
		}

		if idx == length {
			found = true
			return
		}

		// 处理边界
		if i < 0 || j < 0 || i >= m || j >= n {
			return
		}

		if board[i][j] != word[idx] {
			return
		}

		// 已经访问过就不选择，不可重复选
		if visited[i][j] {
			return
		}

		visited[i][j] = true
		backtrack(i-1, j, idx+1)
		backtrack(i+1, j, idx+1)
		backtrack(i, j-1, idx+1)
		backtrack(i, j+1, idx+1)
		visited[i][j] = false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			backtrack(i, j, 0)
		}
	}

	return found
}
