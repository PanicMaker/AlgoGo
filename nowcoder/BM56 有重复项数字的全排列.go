package nowcoder

// https://www.nowcoder.com/practice/a43a2b986ef34843ac4fdd9159b69863

func permuteUnique(num []int) [][]int {
	n := len(num)
	if n == 0 {
		return nil
	}

	ans := make([][]int, 0)
	path := make([]int, 0)
	used := make(map[int]bool)

	var backtarck func(i int)
	backtarck = func(i int) {
		if i == n {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}

		for j := 0; j < n; j++ {
			if used[j] {
				continue
			}

			if j > 0 && num[j] == num[j-1] && used[j-1] {
				continue
			}

			path = append(path, num[j])
			used[j] = true
			backtarck(i + 1)
			path = path[:len(path)-1]
			used[j] = false
		}
	}

	backtarck(0)

	return ans
}
