package nowcoder

// https://www.nowcoder.com/practice/3d8d6a8e516e4633a2244d2934e5aa47

func maxArea(height []int) int {
	n := len(height)
	if n < 2 {
		return 0
	}

	l, r := 0, n-1
	res := 0

	for l < r {
		capacity := min(height[l], height[r]) * (r - l)

		res = max(res, capacity)

		if height[l] < height[r] {
			l++
		} else {
			r--
		}

	}

	return res
}
