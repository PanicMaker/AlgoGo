package queue

// https://leetcode.cn/problems/time-needed-to-buy-tickets/

func timeRequiredToBuy(tickets []int, k int) int {
	queue := make([]int, len(tickets))
	for i := 0; i < len(tickets); i++ {
		queue[i] = i
	}

	time := 0

	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]
		time++
		tickets[head]--

		if head == k && tickets[k] == 0 {
			return time
		}

		if tickets[head] == 0 {
			continue
		}

		queue = append(queue, head)
	}

	return time
}
