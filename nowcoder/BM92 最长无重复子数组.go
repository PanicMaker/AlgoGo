package nowcoder

// https://www.nowcoder.com/practice/b56799ebfd684fb394bd315e89324fb4

func maxLength(arr []int) int {

	n := len(arr)

	window := make(map[int]int)

	l, r := 0, 0
	res := 0

	for r < n {
		num := arr[r]
		r++

		window[num]++

		for window[num] > 1 {
			d := arr[l]
			l++
			window[d]--
		}

		res = max(res, r-l)
	}

	return res
}
