package nowcoder

import (
	"container/heap"
	"sort"
)

// https://www.nowcoder.com/practice/4edf6e6d01554870a12f218c94e8a299

// IntHeap 定义一个整数最小堆，存储活动结束的时间点
// 我们用它来管理当前正在使用的主持人，堆顶是结束最早的那个活动的时间
type IntHeap []int

// 以下五个方法是实现 container/heap.Interface 接口所必需的
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // 决定是最小堆还是最大堆，这里是最小堆
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push 和 Pop 方法需要使用指针接收者，因为它们会修改切片的长度
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop 方法：正确的实现
// heap.Pop 会先将堆顶和最后一个元素交换，然后调用此方法
// 此方法只需要移除最后一个元素即可
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1] // 正确的缩容方式
	return x
}

// 计算所需的最少主持人数量
// n: 活动数量
// startEnd: 一个二维数组，每个子数组表示一个活动的 [开始时间, 结束时间]
func minimumNumberOfHost1(n int, startEnd [][]int) int {
	if n == 0 {
		return 0
	}
	// 核心思想：贪心算法
	// 1. 将所有活动按开始时间升序排序。
	// 2. 维护一个最小堆，存储当前所有正在进行的活动的结束时间。
	// 3. 遍历排序后的活动：
	//    a. 查看堆顶（最早结束的活动）。如果当前活动开始时间 >= 堆顶的结束时间，
	//       说明那个主持人已经空闲了，可以复用。我们将堆顶弹出。
	//    b. 将当前活动的结束时间压入堆中，表示分配了一个主持人（可能是复用的，也可能是新的）。
	//    c. 堆的大小就代表当前占用的主持人数量。我们需要记录这个过程中的峰值。
	// 步骤1：按开始时间对所有活动进行排序
	sort.Slice(startEnd, func(i, j int) bool {
		return startEnd[i][0] < startEnd[j][0]
	})
	// 步骤2：初始化最小堆，用于存储正在进行的活动的结束时间
	h := &IntHeap{}
	heap.Init(h)
	// 用于记录所需主持人的最大数量（即堆大小的峰值）
	maxHosts := 0
	// 步骤3：遍历所有活动
	for _, activity := range startEnd {
		start, end := activity[0], activity[1]
		// 检查是否有主持人可以被释放和复用
		// 如果堆不为空，并且堆顶（最早结束时间）小于等于当前活动的开始时间
		// 说明主持人已经空闲，可以主持新活动，将其从堆中移除
		// 注意这里用 for 循环，因为可能多个主持人同时变得空闲
		for h.Len() > 0 && (*h)[0] <= start {
			heap.Pop(h)
		}
		// 为当前活动分配一个主持人，将其结束时间加入堆中
		heap.Push(h, end)
		// 更新所需的最大主持人数量
		// 堆的当前大小就是此刻需要的主持人数量
		if h.Len() > maxHosts {
			maxHosts = h.Len()
		}
	}
	return maxHosts
}

type Event struct {
	time int // 事件发生的时间
	typ  int // 事件类型：+1 表示活动开始，-1 表示活动结束
}

func minmumNumberOfHost2(n int, startEnd [][]int) int {
	// 1. 如果没有活动，则不需要主持人
	if n == 0 {
		return 0
	}
	// 2. 准备事件点
	// 每个活动 [start, end] 产生两个事件：
	// 一个在 start 时间的 +1 事件 (需要主持人)
	// 一个在 end 时间的 -1 事件 (释放主持人)
	events := make([]Event, 0, n*2)
	for _, activity := range startEnd {
		start := activity[0]
		end := activity[1]
		events = append(events, Event{time: start, typ: 1}) // 活动开始，主持人需求 +1
		events = append(events, Event{time: end, typ: -1})  // 活动结束，主持人需求 -1
	}
	// 3. 收集并排序事件
	// 先按时间升序排序。
	// 如果时间相同，-1 事件 (活动结束) 应该在 +1 事件 (活动开始) 之前处理。
	// 这样可以确保如果一个活动恰好在另一个活动结束的瞬间开始，可以复用同一个主持人。
	sort.Slice(events, func(i, j int) bool {
		if events[i].time != events[j].time {
			return events[i].time < events[j].time
		}
		// 时间相同：-1 事件 (结束) 优先于 +1 事件 (开始)
		return events[i].typ < events[j].typ
	})
	// 4. 扫描事件点并记录最大值
	currentHosts := 0 // 当前活跃的主持人数量
	maxHosts := 0     // 整个过程中所需的最大主持人数量
	for _, event := range events {
		currentHosts += event.typ // 根据事件类型增减主持人数量
		if currentHosts > maxHosts {
			maxHosts = currentHosts // 更新最大值
		}
	}
	return maxHosts
}
