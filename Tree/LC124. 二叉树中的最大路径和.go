package Tree

import "math"

// https://leetcode.cn/problems/binary-tree-maximum-path-sum/description/

func maxPathSum(root *TreeNode) int {
	// res 用于存储当前找到的最大路径和，初始值为最小整数
	res := math.MinInt

	// 定义深度优先搜索函数 dfs，返回的是以当前节点为根的最大单边路径和
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		// 递归终止条件：如果节点为空，返回 0
		if node == nil {
			return 0
		}

		// 递归计算左右子树的最大路径和，忽略负值（如果 left 或 right 为负，则取 0）
		left := max(dfs(node.Left), 0)
		right := max(dfs(node.Right), 0)

		// 更新 res，计算通过当前节点的最大路径和
		// 有四种情况：
		// 1. 左子树 + 当前节点 + 右子树 (node.Val + left + right)
		// 2. 左子树 + 当前节点 (node.Val + left)
		// 3. 右子树 + 当前节点 (node.Val + right)
		// 4. 当前节点本身 (node.Val)
		res = max(res, node.Val+left+right)

		// 返回从当前节点出发的单边最大路径和
		// 即：只能选择左子树或右子树中的一个加上当前节点
		return max(node.Val+left, node.Val+right, node.Val)
	}

	// 从根节点开始进行深度优先搜索
	dfs(root)

	// 返回最大路径和
	return res
}
