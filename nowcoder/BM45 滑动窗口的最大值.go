package nowcoder

// https://www.nowcoder.com/practice/1624bc35a45c42c0bc17d17fa0cba788

// 滑动窗口
func maxInWindows1(num []int, size int) []int {
	if size > len(num) || size == 0 {
		return nil
	}

	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	maxVal := func(nums []int) int {
		res := nums[0]
		for _, v := range nums {
			res = max(res, v)
		}
		return res
	}(num[:size])

	ans := make([]int, 0)
	ans = append(ans, maxVal)

	l, r := 0, size
	window := []int(num[:size])

	for r < len(num) {
		new := num[r]
		r++

		window = append(window, new)

		for len(window) > size {
			old := num[l]
			l++
			window = window[1:]

			if old == maxVal {
				if old <= new {
					maxVal = new
					ans = append(ans, new)
				} else {
					// TODO: 在新窗口中选出最大值
					maxVal = window[0]
					for _, v := range window {
						if v > maxVal {
							maxVal = v
						}
					}
					ans = append(ans, maxVal)
				}
			} else {
				if maxVal <= new {
					maxVal = new
					ans = append(ans, new)
				} else {
					ans = append(ans, maxVal)
				}
			}
		}

	}

	return ans
}

// 双向队列
func maxInWindows2(num []int, size int) []int {
	if size > len(num) || size == 0 {
		return nil
	}

	ans := []int{}
	deque := []int{}

	for i := 0; i < len(num); i++ {
		if len(deque) > 0 && deque[0] <= i-size {
			deque = deque[1:]
		}

		for len(deque) > 0 && num[deque[len(deque)-1]] < num[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)

		if i >= size-1 {
			ans = append(ans, num[deque[0]])
		}
	}

	return ans
}
