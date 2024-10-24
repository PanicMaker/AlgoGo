package LinkedList

// https://leetcode.cn/problems/merge-k-sorted-lists

// 循环迭代合并
func mergeKLists1(lists []*ListNode) *ListNode {

	// 合并两个有序链表
	mergeTwoLists := func(list1 *ListNode, list2 *ListNode) *ListNode {
		tmp := &ListNode{-1, nil}
		p := tmp
		p1 := list1
		p2 := list2

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

	var pre *ListNode

	for _, list := range lists {
		pre = mergeTwoLists(pre, list)
	}

	return pre
}

// 分治
func mergeKLists(lists []*ListNode) *ListNode {
	// 合并两个有序链表的函数
	mergeTwoLists := func(list1 *ListNode, list2 *ListNode) *ListNode {
		// 定义一个临时链表节点，作为合并后的链表头部
		tmp := &ListNode{-1, nil}
		p := tmp
		p1 := list1
		p2 := list2

		// 遍历两个链表，按顺序合并
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

		// 如果其中一个链表还没遍历完，直接链接到结果链表的末尾
		if p1 != nil {
			p.Next = p1
		}
		if p2 != nil {
			p.Next = p2
		}
		return tmp.Next
	}

	// 定义递归的合并函数
	var merge func(l int, r int) *ListNode
	merge = func(l int, r int) *ListNode {
		// 如果左右边界相等，直接返回对应的链表
		if l == r {
			return lists[l]
		}
		// 如果左边界大于右边界，返回空
		if l > r {
			return nil
		}
		// 计算中间位置
		mid := (l + r) >> 1
		// 递归合并左边和右边的链表
		return mergeTwoLists(merge(l, mid), merge(mid+1, r))
	}
	// 合并从0到len(lists)-1的所有链表
	return merge(0, len(lists)-1)
}
