package nowcoder

// https://www.nowcoder.com/practice/7a71a88cdf294ce6bdf54c899be967a2

func solveBM61I(matrix [][]int) int {
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

func solveBM61II(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])

	// dp[i][j] 表示从 (i,j) 出发的最长递增路径长度
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		// 如果已经计算过，直接返回
		if dp[i][j] > 0 {
			return dp[i][j]
		}

		// 初始默认为1（自己本身）
		maxLen := 1

		// 向四个方向探索
		dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			if x >= 0 && y >= 0 && x < m && y < n && matrix[x][y] < matrix[i][j] {
				// 如果相邻点的值更小，则可以构成递增路径
				currLen := 1 + dfs(x, y)
				if currLen > maxLen {
					maxLen = currLen
				}
			}
		}

		// 将结果存入 dp 数组
		dp[i][j] = maxLen
		return maxLen
	}

	// 遍历所有起点
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans = max(ans, dfs(i, j))
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
