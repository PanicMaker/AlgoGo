package greddy

import (
	"container/heap"
	"sort"
)

// https://leetcode.cn/problems/maximum-number-of-events-that-can-be-attended/description/

type hp struct {
	sort.IntSlice
}

func (h *hp) Push(val any) {
	h.IntSlice = append(h.IntSlice, val.(int))
}

func (h *hp) Pop() any {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func maxEvents(events [][]int) int {
	lastDay := 0
	for _, v := range events {
		lastDay = max(lastDay, v[1])
	}

	groups := make([][]int, lastDay+1)
	for _, e := range events {
		groups[e[0]] = append(groups[e[0]], e[1])
	}

	ans := 0
	h := &hp{}
	heap.Init(h)
	for i, g := range groups {
		// 删除过期会议
		for h.Len() > 0 && h.IntSlice[0] < i {
			heap.Pop(h)
		}
		// 新增可以参加的会议
		for _, endDay := range g {
			heap.Push(h, endDay)
		}
		// 参加一个结束时间最早的会议
		if h.Len() > 0 {
			ans++
			heap.Pop(h)
		}
	}

	return ans
}
