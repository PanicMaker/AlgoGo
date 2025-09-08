package nowcoder

// https://www.nowcoder.com/practice/76039109dd0b47e994c08d8319faa352

func candy(arr []int) int {
	n := len(arr)

	if n == 1 {
		return 1
	}

	arrange := make([]int, n)
	for i := range arrange {
		arrange[i]++
	}

	for i := 1; i < n; i++ {
		if arr[i] > arr[i-1] {
			arrange[i] = arrange[i-1] + 1
		}
	}

	for i := n - 1; i > 0; i-- {
		if arr[i] < arr[i-1] && arrange[i] >= arrange[i-1] {
			arrange[i-1] = arrange[i] + 1
		}
	}

	res := 0
	for _, v := range arrange {
		res += v
	}

	return res
}
