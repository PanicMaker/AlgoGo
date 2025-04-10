package LinkedList

// https://www.nowcoder.com/practice/02bf49ea45cd486daa031614f9bd6fc3

func oddEvenList(head *ListNode) *ListNode {
	// write code here
	if head == nil || head.Next == nil {
		return head
	}

	odd, even := head, head.Next
	evenHead := even

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenHead

	return head
}
