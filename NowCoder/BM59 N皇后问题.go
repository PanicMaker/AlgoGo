package nowcoder

// https://www.nowcoder.com/practice/c76408782512486d91eea181107293b6

func Nqueen(n int) int {
	ans := 0
	column := make([]int, n)

	isValid := func(row, col int) bool {
		for r := 0; r < row; r++ {
			l := column[r]
			if l == col || row+col == r+l || row-col == r-l {
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
