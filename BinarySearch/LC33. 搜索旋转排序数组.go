package binarysearch

// https://leetcode.cn/problems/search-in-rotated-sorted-array/description

func search33(nums []int, target int) int {
	left, right := 0, len(nums)-1

	// 找到旋转点
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	minIndex := left

	// 判断 target 属于哪一段
	if target >= nums[minIndex] && target <= nums[len(nums)-1] {
		left, right = minIndex, len(nums)-1
	} else {
		left, right = 0, minIndex-1
	}

	// 标准二分查找
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
