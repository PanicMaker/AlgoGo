package nowcoder

// https://www.nowcoder.com/practice/7a71a88cdf294ce6bdf54c899be967a2

func solveBM61(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])

	ans := 0
	path := make([][2]int, 0)
	used := make(map[[2]int]bool)

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= m || j >= n {
			return
		}

		if used[[2]int{i, j}] {
			return
		}

		if len(path) > 0 {
			last := path[len(path)-1]
			if matrix[i][j] <= matrix[last[0]][last[1]] {
				return
			}
		}

		node := [2]int{i, j}
		path = append(path, node)
		used[node] = true

		ans = max(ans, len(path))

		dfs(i-1, j)
		dfs(i, j-1)
		dfs(i+1, j)
		dfs(i, j+1)

		path = path[:len(path)-1]
		used[node] = false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j)
		}
	}

	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return b
}
