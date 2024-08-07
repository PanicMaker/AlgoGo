package BackTrack

// https://leetcode.cn/problems/letter-case-permutation/description/

func letterCasePermutation(s string) []string {
	n := len(s)

	ans := make([]string, 0)
	path := make([]byte, 0)

	var backtrack func(i int)
	backtrack = func(i int) {
		if i == n {
			tmp := make([]byte, len(path))
			copy(tmp, path)
			ans = append(ans, string(tmp))
			return
		}

		c := s[i]

		// 直接添加当前字符
		path = append(path, c)
		backtrack(i + 1)
		path = path[:len(path)-1]

		// 如果是字母，考虑大小写转换
		if c >= 'A' && c <= 'Z' {
			path = append(path, c+32)
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
		if c >= 'a' && c <= 'z' {
			path = append(path, c-32)
			backtrack(i + 1)
			path = path[:len(path)-1]
		}

	}

	backtrack(0)

	return ans
}
