package nowcoder

// https://www.nowcoder.com/practice/6d29638c85bb4ffd80c020fe244baf11

func LCS(s1 string, s2 string) string {
	m, n := len(s1), len(s2)

	dp := make([][]int, m+1)
	dir := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dir[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if s1[i] == s2[j] {
				dp[i+1][j+1] = 1 + dp[i][j]
				dir[i+1][j+1] = 2
			} else {
				if dp[i][j+1] > dp[i+1][j] {
					dp[i+1][j+1] = dp[i][j+1]
					dir[i+1][j+1] = 3
				} else {
					dp[i+1][j+1] = dp[i+1][j]
					dir[i+1][j+1] = 1
				}
			}
		}
	}

	var res []byte
	i, j := m, n
	for i > 0 && j > 0 {
		switch dir[i][j] {
		case 1:
			j--
		case 2:
			res = append(res, s1[i-1])
			i--
			j--
		case 3:
			i--
		}
	}

	if len(res) == 0 {
		return "-1"
	}

	for left, right := 0, len(res)-1; left < right; left, right = left+1, right-1 {
		res[left], res[right] = res[right], res[left]
	}

	return string(res)
}
