package DP

import "math"

// https://leetcode.cn/problems/length-of-the-longest-subsequence-that-sums-to-target/description/

func lengthOfLongestSubsequence1(nums []int, target int) int {
	n := len(nums)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, target+1)
		for j := range dp[i] {
			dp[i][j] = math.MinInt
		}
	}
	for i := range dp {
		dp[i][0] = 0
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < target+1; j++ {
			if j < nums[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i-1]]+1)
			}
		}
	}

	if dp[n][target] < 0 {
		return -1
	}

	return dp[n][target]
}

func lengthOfLongestSubsequence2(nums []int, target int) int {
	dp := make([]int, target+1)
	for i := range dp {
		dp[i] = math.MinInt
	}
	dp[0] = 0

	for _, num := range nums {
		for j := target; j >= num; j-- {
			dp[j] = max(dp[j], dp[j-num]+1)
		}
	}

	if dp[target] < 0 {
		return -1
	}

	return dp[target]
}
