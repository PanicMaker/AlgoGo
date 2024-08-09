package Combination

// https://leetcode.cn/problems/combination-sum/description

func combinationSum(candidates []int, target int) [][]int {
	n := len(candidates)
	ans := make([][]int, 0)
	path := make([]int, 0)

	var backtrack func(i int, target int)
	backtrack = func(i int, target int) {

		if target == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}

		if target < 0 {
			return
		}

		for j := i; j < n; j++ {
			newTarget := target - candidates[j]
			path = append(path, candidates[j])
			backtrack(j+1, newTarget)
			target += candidates[j]
			path = path[:len(path)-1]
		}
	}

	backtrack(0, target)

	return ans
}
