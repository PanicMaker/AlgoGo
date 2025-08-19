package nowcoder

func generateParenthesis(n int) []string {
	var ans []string
	var backtrack func(path string, left, right int)
	backtrack = func(path string, left, right int) {
		if len(path) == 2*n {
			ans = append(ans, path)
			return
		}
		if left < n {
			backtrack(path+"(", left+1, right)
		}
		if right < left {
			backtrack(path+")", left, right+1)
		}
	}
	backtrack("", 0, 0)
	return ans
}
