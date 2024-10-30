package DP

// https://leetcode.cn/problems/edit-distance/description/

// 递归
func minDistance1(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}

	var dfs func(i, j int) (res int)
	dfs = func(i, j int) (res int) {

		if i < 0 {
			return j + 1
		}

		if j < 0 {
			return i + 1
		}

		if memo[i][j] != -1 {
			return memo[i][j]
		}

		defer func() {
			memo[i][j] = res
		}()

		if word1[i] == word2[j] {
			return dfs(i-1, j-1)
		}

		// dfs(i-1, j) 插入一个字母
		// dfs(i, j-1) 删除一个字母
		// dfs(i-1, j-1)) 替换一个字母
		return min(dfs(i-1, j), dfs(i, j-1), dfs(i-1, j-1)) + 1
	}

	return dfs(m-1, n-1)
}

// 递推
func minDistance2(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}

	for i, x := range word1 {
		for j, y := range word2 {
			if x == y {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = min(dp[i][j+1], dp[i+1][j], dp[i][j]) + 1
			}
		}
	}

	return dp[m][n]
}
