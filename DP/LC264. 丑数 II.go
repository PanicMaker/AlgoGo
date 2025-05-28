package DP

// https://leetcode.cn/problems/ugly-number-ii/

import "container/heap"

func nthUglyNumberI(n int) int {

	dp := make([]int, n)
	dp[0] = 1

	a, b, c := 0, 0, 0

	for i := 1; i < n; i++ {
		n2, n3, n5 := dp[a]*2, dp[b]*3, dp[c]*5
		dp[i] = min(n2, n3, n5)
		if dp[i] == n2 {
			a++
		}
		if dp[i] == n3 {
			b++
		}
		if dp[i] == n5 {
			c++
		}
	}

	return dp[n-1]
}

type intHeap []int

func (hp intHeap) Len() int {
	return len(hp)
}

func (hp intHeap) Less(i, j int) bool {
	return hp[i] < hp[j]
}

func (hp intHeap) Swap(i, j int) {
	hp[i], hp[j] = hp[j], hp[i]
}

func (hp *intHeap) Pop() any {

}

func (hp *intHeap) Push(x any) {
	*hp = append(*hp, x.(int))
}

func nthUglyNumberII(n int) int {

}
