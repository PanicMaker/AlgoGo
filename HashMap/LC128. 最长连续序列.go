package HashMap

// https://leetcode.cn/problems/longest-consecutive-sequence

func longestConsecutive(nums []int) int {
	numMap := make(map[int]bool)
	for _, num := range nums {
		numMap[num] = true
	}
	maxLen := 0

	for _, num := range nums {
		if !numMap[num-1] {
			cur := 1
			for numMap[num+1] {
				cur++
				num++
			}

			if cur > maxLen {
				maxLen = cur
			}
		}
	}

	return maxLen
}
