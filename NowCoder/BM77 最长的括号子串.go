package nowcoder

// https://www.nowcoder.com/practice/45fd68024a4c4e97a8d6c45fc61dc6ad

func longestValidParentheses1(s string) int {
	stack := make([]int, 0)

	res := 0
	length := 0
	start := -1

	for i, v := range s {
		if v == '(' {
			stack = append(stack, i)
		} else {
			if len(stack) == 0 {
				start = i
			} else {
				index := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if len(stack) > 0 {
					length = i - index
				} else {
					length = i - start
				}
			}
		}

		res = max(res, length)
	}

	return res
}

func longestValidParentheses2(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	maxLen := 0

	for i := 1; i <= n; i++ {
		if s[i-1] == ')' {
			if i-2 >= 0 && s[i-2] == '(' {
				dp[i] = dp[i-2] + 2
			}
			if i-2 >= 0 && s[i-2] == ')' {
				prevIndex := i - 1 - dp[i-1] - 1
				if prevIndex >= 0 && s[prevIndex] == '(' {
					dp[i] = dp[i-1] + 2
					if prevIndex > 0 {
						dp[i] += dp[prevIndex]
					}
				}
			}
		}

		maxLen = max(maxLen, dp[i])
	}

	return maxLen
}

func longestValidParentheses3(s string) int {
	n := len(s)
	maxLen := 0
	left, right := 0, 0

	for i := range s {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			maxLen = max(maxLen, 2*right)
		}

		if right > left {
			left, right = 0, 0
		}
	}

	left, right = 0, 0

	for i := n - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			maxLen = max(maxLen, 2*right)
		}

		if left > right {
			left, right = 0, 0
		}
	}

	return maxLen
}
