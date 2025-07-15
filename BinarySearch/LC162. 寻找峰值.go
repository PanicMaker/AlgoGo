package binarysearch

// https://leetcode.cn/problems/find-peak-element/

func findPeakElement(nums []int) int {
	l, r := 0, len(nums)

	for l < r {
		mid := l + (r-l)>>2

		if nums[mid] > nums[mid+1] {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}
