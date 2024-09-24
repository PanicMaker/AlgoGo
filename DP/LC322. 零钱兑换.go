package DP

// https://leetcode.cn/problems/coin-change/description/

func coinChange1(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	n := len(coins)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, amount+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(i int, target int) int
	dfs = func(i, target int) int {
		if i < 0 {
			if target == 0 {
				return 0
			}
			return amount + 1
		}

		if memo[i][target] != -1 {
			return memo[i][target]
		}

		if target < coins[i] {
			return dfs(i-1, target)
		}

		res := min(dfs(i-1, target), dfs(i, target-coins[i])+1)
		memo[i][target] = res

		return res
	}

	ans := dfs(n-1, amount)

	if ans < amount+1 {
		return ans
	}

	return -1
}

func coinChange2(coins []int, amount int) int {
	memo := make([]int, amount+1)
	// 备忘录初始化为一个不会被取到的特殊值，代表还未被计算
	for i := range memo {
		memo[i] = -100
	}

	var dfs func(target int) int
	dfs = func(target int) int {
		if target == 0 {
			return 0
		}

		if target < 0 {
			return -1
		}

		// 查备忘录，防止重复计算
		if memo[target] != -100 {
			return memo[target]
		}

		res := amount + 1
		for _, coin := range coins {
			result := dfs(target - coin)

			if result == -1 {
				continue
			}

			res = min(res, result+1)
		}

		// 把计算结果存入备忘录
		if res == amount+1 {
			memo[target] = -1
		} else {
			memo[target] = res
		}

		return memo[target]
	}

	return dfs(amount)
}

func coinChange3(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i < amount+1; i++ {
		dp[i] = amount + 1
	}

	for i := 1; i < amount+1; i++ {
		for _, coin := range coins {
			if i >= coin {
				dp[i] = min(dp[i], dp[i-coin]+1)

			}
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}
