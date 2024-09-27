package DP

import "math"

// https://leetcode.cn/problems/perfect-squares/description/

func numSquares1(n int) int {
	m := int(math.Sqrt(float64(n)))
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[0][j] = n + 1
		}
	}
	dp[0][0] = 0

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if j < i*i {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-i*i]+1)
			}
		}
	}

	return dp[m][n]
}

func numSquares2(n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = n + 1
	}
	dp[0] = 0

	for i := 1; i < n+1; i++ {
		for j := i * i; j < n+1; j++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
	}

	return dp[n]
}
