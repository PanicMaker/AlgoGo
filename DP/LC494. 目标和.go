package DP

// https://leetcode.cn/problems/target-sum/description/

func findTargetSumWays(nums []int, target int) int {
	n := len(nums)
	res := 0

	var dfs func(i int, cur int)
	dfs = func(i, cur int) {
		if i == n {
			if cur == target {
				res++
			}
			return
		}

		dfs(i+1, cur+nums[i])
		dfs(i+1, cur-nums[i])

	}

	dfs(0, 0)

	return res
}
