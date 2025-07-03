package String

import (
	"strings"
)

// https://leetcode.cn/problems/find-the-k-th-character-in-string-game-i

func kthCharacter(k int) byte {
	words := "abcdefghigklmnopqrstuvwxyz"

	bytes := []byte{'a'}

	for i := 0; i < k; i++ {
		n := len(bytes)
		if n >= k {
			break
		}
		for j := 0; j < n; j++ {
			index := strings.IndexByte(words, bytes[j])
			bytes = append(bytes, words[(index+1)%26])
		}
	}

	return bytes[k-1]
}
