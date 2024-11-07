package DP

import "math"

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/description/

func maxProfit1(prices []int) int {
	n := len(prices)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, 2)
	}

	var dfs func(i, hold int) (res int)
	dfs = func(i, hold int) (res int) {
		if i < 0 {
			if hold == 1 {
				return math.MinInt
			}
			return res
		}

		if memo[i][hold] != 0 {
			return memo[i][hold]
		}

		defer func() {
			memo[i][hold] = res
		}()

		if hold == 1 {
			return max(dfs(i-1, 1), dfs(i-1, 0)-prices[i])
		}

		return max(dfs(i-1, 0), dfs(i-1, 1)+prices[i])
	}

	return dfs(len(prices)-1, 0)
}

func maxProfit2(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n+1)
	dp[0][1] = -prices[0]

	for i, p := range prices {
		dp[i+1][0] = max(dp[i][0], dp[i][1]+p)
		dp[i+1][1] = max(dp[i][1], dp[i][0]-p)
	}

	return dp[n][0]
}
