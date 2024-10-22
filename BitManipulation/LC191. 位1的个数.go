package bitmanipulation

// https://leetcode.cn/problems/number-of-1-bits/description/

func hammingWeight(n int) int {
	res := 0
	for n != 0 {
		// 判断n是不是奇数（ 2进制末尾为1）
		if n%2 == 1 {
			res++
		}
		// n向右移一位
		n = n >> 1
	}
	return res
}
