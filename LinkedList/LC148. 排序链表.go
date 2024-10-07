package LinkedList

import "github.com/nxadm/tail"

// https://leetcode.cn/problems/sort-list/description

func sortList(head *ListNode) *ListNode {
	merge := func(l1, l2 *ListNode) *ListNode {
		tmp := &ListNode{-1, nil}
		p := tmp
		p1 := l1
		p2 := l2

		for p1 != nil && p2 != nil {
			if p1.Val > p2.Val {
				p.Next = p2
				p2 = p2.Next
			} else {
				p.Next = p1
				p1 = p1.Next
			}
			p = p.Next
		}

		if p1 != nil {
			p.Next = p1
		}

		if p2 != nil {
			p.Next = p2
		}

		return tmp.Next
	}

	middleNode := func(start, end *ListNode) *ListNode {
		slow, fast := start, start

		for fast != end && fast.Next != end {
			fast = fast.Next.Next
			slow = slow.Next
		}

		return slow
	}

	var sort func(head, tail *ListNode) *ListNode
	sort = func(head, tail *ListNode) *ListNode {
		if head == nil {
			return nil
		}

		if head.Next == tail {
			head.Next = nil
			return head
		}

		mid := middleNode(head, tail)

		return merge(sort(head, mid), sort(mid, tail))
	}

	return sort(head, nil)
}
