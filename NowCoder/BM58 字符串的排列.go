package nowcoder

import "sort"

// https://www.nowcoder.com/practice/fe6b651b66ae47d7acce78ffdd9a96c7

func Permutation(str string) []string {
	n := len(str)
	if n == 0 {
		return []string{}
	}

	ans := make([]string, 0)
	path := make([]byte, 0)
	used := make([]bool, n)
	s := []byte(str)
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })

	var dfs func()
	dfs = func() {
		if len(path) == n {
			ans = append(ans, string(path))
			return
		}
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			if i > 0 && s[i] == s[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			path = append(path, s[i])
			dfs()
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	dfs()
	return ans
}
