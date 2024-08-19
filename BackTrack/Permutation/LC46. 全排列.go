package Permutation

// https://leetcode.cn/problems/permutations/description/

func permute(nums []int) [][]int {
	n := len(nums)

	ans := make([][]int, 0)
	path := make([]int, 0)
	used := make(map[int]bool)

	var backtarck func(i int)
	backtarck = func(i int) {
		if len(path) == n {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}

		for j := 0; j < n; j++ {
			if used[nums[j]] {
				continue
			}

			path = append(path, nums[j])
			used[nums[j]] = true
			backtarck(i + 1)
			path = path[:len(path)-1]
			used[nums[j]] = false
		}
	}

	backtarck(0)

	return ans
}
