package main

import (
	"strconv"
)

// https://leetcode.cn/problems/palindrome-number/description/

func isPalindromeI(x int) bool {
	str := strconv.Itoa(x)

	i, j := 0, len(str)-1

	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func isPalindromeII(x int) bool {
	if x < 0 {
		return false
	}

	copyNum := x
	reverse := 0

	for copyNum > 0 {
		temp := copyNum % 10
		copyNum /= 10
		reverse = reverse*10 + temp
	}

	if x == reverse {
		return true
	}

	return false
}
