package nowcoder

func maxProfitBM82(prices []int) int {
	n := len(prices)

	const INF = -(1 << 30)

	dp := make([][5]int, n+1)
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[0][2] = INF
	dp[0][3] = INF
	dp[0][4] = INF

	// ·dp[i][o]表示到第天为止没有买过股票的最大收益
	// ·dp[i][1]表示到第天为止买过一次股票还没有卖出的最大收益
	// ·dp[i][2]表示到第天为止买过一次也卖出过一次股票的最大收益
	// ·dp[i][3]表示到第天为止买过两次只卖出过一次股票的最大收益
	// ·dp[i][4]表示到第天为止买过两次同时也买出过两次股票的最大收益
	for i, p := range prices {
		dp[i+1][1] = max(dp[i][1], dp[i][0]-p)
		dp[i+1][2] = max(dp[i][2], dp[i][1]+p)
		dp[i+1][3] = max(dp[i][3], dp[i][2]-p)
		dp[i+1][4] = max(dp[i][4], dp[i][3]+p)
	}

	return max(dp[n][0], dp[n][2], dp[n][4])
}
