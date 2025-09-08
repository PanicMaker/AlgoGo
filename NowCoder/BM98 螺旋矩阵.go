package nowcoder

// https://www.nowcoder.com/practice/7edf70f2d29c4b599693dc3aaeea1d31

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	m, n := len(matrix), len(matrix[0])
	left, right := 0, n-1
	top, bottom := 0, m-1
	res := make([]int, 0)
	for left <= right && top <= bottom {
		// ← 上行：从左到右
		for j := left; j <= right; j++ {
			res = append(res, matrix[top][j])
		}
		top++
		// ↓ 右列：从上到下
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		// → 下行：从右到左
		if top <= bottom {
			for j := right; j >= left; j-- {
				res = append(res, matrix[bottom][j])
			}
			bottom--
		}
		// ↑ 左列：从下到上
		if left <= right {
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}
	return res
}
