package Array

// https://leetcode.cn/problems/maximum-sum-circular-subarray/description/

func maxSubarraySumCircular(nums []int) int {
	n := len(nums)
	dp1 := make([]int, n)
	dp1[0] = nums[0]
	dp2 := make([]int, n)
	dp2[0] = nums[0]

	maxVal, minVal, sum := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		dp1[i] = max(dp1[i-1]+nums[i], nums[i])
		maxVal = max(maxVal, dp1[i])
		dp2[i] = min(dp2[i-1]+nums[i], nums[i])
		minVal = min(minVal, dp1[i])
	}

	return max(maxVal, sum-minVal)
}
