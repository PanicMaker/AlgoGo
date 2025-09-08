package algogo

// https://leetcode.cn/problems/reverse-integer/description/

func reverse(x int) int {
	const (
		minInt32 = -1 << 31
		maxInt32 = 1<<31 - 1
	)

	var result int
	for x != 0 {
		digit := x % 10
		x /= 10

		if result > maxInt32/10 || (result == maxInt32/10 && digit > 7) {
			return 0 // overflow
		}
		if result < minInt32/10 || (result == minInt32/10 && digit < -8) {
			return 0 // underflow
		}

		result = result*10 + digit
	}

	return result
}
