package nowcoder

// https://www.nowcoder.com/practice/508378c0823c423baa723ce448cbfd0c

func hasPathSum(root *TreeNode, sum int) bool {
	// 如果根节点为空，返回false
	if root == nil {
		return false
	}

	// 如果当前节点是叶子节点，检查节点值是否等于目标和
	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	}

	// 递归检查左右子树，目标和减去当前节点的值
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
