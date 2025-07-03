package graph

// https://leetcode.cn/problems/number-of-islands/description

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == '0' {
			return
		}

		grid[i][j] = '0'

		dfs(i, j+1)
		dfs(i, j-1)
		dfs(i-1, j)
		dfs(i+1, j)
	}

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				ans++
				dfs(i, j)
			}
		}
	}

	return ans
}
