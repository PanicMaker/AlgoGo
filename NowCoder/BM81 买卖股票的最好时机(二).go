package nowcoder

// https://www.nowcoder.com/practice/9e5e3c2603064829b0a0bbfca10594e9

func maxProfit(prices []int) int {
	n := len(prices)

	dp := make([][2]int, n+1)
	dp[0][1] = -prices[0]

	for i, p := range prices {
		dp[i+1][0] = max(dp[i][0], dp[i][1]+p)
		dp[i+1][1] = max(dp[i][1], dp[i][0]-p)
	}

	return dp[n][0]
}
