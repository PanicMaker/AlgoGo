package Array

import (
	"math"
	"sort"
)

// https://leetcode.cn/problems/minimum-size-subarray-sum/

func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	sum := 0
	res := math.MaxInt

	for right < len(nums) {
		sum += nums[right]
		right++

		for sum >= target {
			res = min(res, right-left)
			sum -= nums[left]
			left++
		}
	}

	if res == math.MaxInt {
		return 0
	}

	return res
}

func minSubArrayLen2(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	// 1. 构建前缀和数组
	// preSum[i] 存储 nums[0...i-1] 的和
	// 数组长度为 n+1，preSum[0] = 0，方便计算
	preSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}
	res := n + 1 // 初始化结果为一个不可能的值
	// 2. 遍历每个可能的子数组起点
	// i 对应的是原数组的索引起点，对应到 preSum 数组中是 preSum[i]
	for i := 0; i < n; i++ {
		// 目标：寻找一个 j > i，使得 preSum[j] - preSum[i] >= target
		// 等价于：preSum[j] >= target + preSum[i]
		searchTarget := target + preSum[i]
		// 3. 在 preSum 数组的 [i+1, n] 范围内二分查找
		// Go 的 sort.SearchInts 在整个切片中查找，返回第一个 >= target 的索引
		// 我们需要在 preSum[i+1:] 这个子切片上进行查找
		// sort.SearchInts 返回的是在 *整个* preSum 数组中的索引
		// bound 是满足 preSum[bound] >= searchTarget 的最小索引
		// 注意这里的查找范围，因为j必须大于i
		// sort.Search 是一个更通用的二分查找函数
		// 它在 [0, len) 中查找第一个使 f(k) 为 true 的索引 k
		bound := sort.Search(len(preSum), func(k int) bool {
			return preSum[k] >= searchTarget
		})

		// 如果找到了这样的 bound 并且它在数组范围内
		// preSum[bound] 实际上是 j 的位置
		if bound <= n {
			// j = bound 是 preSum 数组中的索引
			// 对应的子数组长度是 j - i
			length := bound - i
			if length < res {
				res = length
			}
		}
	}
	if res == n+1 {
		return 0
	}
	return res
}
