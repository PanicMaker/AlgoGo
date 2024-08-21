package Permutation

// https://leetcode.cn/problems/n-queens-ii/description/

func totalNQueens(n int) int {
	ans := 0
	column := make([]int, n)

	isValid := func(row int, col int) bool {
		for r := 0; r < row; r++ {
			c := column[r]
			if col == c || row-r == col-c || row-r == c-col {
				return false
			}
		}
		return true
	}

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			ans++
			return
		}

		for col := 0; col < n; col++ {
			if isValid(row, col) {
				column[row] = col
				backtrack(row + 1)
			}
		}
	}

	backtrack(0)

	return ans
}
