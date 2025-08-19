package nowcoder

// https://www.nowcoder.com/practice/5164f38b67f846fb8699e9352695cd2f

func LIS(arr []int) int {
	n := len(arr)

	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
	}

	maxLength := 1

	for _, v := range dp {
		maxLength = max(v, maxLength)
	}

	return maxLength
}
