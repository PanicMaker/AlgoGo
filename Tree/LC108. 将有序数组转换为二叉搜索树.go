package Tree

// https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/

func sortedArrayToBST(nums []int) *TreeNode {

	var dfs func(lo, hi int) *TreeNode
	dfs = func(lo, hi int) *TreeNode {
		if lo > hi {
			return nil
		}

		mid := lo + (hi-lo)/2

		root := &TreeNode{
			Val: nums[mid],
		}
		root.Left = dfs(lo, mid-1)
		root.Right = dfs(mid+1, hi)
		return root
	}

	return dfs(0, len(nums)-1)
}
