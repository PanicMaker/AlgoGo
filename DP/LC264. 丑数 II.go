package DP

import "container/heap"

// https://leetcode.cn/problems/ugly-number-ii/

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
	old := hp
	n := len(*old)
	if n == 0 {
		return nil
	}
	x := (*old)[n-1]
	*old = (*old)[:n-1]
	return x
}

func (hp *intHeap) Push(x any) {
	*hp = append(*hp, x.(int))
}

// 使用最小堆解决
func nthUglyNumberII(n int) int {
	if n == 1 {
		return 1
	}

	hp := &intHeap{1}
	seen := map[int]bool{1: true}

	for i := 1; i < n; i++ {
		x := (*hp)[0]
		seen[x] = true
		heap.Pop(hp)

		for _, factor := range []int{2, 3, 5} {
			next := x * factor
			if !seen[next] {
				heap.Push(hp, next)
				seen[next] = true
			}
		}
	}

	return (*hp)[0]
}
