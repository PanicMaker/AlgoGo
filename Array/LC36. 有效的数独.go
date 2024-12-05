package Array

// https://leetcode.cn/problems/valid-sudoku/description/

func isValidSudoku(board [][]byte) bool {
	row := make([]map[int]bool, 9)
	for i := 0; i < 9; i++ {
		row[i] = make(map[int]bool)
	}

	col := make([]map[int]bool, 9)
	for i := 0; i < 9; i++ {
		col[i] = make(map[int]bool)
	}

	triangel := make(map[int]map[int]bool)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			if board[i][j] == '.' {
				continue
			}

			val := board[i][j] - '1'

			// 检测行
			if _, ok := row[i][int(val)]; ok {
				return false
			} else {
				row[i][int(val)] = true
			}

			// 检测列
			if _, ok := col[j][int(val)]; ok {
				return false
			} else {
				col[j][int(val)] = true
			}

			// 检测 3*3 九宫格
			tri := i/3*3 + j/3
			if _, ok := triangel[tri][int(val)]; ok {
				return false
			} else {
				if triangel[tri] == nil {
					triangel[tri] = make(map[int]bool)
				}
				triangel[tri][int(val)] = true
			}

		}
	}

	return true
}
