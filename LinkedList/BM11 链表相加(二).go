package LinkedList

// https://www.nowcoder.com/practice/c56f6c70fb3f4849bc56e33ff2a50b6b

func addInList(head1 *ListNode, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}

	dummy := &ListNode{}
	stack1, stack2 := make([]*ListNode, 0), make([]*ListNode, 0)

	// 将链表节点压入栈中
	for p1 := head1; p1 != nil; p1 = p1.Next {
		stack1 = append(stack1, p1)
	}
	for p2 := head2; p2 != nil; p2 = p2.Next {
		stack2 = append(stack2, p2)
	}

	carry := 0
	// 逐位相加，处理两个栈中的节点
	for len(stack1) > 0 || len(stack2) > 0 || carry > 0 {
		sum := carry
		if len(stack1) > 0 {
			sum += stack1[len(stack1)-1].Val
			stack1 = stack1[:len(stack1)-1] // 弹出栈顶元素
		}
		if len(stack2) > 0 {
			sum += stack2[len(stack2)-1].Val
			stack2 = stack2[:len(stack2)-1] // 弹出栈顶元素
		}

		carry = sum / 10
		tmp := &ListNode{
			Val: sum % 10,
		}
		tmp.Next = dummy.Next
		dummy.Next = tmp
	}

	return dummy.Next
}
