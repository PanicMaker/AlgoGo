package algogo

// https://leetcode.cn/problems/ugly-number/description/

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	factors := []int{2, 3, 5}

	for _, num := range factors {
		for n%num == 0 {
			n /= num
		}
	}

	return n == 1
}
