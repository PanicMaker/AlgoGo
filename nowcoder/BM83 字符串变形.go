package nowcoder

import (
	"strings"
)

// https://www.nowcoder.com/practice/c3120c1c1bc44ad986259c0cf0f0b80e

func trans(s string, n int) string {
	strs := strings.Split(s, " ")

	for i, j := 0, len(strs)-1; i < j; i, j = i+1, j-1 {
		strs[i], strs[j] = strs[j], strs[i]
	}

	for i := 0; i < len(strs); i++ {
		bytes := []byte(strs[i])
		for j := 0; j < len(bytes); j++ {
			if bytes[j] >= 'A' && bytes[j] <= 'Z' {
				bytes[j] += 32
			} else if bytes[j] >= 'a' && bytes[j] <= 'z' {
				bytes[j] -= 32
			}
		}
		strs[i] = string(bytes)
	}
	return strings.Join(strs, " ")
}
