package DP

// https://leetcode.cn/problems/coin-change-ii/description/

func change1(amount int, coins []int) int {
	n := len(coins)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1
	}
	for j := range dp[0] {
		dp[0][j] = 0
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < amount+1; j++ {
			if j < coins[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			}
		}
	}

	return dp[n][amount]
}

func change2(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for j := 1; j < amount+1; j++ {
			if j < coin {
				dp[j] = dp[j]
			} else {
				dp[j] = dp[j] + dp[j-coin]
			}
		}
	}

	return dp[amount]
}
