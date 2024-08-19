package Combination

// https://leetcode.cn/problems/generate-parentheses/description/

// 选或不选
func generateParenthesisI(n int) []string {
	m := 2 * n

	ans := make([]string, 0)
	path := make([]byte, 0)

	var backtrack func(i int, left int)
	backtrack = func(i int, left int) {
		if len(path) == m {
			ans = append(ans, string(path))
			return
		}

		if left < n {
			path = append(path, '(')
			backtrack(i+1, left+1)
			path = path[:len(path)-1]
		}

		if i-left < left {
			path = append(path, ')')
			backtrack(i+1, left)
			path = path[:len(path)-1]
		}
	}

	backtrack(0, 0)

	return ans
}

// 枚举答案
func generateParenthesisII(n int) []string {
	m := 2 * n

	ans := make([]string, 0)
	path := make([]byte, 0)

	var backtrack func(i int, left int)
	backtrack = func(i int, left int) {
		if left > n || i-left > left {
			return
		}

		if len(path) == m {
			ans = append(ans, string(path))
			return
		}

		for j := i; j < m; j++ {
			path = append(path, '(')
			backtrack(j+1, left+1)
			path = path[:len(path)-1]

			path = append(path, ')')
			backtrack(j+1, left)
			path = path[:len(path)-1]
		}
	}

	backtrack(0, 0)

	return ans
}
