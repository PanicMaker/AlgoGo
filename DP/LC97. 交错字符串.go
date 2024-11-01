package DP

// https://leetcode.cn/problems/interleaving-string

func isInterleave1(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	l := len(s3)

	if m+n != l {
		return false
	}

	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}

	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {

		if i < 0 {
			for idx := range s2[:j+1] {
				if s2[idx] != s3[idx] {
					return false
				}
			}
			return true
		}

		if j < 0 {
			for idx := range s1[:i+1] {
				if s1[idx] != s3[idx] {
					return false
				}
			}
			return true
		}

		if memo[i][j] != -1 {
			return memo[i][j] == 1
		}

		res := (dfs(i-1, j) && s1[i] == s3[i+j+1]) || (dfs(i, j-1) && s2[j] == s3[i+j+1])

		if res {
			memo[i][j] = 1
		} else {
			memo[i][j] = 0
		}

		return memo[i][j] == 1
	}

	return dfs(m-1, n-1)
}

func isInterleave2(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	l := len(s3)

	// 如果 s1 和 s2 的总长度不等于 s3 的长度，直接返回 false
	if m+n != l {
		return false
	}

	// 初始化 dp 数组
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	// 初始化 dp 的初始状态
	dp[0][0] = true

	// 填充第一列
	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}

	// 填充第一行
	for j := 1; j <= n; j++ {
		dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
	}

	// 填充剩余的 dp 数组
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i+1][j+1] = (dp[i][j+1] && s1[i] == s3[i+j+1]) || (dp[i+1][j] && s2[j] == s3[i+j+1])
		}
	}

	return dp[m][n]
}
