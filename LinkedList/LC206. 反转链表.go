package LinkedList

// https://leetcode.cn/problems/reverse-linked-list/

// 迭代
func reverseList1(head *ListNode) *ListNode {
	pre := (*ListNode)(nil)
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

// 递归

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList2(head.Next)

	head.Next.Next = head
	head.Next = nil

	return newHead
}
