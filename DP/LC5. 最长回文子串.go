package DP

// https://leetcode.cn/problems/longest-palindromic-substring/description/

// 中心扩散
func longestPalindromeI(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	maxLen := 0
	startIndex := 0
	// 辅助函数：从某个中心点开始向两边扩展
	expandAroundCenter := func(left, right int) {
		for left >= 0 && right < n && s[left] == s[right] {
			currLen := right - left + 1
			if currLen > maxLen {
				maxLen = currLen
				startIndex = left
			}
			left--
			right++
		}
	}
	for i := 0; i < n; i++ {
		// 奇数长度的回文（以 s[i] 为中心）
		expandAroundCenter(i, i)
		// 偶数长度的回文（以 s[i], s[i+1] 为中心）
		expandAroundCenter(i, i+1)
	}
	return s[startIndex : startIndex+maxLen]
}

// 动态规划
func longestPalindromeII(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	dp := make([][]bool, n)
	for i := range n {
		dp[i] = make([]bool, n)
	}

	maxLen := 1
	startIndex := 0

	for i := range n {
		dp[i][i] = true
	}

	for l := 2; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			j := i + l - 1

			if s[i] == s[j] {
				if l == 2 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] && l > maxLen {
				maxLen = l
				startIndex = i
			}
		}
	}

	return s[startIndex : startIndex+maxLen]
}
