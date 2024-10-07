package Tree

// https://leetcode.cn/problems/convert-sorted-list-to-binary-search-tree/

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortedListToBST(head *ListNode) *TreeNode {

	middleNode := func(start, end *ListNode) *ListNode {
		slow, fast := start, start

		for fast != end && fast.Next != end {
			fast = fast.Next.Next
			slow = slow.Next
		}

		return slow
	}

	var dfs func(start, end *ListNode) *TreeNode
	dfs = func(start, end *ListNode) *TreeNode {
		if start == end {
			return nil
		}

		mid := middleNode(start, end)

		root := &TreeNode{
			Val: mid.Val,
		}
		root.Left = dfs(start, mid)
		root.Right = dfs(mid.Next, end)
		return root
	}

	return dfs(head, nil)
}

func sortedListToBST2(head *ListNode) *TreeNode {
	nums := make([]int, 0)
	for cur := head; cur != nil; cur = cur.Next {
		nums = append(nums, cur.Val)
	}

	var dfs func(nums []int, lo, hi int) *TreeNode
	dfs = func(nusm []int, lo, hi int) *TreeNode {
		if lo > hi {
			return nil
		}

		mid := lo + (hi-lo)/2

		root := &TreeNode{
			Val: nums[mid],
		}
		root.Left = dfs(nums, lo, mid-1)
		root.Right = dfs(nums, mid+1, hi)
		return root
	}

	return dfs(nums, 0, len(nums)-1)
}
