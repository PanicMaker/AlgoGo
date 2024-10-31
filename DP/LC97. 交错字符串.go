package DP

// https://leetcode.cn/problems/interleaving-string

func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	l := len(s3)

	if m+n != l {
		return false
	}

	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}

	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {

		if i < 0 {
			for idx := range s2[:j+1] {
				if s2[idx] != s3[idx] {
					return false
				}
			}
			return true
		}

		if j < 0 {
			for idx := range s1[:i+1] {
				if s1[idx] != s3[idx] {
					return false
				}
			}
			return true
		}

		if memo[i][j] != -1 {
			return memo[i][j] == 1
		}

		res := (dfs(i-1, j) && s1[i] == s3[i+j+1]) || (dfs(i, j-1) && s2[j] == s3[i+j+1])

		if res {
			memo[i][j] = 1
		} else {
			memo[i][j] = 0
		}

		return memo[i][j] == 1
	}

	return dfs(m-1, n-1)
}
