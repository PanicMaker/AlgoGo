package Sort

func CountingSort(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	minVal, maxVal := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if maxVal < nums[i] {
			maxVal = nums[i]
		} else if minVal > nums[i] {
			minVal = nums[i]
		}
	}

	count := make([]int, maxVal-minVal+1)
	for _, v := range nums {
		count[v-minVal]++
	}

	// 累加计数数组
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	res := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		v := nums[i]
		res[count[v-minVal]-1] = v
		count[v-minVal]--
	}

	return res
}
