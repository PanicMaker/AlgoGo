package Combination

// https://leetcode.cn/problems/generate-parentheses/description/

func generateParenthesis(n int) []string {
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
