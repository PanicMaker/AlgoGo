package nowcoder

// https://www.nowcoder.com/practice/e0cc33a83afe4530bcec46eba3325116

func lowestCommonAncestor(root *TreeNode, o1 int, o2 int) int {
	if root == nil {
		return -1
	}

	if o1 == root.Val || o2 == root.Val {
		return root.Val
	}

	left := lowestCommonAncestor(root.Left, o1, o2)
	right := lowestCommonAncestor(root.Right, o1, o2)

	if left != -1 && right != -1 {
		return root.Val
	}

	if left != -1 {
		return left
	}

	return right
}
