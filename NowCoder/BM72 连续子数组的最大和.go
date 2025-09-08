package nowcoder

// https://www.nowcoder.com/practice/459bd355da1549fa8a49e350bf3df484

func FindGreatestSumOfSubArray(array []int) int {
	n := len(array)
	dp := make([]int, n)
	dp[0] = array[0]

	res := array[0]

	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+array[i], array[i])
		res = max(res, dp[i])
	}
	return res
}
