package nowcoder

// https://www.nowcoder.com/practice/6fe0302a058a4e4a834ee44af88435c7

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	dp[0] = 0
	dp[1] = 0
	dp[2] = min(cost[0], dp[1]+cost[1])

	for i := 3; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[len(cost)]
}
