package nowcoder

// https://www.nowcoder.com/practice/2e95333fbdd4451395066957e24909cc

func rotateMatrix(mat [][]int, n int) [][]int {
	// Step 1: 转置矩阵 (行变列)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			mat[i][j], mat[j][i] = mat[j][i], mat[i][j]
		}
	}
	// Step 2: 每一行反转（左右翻转）
	for i := 0; i < n; i++ {
		for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
			mat[i][l], mat[i][r] = mat[i][r], mat[i][l]
		}
	}
	return mat
}
