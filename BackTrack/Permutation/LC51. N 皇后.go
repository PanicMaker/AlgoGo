package Permutation

import "strings"

// https://leetcode.cn/problems/n-queens/description/

func solveNQueensI(n int) [][]string {
	ans := make([][]string, 0)
	column := make([]int, n)

	isValid := func(row int, col int) bool {
		for Row := 0; Row < row; Row++ {
			Col := column[Row]
			if row+col == Row+Col || row-col == Row-Col {
				return false
			}
		}
		return true
	}

	var backtrack func(row int, set []int)
	backtrack = func(row int, set []int) {
		if row == n {
			tmp := make([]string, 0)
			for _, c := range column {
				str := strings.Repeat(".", c) + "Q" + strings.Repeat(".", n-c-1)
				tmp = append(tmp, str)
			}

			ans = append(ans, tmp)
			return
		}

		for i, c := range set {
			if isValid(row, c) {
				column[row] = c
				backtrack(row+1, append(set[:i], set[i+1:]...))
			}
		}
	}

	set := make([]int, n)
	for i, _ := range set {
		set[i] = i
	}

	backtrack(0, set)

	return ans
}

func solveNQueensII(n int) [][]string {
	ans := make([][]string, 0)
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
			tmp := make([]string, n)
			for r := 0; r < n; r++ {
				str := strings.Repeat(".", column[r]) + "Q" + strings.Repeat(".", n-column[r]-1)
				tmp[r] = str
			}
			ans = append(ans, tmp)
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
