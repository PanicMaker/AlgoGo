package Array

// https://www.nowcoder.com/practice/e016ad9b7f0b45048c58a9f27ba618bf?

func findKth(a []int, n int, K int) int {

	var quickSelect func(left, right, k int) int
	quickSelect = func(left, right, k int) int {
		if left == right {
			return a[left]
		}

		pivot := a[right]
		storeIndex := left
		for i := left; i < right; i++ {
			if a[i] >= pivot {
				a[i], a[storeIndex] = a[storeIndex], a[i]
				storeIndex++
			}
		}

		a[right], a[storeIndex] = a[storeIndex], a[right]

		count := storeIndex - left + 1

		if k == count {
			return a[storeIndex]
		} else if k < count {
			return quickSelect(left, storeIndex-1, k)
		} else {
			return quickSelect(storeIndex+1, right, k-count)
		}
	}

	return quickSelect(0, n-1, K)
}
