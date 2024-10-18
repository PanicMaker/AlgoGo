package bitmanipulation

// https://leetcode.cn/problems/single-number/

func singleNumber1(nums []int) int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	for k, v := range m {
		if v == 1 {
			return k
		}
	}

	return -30001
}

// 异或运算
func singleNumber2(nums []int) int {
	x := 0

	for _, v := range nums {
		x ^= v
	}

	return x
}
