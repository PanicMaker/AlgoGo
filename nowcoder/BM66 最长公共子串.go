package nowcoder

// https://www.nowcoder.com/practice/f33f5adc55f444baa0e0ca87ad8a6aac

func LCS66(str1 string, str2 string) string {
	m, n := len(str1), len(str2)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	maxLen := 0
	index := -1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if str1[i] == str2[j] {
				dp[i+1][j+1] = 1 + dp[i][j]
				if dp[i+1][j+1] > maxLen {
					maxLen = dp[i+1][j+1]
					index = i
				}
			} else {
				dp[i+1][j+1] = 0
			}
		}
	}

	return str1[index-maxLen : index]
}
