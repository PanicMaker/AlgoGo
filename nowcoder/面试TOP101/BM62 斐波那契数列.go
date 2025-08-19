package nowcoder

// https://www.nowcoder.com/practice/c6c7742f5ba7442aada113136ddea0c3

func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func FibonacciII(n int) int {
	dp := make([]int, 41)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func FibonacciIII(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}
