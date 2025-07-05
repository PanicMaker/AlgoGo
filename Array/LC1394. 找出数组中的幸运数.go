package Array

// https://leetcode.cn/problems/find-lucky-integer-in-an-array/

func findLucky(arr []int) int {
	count := make(map[int]int)
	lucky := -1

	for _, v := range arr {
		count[v]++
	}
	for _, v := range arr {
		if count[v] == v && v > lucky {
			lucky = v
		}
	}
	return lucky
}
