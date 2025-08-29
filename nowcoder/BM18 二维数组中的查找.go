package nowcoder

// https://www.nowcoder.com/practice/abc3fe2ce8e146608e868a70efebf62e

func Find(target int, array [][]int) bool {
	m, n := len(array), len(array[0])

	i, j := m-1, 0

	for i >= 0 && j < n {
		num := array[i][j]
		if num == target {
			return true
		} else if num < target {
			j++
		} else {
			i--
		}
	}

	return false
}
