package BackTrack

// https://leetcode.cn/problems/palindrome-partitioning/description/

func partition(s string) [][]string {
	n := len(s)

	ans := make([][]string, 0)
	path := make([]string, 0)

	var backtrack func(i int)
	backtrack = func(i int) {
		if i == n {
			tmp := make([]string, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}

		for j := i; j < n; j++ {
			str := s[i : j+1]
			if isPalindrome(str) {
				path = append(path, str)
				backtrack(j + 1)
				path = path[:len(path)-1]
			}
		}
	}

	backtrack(0)

	return ans
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	left, right := 0, len(s)-1

	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}
