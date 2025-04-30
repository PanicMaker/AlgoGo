package Tree

// https://www.nowcoder.com/practice/a9d0ecbacef9410ca97463e4a5c83be7

func Mirror(pRoot *TreeNode) *TreeNode {
	// write code here
	if pRoot == nil {
		return pRoot
	}

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Left == nil && node.Right == nil {
			return
		}

		dfs(node.Left)
		dfs(node.Right)
		node.Left, node.Right = node.Right, node.Left
	}

	dfs(pRoot)

	return pRoot
}
