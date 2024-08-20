package Tree

import "math"

// https://leetcode.cn/problems/binary-tree-maximum-path-sum/description/

func maxPathSum(root *TreeNode) int {
	res := math.MinInt

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		res = max(res, node.Val+left+right, node.Val+left, node.Val+right, node.Val)

		return max(node.Val+left, node.Val+right, node.Val)
	}

	dfs(root)

	return res
}
