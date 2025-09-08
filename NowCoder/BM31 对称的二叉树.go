package nowcoder

// https://www.nowcoder.com/practice/ff05d44dfdb04e1d83bdbdab320efbcb

func isSymmetrical(pRoot *TreeNode) bool {
	if pRoot == nil {
		return true
	}

	var dfs func(p *TreeNode, q *TreeNode) bool
	dfs = func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}

		if (p == nil && q != nil) || (p != nil && q == nil) {
			return false
		}

		if p.Val != q.Val {
			return false
		}

		return dfs(p.Left, q.Right) && dfs(p.Right, q.Left)
	}

	return dfs(pRoot.Left, pRoot.Right)
}
