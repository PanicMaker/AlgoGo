package Array

// https://leetcode.cn/problems/set-matrix-zeroes/description/

func setZeroesI(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	newMatrix := make([][]int, len(matrix))
	for i, row := range matrix {
		newMatrix[i] = make([]int, len(row))
		copy(newMatrix[i], row)
	}

	row := make([]int, m)
	col := make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 && newMatrix[i][j] == 0 {
				if row[i] != 1 {
					for x := 0; x < n; x++ {
						matrix[i][x] = 0
					}
					row[i] = 1
				}

				if col[j] != 1 {
					for y := 0; y < m; y++ {
						matrix[y][j] = 0
					}
					col[j] = 1
				}
			}
		}
	}
}

func setZeroesII(matrix [][]int) {
	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))
	for i, r := range matrix {
		for j, v := range r {
			if v == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i, r := range matrix {
		for j := range r {
			if row[i] || col[j] {
				r[j] = 0
			}
		}
	}
}
