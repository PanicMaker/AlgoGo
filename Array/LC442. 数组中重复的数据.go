package Array

// https://leetcode.cn/problems/find-all-duplicates-in-an-array/description/

func findDuplicates(nums []int) []int {
	ans := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 || i == nums[i]-1 {
			continue
		}

		if nums[i] == nums[nums[i]-1] {
			ans = append(ans, nums[i])
			nums[i] *= -1
		} else {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
			i--
		}
	}

	return ans
}
