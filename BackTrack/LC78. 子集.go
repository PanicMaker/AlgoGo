package BackTrack

import (
	"slices"
)

// https://leetcode.cn/problems/subsets/

func subsetsI(nums []int) [][]int {
	n := len(nums)

	ans := make([][]int, 0)
	path := make([]int, 0)

	var dfs func(i int)
	dfs = func(i int) {
		ans = append(ans, slices.Clone(path))

		for j := i; j < n; j++ {
			path = append(path, nums[j])
			dfs(j + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(0)

	return ans
}

func subsetsII(nums []int) [][]int {
	n := len(nums)

	ans := make([][]int, 0)
	path := make([]int, 0)

	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, slices.Clone(path))
			return
		}

		// 当前数字不选
		dfs(i + 1)

		// 选择当前数字
		path = append(path, nums[i])
		dfs(i + 1)
		path = path[:len(path)-1]
	}

	dfs(0)

	return ans
}
