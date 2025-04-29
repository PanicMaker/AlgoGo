package slidingwindow

// https://leetcode.cn/problems/count-subarrays-where-max-element-appears-at-least-k-times/description

func countSubarrays(nums []int, k int) int64 {
	maxVal := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		maxVal = max(maxVal, nums[i])
	}

	left, right := 0, 0
	res := 0
	count := 0

	for ; right < n; right++ {
		if nums[right] == maxVal {
			count++
		}

		for count >= k {
			res += n - right

			if nums[left] == maxVal {
				count--
			}
			left++
		}
	}

	return int64(res)
}
