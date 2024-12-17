package Array

// https://leetcode.cn/problems/rotate-image/

func rotate(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for row := 0; row < n; row++ {
		for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
			matrix[row][i], matrix[row][j] = matrix[row][j], matrix[row][i]
		}
	}
}
