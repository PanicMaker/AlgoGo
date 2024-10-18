package binarysearch

// https://leetcode.cn/problems/maximum-count-of-positive-integer-and-negative-integer/description/

func maximumCount(nums []int) int {
	search := func(nums []int, target int) int {
		left, right := 0, len(nums)

		for left < right {
			mid := left + (right-left)/2

			if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid
			}
		}

		return left
	}

	// 找到最左侧0的位置，不存在则是适合插入0的位置
	left := search(nums, 0)

	// 找到最右侧0位置的右边一位
	right := search(nums, 1)

	neg := left
	pos := len(nums) - right

	return max(pos, neg)
}
