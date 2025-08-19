package nowcoder

// https://www.nowcoder.com/practice/96bd6684e04a44eb80e6a68efc0ec6c5

func InversePairs(nums []int) int {
	// 调用归并排序并统计逆序对数量
	_, res := MergeSortAndCount(nums)
	// 按题意对结果取模
	return res % 1000000007
}

// 归并排序并统计逆序对
func MergeSortAndCount(nums []int) ([]int, int) {
	// 递归终止条件：单个元素没有逆序对
	if len(nums) <= 1 {
		return nums, 0
	}

	// 分治：分成左右两部分
	mid := len(nums) / 2
	al, countLeft := MergeSortAndCount(nums[:mid])  // 左半部分及其逆序对数
	ar, countRight := MergeSortAndCount(nums[mid:]) // 右半部分及其逆序对数
	merged, splitCount := mergeAndCount(al, ar)     // 合并并统计跨区间逆序对

	// 返回合并后的有序数组和总逆序对数
	return merged, countLeft + countRight + splitCount
}

// 合并两个有序数组，并统计逆序对数量
func mergeAndCount(arr1, arr2 []int) ([]int, int) {
	tmp := make([]int, 0, len(arr1)+len(arr2)) // 存放合并后的有序数组
	i, j, count := 0, 0, 0

	// 合并两个有序数组，同时统计逆序对
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			tmp = append(tmp, arr1[i])
			i++
		} else {
			tmp = append(tmp, arr2[j])
			j++
			// 关键：arr1[i] > arr2[j]，则arr1[i]及其后面的元素都比arr2[j]大
			// 所以逆序对数量加上左半部分剩余元素个数
			count += len(arr1) - i
		}
	}

	// 把剩余元素补到结果数组
	tmp = append(tmp, arr1[i:]...)
	tmp = append(tmp, arr2[j:]...)

	return tmp, count
}
