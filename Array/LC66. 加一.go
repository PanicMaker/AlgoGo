package Array

import (
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/plus-one/description/

// 会超出Int最大长度
func plusOne1(digits []int) []int {
	num := 0

	for i := range digits {
		num = num*10 + digits[i]
	}

	numStrs := strings.Split(string(num), "")
	res := make([]int, 0)

	for i := range numStrs {
		num, _ := strconv.Atoi(numStrs[i])
		res = append(res, num)
	}

	return res
}

func plusOne2(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += 1
		digits[i] %= 10

		if digits[i] != 0 {
			return digits
		}
	}

	newArr := make([]int, len(digits)+1)
	newArr[0] = 1
	return newArr
}
