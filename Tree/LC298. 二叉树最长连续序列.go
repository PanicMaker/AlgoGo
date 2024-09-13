package Tree

// https://leetcode.cn/problems/binary-tree-longest-consecutive-sequence/description/

func longestConsecutive(root *TreeNode) int {
	ans := 0

	var dfs func(node *TreeNode, pre *TreeNode, path int)
	dfs = func(node *TreeNode, pre *TreeNode, path int) {
		if node == nil {
			return
		}

		if pre != nil && pre.Val+1 == node.Val {
			path++
		} else {
			path = 1
		}

		ans = max(ans, path)

		dfs(node.Left, node, path)
		dfs(node.Right, node, path)
	}

	dfs(root, nil, 0)

	return ans
}
