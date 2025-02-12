package Array

// https://leetcode.cn/problems/find-all-duplicates-in-an-array/description/

func findDuplicates(nums []int) []int {
	// 用于存储最终结果的数组
	ans := make([]int, 0)

	// 遍历数组中的每一个元素
	for i := 0; i < len(nums); i++ {
		// 如果当前元素已经被标记为负数（表示已经处理过），或者当前元素已经在正确的位置上（即 nums[i] == i+1），则跳过
		if nums[i] < 0 || i == nums[i]-1 {
			continue
		}

		// 如果当前元素与它应该在的位置上的元素相等，说明找到了一个重复的元素
		if nums[i] == nums[nums[i]-1] {
			// 将这个重复的元素加入到结果数组中
			ans = append(ans, nums[i])
			// 将这个元素标记为负数，表示已经处理过
			nums[i] *= -1
		} else {
			// 如果当前元素与它应该在的位置上的元素不相等，则交换这两个元素的位置
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
			// 由于交换后当前元素可能还没有处理完，所以需要将索引减1，以便在下一次循环中重新处理当前元素
			i--
		}
	}

	// 返回最终的结果数组
	return ans
}
