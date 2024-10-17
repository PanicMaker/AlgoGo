package binarysearch

// https://leetcode.cn/problems/search-insert-position/description

// 左闭又开
func searchInsert1(nums []int, target int) int {
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

// 闭区间
func searchInsert2(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
