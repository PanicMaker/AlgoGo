package binarysearch

// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/description/

// O(n)
func findMin1(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		if nums[left] > nums[right] {
			left++
		} else {
			right--
		}
	}

	return nums[left]
}

// 二分查找,O(log n)
func findMin2(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}
