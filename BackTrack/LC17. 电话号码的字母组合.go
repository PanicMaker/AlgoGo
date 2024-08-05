package BackTrack

// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/

func letterCombinations(digits string) []string {
	dict := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	ans := make([]string, 0)

	n := len(digits)
	if n == 0 {
		return ans
	}
	path := make([]byte, 0)

	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, string(path))
			return
		}

		for _, c := range dict[digits[i]-'0'] {
			path = append(path, byte(c))
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(0)

	return ans
}
