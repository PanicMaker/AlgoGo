package Combination

// https://leetcode.cn/problems/combinations/description/

// 枚举下一个数选什么
func combineI(n int, k int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)

	var backtrack func(i int)
	backtrack = func(i int) {
		if n-i+1 < k-len(path) {
			return
		}

		if len(path) == k {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}

		for j := i; j < n+1; j++ {
			path = append(path, j)
			backtrack(j + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(1)

	return ans
}

// 选或不选
func combineII(n int, k int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)

	var backtrack func(i int)
	backtrack = func(i int) {
		if len(path) == k {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}

		// 可选数大于还需要的，可以跳过选择
		if n-i+1 > k-len(path) {
			backtrack(i + 1)
		}

		// 选择当前数
		path = append(path, i)
		backtrack(i + 1)
		path = path[:len(path)-1]
	}

	backtrack(1)

	return ans
}
