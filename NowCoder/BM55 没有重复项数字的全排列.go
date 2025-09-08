package nowcoder

// https://www.nowcoder.com/practice/4bcf3081067a4d028f95acee3ddcd2b1

func permute(num []int) [][]int {
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
			if used[num[j]] {
				continue
			}

			path = append(path, num[j])
			used[num[j]] = true
			backtarck(i + 1)
			path = path[:len(path)-1]
			used[num[j]] = false
		}
	}

	backtarck(0)

	return ans
}
