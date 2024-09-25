package DP

// https://leetcode.cn/problems/target-sum/description/

// 递归
func findTargetSumWays1(nums []int, target int) int {
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

// 递归
func findTargetSumWays2(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	capacity := sum - abs(target)
	if capacity < 0 || capacity%2 == 1 {
		return 0
	}

	n := len(nums)

	var dfs func(i int, cur int) int
	dfs = func(i, cur int) int {
		if i < 0 {
			if cur == 0 {
				return 1
			}
			return 0
		}

		if cur < nums[i] {
			return dfs(i-1, cur)
		}

		return dfs(i-1, cur) + dfs(i-1, cur-nums[i])

	}

	return dfs(n-1, capacity/2)
}

// 递归 + 记忆化搜索
func findTargetSumWays3(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	capacity := sum - abs(target)
	if capacity < 0 || capacity%2 == 1 {
		return 0
	}

	n := len(nums)

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, capacity/2+1)
	}

	var dfs func(i int, cur int) int
	dfs = func(i, cur int) int {
		if i < 0 {
			if cur == 0 {
				return 1
			}
			return 0
		}

		if memo[i][cur] != 0 {
			return memo[i][cur]
		}

		if cur < nums[i] {
			return dfs(i-1, cur)
		}

		res := dfs(i-1, cur) + dfs(i-1, cur-nums[i])
		memo[i][cur] = res

		return res

	}

	return dfs(n-1, capacity/2)
}

func findTargetSumWays4(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	capacity := sum - abs(target)
	if capacity < 0 || capacity%2 == 1 {
		return 0
	}
	capacity /= 2

	n := len(nums)

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, capacity+1)
	}

	for i := 1; i < n; i++ {
		for j := capacity; j >= 0; j-- {
			if j < nums[i] {
				memo[i][j] = memo[i-1][j]
			} else {
				memo[i][j] = memo[i-1][j] + memo[i-1][j-nums[i]]
			}
		}
	}

	return memo[n-1][capacity]
}
