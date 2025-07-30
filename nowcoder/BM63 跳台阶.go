package nowcoder

// https://www.nowcoder.com/practice/8c82a5b80378478f9484d87d1c5f12a4

func jumpFloor(number int) int {

	if number == 0 {
		return 0
	}

	if number == 1 {
		return 1
	}

	if number == 2 {
		return 2
	}

	return jumpFloor(number-1) + jumpFloor(number-2)
}

func jumpFloorII(number int) int {

	dp := make([]int, 41)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2

	for i := 2; i <= number; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[number]
}

func jumpFloorIII(number int) int {

	if number <= 1 {
		return 1
	}

	res := 0
	a, b := 1, 1

	for i := 2; i <= number; i++ {
		res = a + b
		a, b = b, res
	}

	return res
}
