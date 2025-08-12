package DP

import "slices"

// https://leetcode.cn/problems/maximum-number-of-events-that-can-be-attended-ii/description/

func maxValue(events [][]int, k int) int {
	n := len(events)

	// 按会议结束时间升序排序，方便后续二分查找
	slices.SortFunc(events, func(a, b []int) int {
		return a[1] - b[1]
	})

	// dp[i][j] 表示前 i 个会议中，最多参加 j 个会议的最大价值
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}

	// 初始化：参加 0 个会议时价值为 0
	for i := range dp {
		dp[i][0] = 0
	}
	for i := range dp[0] {
		dp[0][i] = 0
	}

	// 二分查找：找到最后一个结束时间 < 当前会议开始时间的会议编号
	search := func(event []int) int {
		left, right := 0, n-1
		for left < right {
			mid := left + (right-left)/2
			if events[mid][1] < event[0] {
				left = mid + 1
			} else {
				right = mid
			}
		}
		return left
	}

	// 动态规划转移
	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			// p 表示不冲突的上一个会议编号
			p := search(events[i-1])
			// 不选当前会议 or 选当前会议
			dp[i][j] = max(dp[i-1][j], dp[p][j-1]+events[i-1][2])
		}
	}

	return dp[n][k]
}
