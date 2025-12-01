package Array

import "math"

// https://leetcode.cn/problems/maximum-subarray-sum-with-length-divisible-by-k/description/

func maxSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	if n == 0 {
		return 0
	}
	// 1. 使用 int64 避免整数溢出
	preSum := make([]int64, n+1)
	for i := 0; i < n; i++ {
		preSum[i+1] = preSum[i] + int64(nums[i])
	}
	// 2. 初始化结果为 int64 的最小值
	maxSum := int64(math.MinInt64)
	// 3. minPreSum 用于存储 {余数 -> 该余数下遇到的最小前缀和}
	// key: j % k 的余数
	// value: 在所有余数为 key 的下标中，所对应的最小前缀和 preSum[j]
	minPreSum := make(map[int]int64)
	// 4. 遍历前缀和数组
	for j := 0; j <= n; j++ {
		currentSum := preSum[j]
		rem := j % k
		// 查找是否存在一个之前的下标 i, 使得 i % k == j % k
		if minP, ok := minPreSum[rem]; ok {
			// 如果存在，说明我们找到了一个有效的子数组
			// 其长度 (j-i) 是 k 的倍数
			// 计算 preSum[j] - preSum[i] 的值，并更新最大和
			// minP 就是我们为当前余数 rem 找到的最小的 preSum[i]
			candidate := currentSum - minP
			if candidate > maxSum {
				maxSum = candidate
			}
		}
		// 无论如何，都要更新当前余数对应的最小前缀和
		// 如果 rem 不存在，或者 currentSum 比已存的值更小，就更新
		if _, ok := minPreSum[rem]; !ok || currentSum < minPreSum[rem] {
			minPreSum[rem] = currentSum
		}
	}
	// 如果 maxSum 没有被更新过 (例如所有满足条件的子数组和都是负数，但比MinInt64大)，
	// 并且没有任何子数组满足条件，我们应该返回0吗？
	// 题意是求“最大和”，如果一个满足条件的子数组都找不到，理论上没有和。
	// 但通常这类问题会默认一个空子数组和为0。如果题目要求找不到就返回0，可以加个判断。
	// 但根据原始代码逻辑，找不到时会返回 MinInt，这里保持一致。
	// 如果所有满足条件的子数组和都是负数，那最大和就是那个“最不负”的负数。
	// 如果一个满足条件的子数组都没有，maxSum会保持MinInt64。这种情况下，可能需要根据题意返回0或错误。
	// 假设至少存在一个长度为k的倍数的子数组，或者可以返回0。这里我们先按找到的最大值返回。
	// 一个长度为0的子数组（j=i）和为0，且长度0%k==0，所以结果至少是0。
	// 我们的代码在 j=i 时，currentSum - minP = 0，所以 maxSum 至少会是 0（如果所有和都为负）。
	if maxSum == int64(math.MinInt64) {
		return 0 // 或者根据题目具体要求返回
	}
	return maxSum
}
