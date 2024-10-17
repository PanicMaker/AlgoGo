package binarysearch

// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/description/

func searchRange1(nums []int, target int) []int {

	search := func(nums []int, target int) int {
		left, right := 0, len(nums)

		for left < right {
			mid := left + (right-left)/2

			if nums[mid] < target {
				left = mid + 1
			} else if nums[mid] > target {
				right = mid
			} else {
				return mid
			}
		}

		return -1
	}

	left := search(nums, target)

	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}

	right := search(nums, target+1) - 1

	return []int{left, right}
}

// 两次二分查找
func searchRange2(nums []int, target int) []int {

	idxL, idxR := -1, -1

	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		} else {
			idxL = mid
			right = mid
		}

	}

	left, right = 0, len(nums)

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		} else {
			idxR = mid
			left = mid + 1
		}
	}

	return []int{idxL, idxR}
}
