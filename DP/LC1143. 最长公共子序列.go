package DP

// https://leetcode.cn/problems/longest-common-subsequence/description/

// 递归
func longestCommonSubsequence1(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)

	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}

	var dfs func(i, j int) (res int)
	dfs = func(i, j int) (res int) {

		if i < 0 || j < 0 {
			return 0
		}

		if memo[i][j] != -1 {
			return memo[i][j]
		}

		defer func() {
			memo[i][j] = res
		}()

		if text1[i] == text2[j] {
			return dfs(i-1, j-1) + 1
		}

		return max(dfs(i-1, j), dfs(i, j-1))
	}

	return dfs(m-1, n-1)
}

// 递推
func longestCommonSubsequence2(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)

	memo := make([][]int, m+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
		for j := range memo[i] {
			memo[i][j] = 0
		}
	}

	for i, x := range text1 {
		for j, y := range text2 {
			if x == y {
				memo[i+1][j+1] = memo[i][j] + 1
			} else {
				memo[i+1][j+1] = max(memo[i+1][j], memo[i][j+1])
			}
		}
	}

	return memo[m][n]
}

// 优化为一个数组
func longestCommonSubsequence3(text1 string, text2 string) int {
	m := len(text2)

	dp := make([]int, m+1)

	for _, x := range text1 {
		pre := 0
		for j, y := range text2 {
			if x == y {
				dp[j+1], pre = pre+1, dp[j]
			} else {
				pre, dp[j+1] = dp[j+1], max(dp[j], dp[j+1])
			}
		}
	}

	return dp[m]
}
