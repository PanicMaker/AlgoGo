package Array

// https://www.nowcoder.com/practice/96bd6684e04a44eb80e6a68efc0ec6c5

func InversePairs(nums []int) int {
	// write code here

	_, res := MergeSortAndCount(nums)
	return res % 1000000007
}

func MergeSortAndCount(nums []int) ([]int, int) {
	if len(nums) <= 1 {
		return nums, 0
	}

	mid := len(nums) / 2
	al, countLeft := MergeSortAndCount(nums[:mid])
	ar, countRight := MergeSortAndCount(nums[mid:])
	merged, splitCount := mergeAndCount(al, ar)

	return merged, countLeft + countRight + splitCount
}

func mergeAndCount(arr1, arr2 []int) ([]int, int) {
	tmp := make([]int, 0, len(arr1)+len(arr2))
	i, j, count := 0, 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			tmp = append(tmp, arr1[i])
			i++
		} else {
			tmp = append(tmp, arr2[j])
			j++
			// 关键：左半部分剩余元素均与arr2[j]构成逆序对
			count += len(arr1) - i
		}
	}

	tmp = append(tmp, arr1[i:]...)
	tmp = append(tmp, arr2[j:]...)

	return tmp, count
}
