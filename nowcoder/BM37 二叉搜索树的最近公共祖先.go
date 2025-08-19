package nowcoder

func lowestCommonAncestorBST(root *TreeNode, p int, q int) int {
	// write code here

	x := root.Val

	if p < x && q < x {
		return lowestCommonAncestorBST(root.Left, p, q)
	}

	if p > x && q > x {
		return lowestCommonAncestorBST(root.Right, p, q)
	}

	return x
}
