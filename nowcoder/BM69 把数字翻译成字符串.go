package nowcoder

import "strconv"

// https://www.nowcoder.com/practice/046a55e6cd274cffb88fc32dba695668

func solveBM69(nums string) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	if nums[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		twoDigit, _ := strconv.Atoi(nums[i-2 : i])

		// 单字数字，有效（非 '0'）
		if nums[i-1] != '0' {
			dp[i] += dp[i-1]
		}

		// 双位数字，有效的范围(10~26)
		if twoDigit >= 10 && twoDigit <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}
