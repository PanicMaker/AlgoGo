package Sort

func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	left, right := 0, len(nums)-1

	mid := left + (right-left)/2

	al := MergeSort(nums[:mid+1])
	ar := MergeSort(nums[mid+1:])

	return merge(al, ar)
}

func merge(arr1 []int, arr2 []int) []int {
	tmp := make([]int, 0)

	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			tmp = append(tmp, arr1[i])
			i++
		} else {
			tmp = append(tmp, arr2[j])
			j++
		}
	}

	tmp = append(tmp, arr1[i:]...)
	tmp = append(tmp, arr2[j:]...)

	return tmp
}
