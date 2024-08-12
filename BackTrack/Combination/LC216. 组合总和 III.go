package Combination

// https://leetcode.cn/problems/combination-sum-iii/

func combinationSum3I(k int, n int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)

	var backtrack func(i int, target int)
	backtrack = func(i int, target int) {
		if target < 0 || len(path) > k {
			return
		}

		if target == 0 && len(path) == k {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}

		for j := i; j < 10; j++ {
			target -= j
			path = append(path, j)
			backtrack(j+1, target)
			target += j
			path = path[:len(path)-1]
		}
	}

	backtrack(1, n)

	return ans
}

func combinationSum3II(k int, n int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)

	var backtrack func(i int, target int)
	backtrack = func(i int, target int) {
		if target < 0 || len(path) > k {
			return
		}

		if target == 0 && len(path) == k {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}

		if i > 9 {
			return
		}

		backtrack(i+1, target)

		path = append(path, i)
		backtrack(i+1, target-i)
		path = path[:len(path)-1]
	}

	backtrack(1, n)

	return ans
}
