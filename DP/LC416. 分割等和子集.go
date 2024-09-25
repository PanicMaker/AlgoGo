package DP

// https://leetcode.cn/problems/partition-equal-subset-sum/description/

// 递归搜索 + 保存计算结果 = 记忆化搜索, 会超时
func canPartition1(nums []int) bool {
	n := len(nums)

	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%2 != 0 {
		return false
	}

	memo := make([][]bool, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]bool, sum/2+1)
	}

	var dfs func(i int, target int) bool
	dfs = func(i, target int) bool {
		if i < 0 {
			return target == 0
		}

		if memo[i][target] {
			return true
		}

		if target < nums[i] {
			return dfs(i-1, target)
		}

		res := dfs(i-1, target) || dfs(i-1, target-nums[i])
		memo[i][target] = res

		return res
	}

	return dfs(n-1, sum/2)
}

func canPartition2(nums []int) bool {
	n := len(nums)

	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2

	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}

	for i := 0; i <= n; i++ {
		dp[i][0] = true
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < target+1; j++ {
			if j < nums[i-1] {
				// 背包容量不足，不能装入第 i 个物品
				dp[i][j] = dp[i-1][j]
			} else {
				// 装入或不装入背包
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}

	return dp[n][target]
}

func canPartition3(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2

	dp := make([]bool, target+1)

	dp[0] = true

	for _, num := range nums {
		for j := target; j >= 0; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}

	return dp[target]
}
