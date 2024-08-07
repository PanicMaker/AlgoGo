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

		path = append(path, c)
		backtrack(i + 1)
		path = path[:len(path)-1]

		if c >= 65 && c <= 90 {
			path = append(path, c+32)
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
		if c >= 97 && c <= 122 {
			path = append(path, c-32)
			backtrack(i + 1)
			path = path[:len(path)-1]
		}

	}

	backtrack(0)

	return ans
}
