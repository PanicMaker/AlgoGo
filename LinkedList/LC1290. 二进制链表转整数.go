package LinkedList

// https://leetcode.cn/problems/convert-binary-number-in-a-linked-list-to-integer

func getDecimalValue(head *ListNode) int {
	p := head
	len := 0
	for p != nil {
		len++
		p = p.Next
	}

	res := 0
	p = head
	for i := len - 1; i >= 0; i-- {
		res += p.Val << i
		p = p.Next
	}

	return res
}
