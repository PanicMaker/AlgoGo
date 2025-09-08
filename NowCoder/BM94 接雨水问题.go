package nowcoder

// https://www.nowcoder.com/practice/31c1aed01b394f0b8b7734de0324e00f

func maxWater(arr []int) int64 {
	if len(arr) == 0 {
		return 0
	}

	l, r := 0, len(arr)-1
	lMax, rMax := 0, 0

	ans := 0

	for l < r {
		if arr[l] > lMax {
			lMax = arr[l]
		}
		if arr[r] > rMax {
			rMax = arr[r]
		}

		if arr[l] < arr[r] {
			ans += lMax - arr[l]
			l++
		} else {
			ans += rMax - arr[r]
			r--
		}
	}

	return int64(ans)
}
