package nowcoder

// https://www.nowcoder.com/practice/b49c3dc907814e9bbfa8437c251b028e

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	n := 0

	for p := head; p != nil; p = p.Next {
		n++
	}

	dummy := &ListNode{}
	dummy.Next = head

	p := dummy

	for n >= k {
		cur := p.Next
		pre := p

		for i := 0; i < k; i++ {
			tmp := cur.Next
			cur.Next = pre
			pre = cur
			cur = tmp
		}

		tail := p.Next
		tail.Next = cur
		p.Next = pre
		p = tail

		n -= k
	}

	return dummy.Next
}
