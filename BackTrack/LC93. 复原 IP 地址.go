package BackTrack

import (
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/restore-ip-addresses/

func restoreIpAddresses(s string) []string {
	n := len(s)

	ans := make([]string, 0)
	path := make([]string, 0)

	var dfs func(i int)
	dfs = func(i int) {
		if i == n && len(path) == 4 {
			ans = append(ans, strings.Join(path, "."))
			return
		}

		for j := i; j < n && j-i < 3; j++ {
			str := s[i : j+1]
			num, _ := strconv.Atoi(str)
			if num > 255 {
				continue
			}

			if j-i == 0 || (j-i > 0 && s[i] != '0') {
				path = append(path, str)
				dfs(j + 1)
				path = path[:len(path)-1]
			}
		}
	}

	dfs(0)
	return ans
}
