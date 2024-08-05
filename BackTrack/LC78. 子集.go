package BackTrack

// https://leetcode.cn/problems/subsets/

func subsets(nums []int) [][]int {
	n := len(nums)

	ans := make([][]int, 0)
	path := make([]int, 0)

	var dfs func(i int)
	dfs = func(i int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		ans = append(ans, tmp)

		for j := i; j < n; j++ {
			path = append(path, nums[j])
			dfs(j + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(0)

	return ans
}
