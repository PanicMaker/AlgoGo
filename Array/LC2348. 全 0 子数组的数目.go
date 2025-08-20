package Array

// https://leetcode.cn/problems/number-of-zero-filled-subarrays/description

func zeroFilledSubarray1(nums []int) int64 {

	maps0 := make(map[int]int)

	for i := 0; i < len(nums); {
		if nums[i] == 0 {
			start := i
			cnt := 0
			for i < len(nums) && nums[i] == 0 {
				cnt++
				i++
			}
			maps0[start] = cnt
		} else {
			i++
		}
	}

	sum := 0

	for _, v := range maps0 {
		tmp := (1 + v) * v / 2
		sum += tmp
	}

	return int64(sum)
}

func zeroFilledSubarray2(nums []int) int64 {
	var res, count int64

	for _, v := range nums {
		if v == 0 {
			count++
			res += count
		} else {
			count = 0
		}
	}

	return res
}
