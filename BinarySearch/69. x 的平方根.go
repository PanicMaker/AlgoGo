package binarysearch

// https://leetcode.cn/problems/sqrtx

func mySqrt(x int) int {
	left, right := 0, x

	for left <= right {
		mid := left + (right-left)/2

		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			left = mid + 1
		} else if mid*mid > x {
			right = mid - 1
		}
	}

	return right
}
