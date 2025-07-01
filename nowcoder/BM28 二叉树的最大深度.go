package nowcoder

// https://www.nowcoder.com/practice/8a2b2bf6c19b4f23a9bdb9b233eefa73

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}
