package Sort

func QuickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	pivot := nums[len(nums)/2]
	left, right := []int{}, []int{}

	for i, v := range nums {
		if i == len(nums)/2 {
			continue
		}
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	return append(append(QuickSort(left), pivot), QuickSort(right)...)
}

func QuickSort2(nums []int) {
	// 三数取中法选择pivot
	mid := len(nums) / 2
	if nums[0] > nums[mid] {
		nums[0], nums[mid] = nums[mid], nums[0]
	}
	if nums[0] > nums[len(nums)-1] {
		nums[0], nums[len(nums)-1] = nums[len(nums)-1], nums[0]
	}
	if nums[mid] > nums[len(nums)-1] {
		nums[mid], nums[len(nums)-1] = nums[len(nums)-1], nums[mid]
	}
	pivot := nums[mid]

	// 原地分区（Lomuto分区）
	left, right := 0, len(nums)-1
	for left < right {
		for left < right && nums[left] <= pivot {
			left++
		}
		for left < right && nums[right] >= pivot {
			right--
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}

	// 尾递归优化：先处理较短的子数组
	if left > 0 {
		QuickSort(nums[:left])
	}
	if left < len(nums)-1 {
		QuickSort(nums[left+1:])
	}
}
