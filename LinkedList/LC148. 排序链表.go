package LinkedList

// https://leetcode.cn/problems/sort-list/description

// 使用归并排序
func sortListI(head *ListNode) *ListNode {
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

// 冒泡排序
func sortListII(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head
	for p.Next != nil {
		tmp := p.Next
		for tmp != nil {
			if tmp.Val < p.Val {
				tmp.Val, p.Val = p.Val, tmp.Val
			}
			tmp = tmp.Next
		}
		p = p.Next
	}

	return head
}

// 选择排序，重排节点
func sortListIII(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{}
	dummy.Next = head
	p := dummy

	for p.Next != nil {
		// 找到最小值节点及其前驱节点
		pre, preMin := p, p
		min := p.Next
		for cur := p.Next; cur != nil; cur = cur.Next {
			if cur.Val < min.Val {
				preMin = pre
				min = cur
			}
			pre = cur // pre移动到当前节点，作为下一个节点的前驱
		}

		// 将最小值节点移动到当前节点之后
		preMin.Next = min.Next
		min.Next = p.Next
		p.Next = min
		p = min
	}

	return dummy.Next
}

// 选择排序，交换值
func sortListIIII(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head
	for p != nil {
		minNode := p
		tmp := p.Next
		for tmp != nil {
			if tmp.Val < minNode.Val {
				minNode = tmp
			}
			tmp = tmp.Next
		}
		// 交换 p 和 minNode 的值
		p.Val, minNode.Val = minNode.Val, p.Val
		p = p.Next
	}

	return head
}
