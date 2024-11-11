package Array

// https://leetcode.cn/problems/h-index/description

func hIndex(citations []int) int {
	res := 0
	n := len(citations)

	for i := 0; i <= n; i++ {
		count := 0

		for _, v := range citations {
			if v >= i {
				count++
			}
		}

		if count >= i {
			res = max(res, i)
		}

	}

	return res
}
