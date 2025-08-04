package nowcoder

// https://www.nowcoder.com/practice/7d21b6be4c6b429bb92d219341c4f8bb

func minPathSum(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = matrix[0][0]

	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + matrix[i][0]
	}

	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + matrix[0][j]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + matrix[i][j]
		}
	}

	return dp[m-1][n-1]
}
